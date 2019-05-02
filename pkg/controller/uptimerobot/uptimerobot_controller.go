/*

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package uptimerobot

import (
	"context"

	monitorsv1 "github.com/cloud104/tks-uptimerobot-controller/pkg/apis/monitors/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	controllerError "sigs.k8s.io/cluster-api/pkg/controller/error"
	"sigs.k8s.io/cluster-api/pkg/util"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	logf "sigs.k8s.io/controller-runtime/pkg/runtime/log"
	"sigs.k8s.io/controller-runtime/pkg/source"
)

var log = logf.Log.WithName("controller")

// AddWithActuator creates a new UptimeRobot Controller and adds it to the Manager with default RBAC. The Manager will set fields on the Controller
// and Start it when the Manager is Started.
func AddWithActuator(mgr manager.Manager, actuator Actuator) error {
	return add(mgr, newReconciler(mgr, actuator))
}

// newReconciler returns a new reconcile.Reconciler
func newReconciler(mgr manager.Manager, actuator Actuator) reconcile.Reconciler {
	return &ReconcileUptimeRobot{Client: mgr.GetClient(), scheme: mgr.GetScheme(), actuator: actuator}
}

// add adds a new Controller to mgr with r as the reconcile.Reconciler
func add(mgr manager.Manager, r reconcile.Reconciler) error {
	// Create a new controller
	c, err := controller.New("uptimerobot-controller", mgr, controller.Options{Reconciler: r})
	if err != nil {
		return err
	}

	// Watch for changes to UptimeRobot
	err = c.Watch(&source.Kind{Type: &monitorsv1.UptimeRobot{}}, &handler.EnqueueRequestForObject{})
	if err != nil {
		return err
	}

	return nil
}

var _ reconcile.Reconciler = &ReconcileUptimeRobot{}

// ReconcileUptimeRobot reconciles a UptimeRobot object
type ReconcileUptimeRobot struct {
	client.Client
	scheme   *runtime.Scheme
	actuator Actuator
}

// Reconcile reads that state of the cluster for a UptimeRobot object and makes changes based on the state
// Automatically generate RBAC rules to allow the Controller to read and write Deployments
// +kubebuilder:rbac:groups=monitors.tks.sh,resources=uptimerobots,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=monitors.tks.sh,resources=uptimerobots/status,verbs=get;update;patch
func (r *ReconcileUptimeRobot) Reconcile(request reconcile.Request) (reconcile.Result, error) {
	// Fetch the UptimeRobot instance
	instance := &monitorsv1.UptimeRobot{}
	err := r.Get(context.TODO(), request.NamespacedName, instance)
	if err != nil {
		if errors.IsNotFound(err) {
			// Object not found, return.  Created objects are automatically garbage collected.
			// For additional cleanup logic use finalizers.
			return reconcile.Result{}, nil
		}
		// Error reading the object - requeue the request.
		return reconcile.Result{}, err
	}

	name := instance.Name
	log.Info("Running reconcile Uptimerobot", "name", name)

	// If object hasn't been deleted and doesn't have a finalizer, add one
	// Add a finalizer to newly created objects.
	if instance.ObjectMeta.DeletionTimestamp.IsZero() &&
		!util.Contains(instance.ObjectMeta.Finalizers, monitorsv1.UptimeRobotFinalizer) {
		instance.Finalizers = append(instance.Finalizers, monitorsv1.UptimeRobotFinalizer)
		if err = r.Update(context.Background(), instance); err != nil {
			log.Info("failed to add finalizer to uptimerobot", "name", name, "err", err)
			return reconcile.Result{}, err
		}

		// Since adding the finalizer updates the object return to avoid later update issues
		return reconcile.Result{}, nil
	}

	if !instance.ObjectMeta.DeletionTimestamp.IsZero() {
		// no-op if finalizer has been removed.
		if !util.Contains(instance.ObjectMeta.Finalizers, monitorsv1.UptimeRobotFinalizer) {
			log.Info("reconciling uptimerobot object causes a no-op as there is no finalizer.", "name", name)
			return reconcile.Result{}, nil
		}

		log.Info("reconciling uptimerobot object %v triggers delete.", "name", name)
		if err := r.actuator.Delete(instance); err != nil {
			log.Error(err, "Error deleting uptimerobot object", "name", name)
			return reconcile.Result{}, err
		}
		// Remove finalizer on successful deletion.
		log.Info("uptimerobot object deletion successful, removing finalizer.", "name", name)
		instance.ObjectMeta.Finalizers = util.Filter(instance.ObjectMeta.Finalizers, monitorsv1.UptimeRobotFinalizer)
		if err := r.Client.Update(context.Background(), instance); err != nil {
			log.Error(err, "Error removing finalizer from uptimerobot object", "name", name)
			return reconcile.Result{}, err
		}
		return reconcile.Result{}, nil
	}

	log.Info("reconciling uptimerobot object triggers idempotent reconcile.", "name", name)
	err = r.actuator.Reconcile(instance)
	if err != nil {
		if requeueErr, ok := err.(*controllerError.RequeueAfterError); ok {
			log.Info("Actuator returned requeue after error", "requeueErr", requeueErr)
			return reconcile.Result{Requeue: true, RequeueAfter: requeueErr.RequeueAfter}, nil
		}
		log.Error(err, "Error reconciling uptimerobot object", "name", name)
		return reconcile.Result{}, err
	}

	// @TODO: When a monitor is removed from the list, delete it
	// @TODO: When a monitor is removed from the list, remove it from status page

	return reconcile.Result{}, nil
}
