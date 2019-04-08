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
	"github.com/k0kubun/pp"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	logf "sigs.k8s.io/controller-runtime/pkg/runtime/log"
	"sigs.k8s.io/controller-runtime/pkg/source"
)

var log = logf.Log.WithName("controller")

/**
* USER ACTION REQUIRED: This is a scaffold file intended for the user to modify with their own Controller
* business logic.  Delete these comments after modifying this file.*
 */

// Add creates a new UptimeRobot Controller and adds it to the Manager with default RBAC. The Manager will set fields on the Controller
// and Start it when the Manager is Started.
func Add(mgr manager.Manager) error {
	return add(mgr, newReconciler(mgr))
}

// newReconciler returns a new reconcile.Reconciler
func newReconciler(mgr manager.Manager) reconcile.Reconciler {
	return &ReconcileUptimeRobot{Client: mgr.GetClient(), scheme: mgr.GetScheme()}
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
	scheme *runtime.Scheme
}

// Reconcile reads that state of the cluster for a UptimeRobot object and makes changes based on the state read
// and what is in the UptimeRobot.Spec
// TODO(user): Modify this Reconcile function to implement your Controller logic.  The scaffolding writes
// a Deployment as an example
// Automatically generate RBAC rules to allow the Controller to read and write Deployments
// +kubebuilder:rbac:groups=apps,resources=deployments,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=apps,resources=deployments/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=monitors.tks.sh,resources=uptimerobots,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=monitors.tks.sh,resources=uptimerobots/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=apps,resources=secrets,verbs=get
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

	// @TODO: Invoke someone here to post to a external source that can be queryed
	pp.Println(instance)

	// @TODO: Where to put this
	// httpClient := &http.Client{}
	// client := uptimerobot.NewClient("yourAPIKeyHere", httpClient)
	// resp, err := client.GetMonitors()
	// if err != nil {
	// 	return reconcile.Result{}, err
	// }
	// pp.Println(resp)

	return reconcile.Result{}, nil
}
