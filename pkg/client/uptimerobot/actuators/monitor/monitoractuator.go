package monitor

import (
	"net/http"
	"net/url"
	"strconv"
	"strings"

	monitorv1 "github.com/cloud104/tks-uptimerobot-controller/pkg/apis/monitors/v1"
	uptimerobot "github.com/cloud104/uptimerobot-api/pkg/api"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/client-go/tools/record"
	logf "sigs.k8s.io/controller-runtime/pkg/runtime/log"
)

var log = logf.Log.WithName("monitor-actuator")

// Actuator ...
type Actuator struct {
	recorder record.EventRecorder
	apiKey   string
}

// NewActuator creates a new uptime robot actuator
func NewActuator(apiKey string, recorder record.EventRecorder) (*Actuator, error) {
	return &Actuator{
		recorder: recorder,
		apiKey:   apiKey,
	}, nil
}

// Reconcile will create or update the uptimerobot monitor
func (a *Actuator) Reconcile(monitor *monitorv1.UptimeRobot) error {
	// Emitt event
	a.recorder.Event(monitor, corev1.EventTypeWarning, "tks-uptimerobot", "monitoractuator Reconcile invoked")

	for _, h := range monitor.Spec.Hosts {
		err := reconcile(&h, monitor, a.apiKey)
		if err != nil {
			return err
		}
	}

	return nil
}

func reconcile(host *monitorv1.UptimeRobotHosts, monitor *monitorv1.UptimeRobot, apiKey string) (err error) {
	// @TODO: Where to put this
	// Make client
	uptimerobotClient := uptimerobot.NewClient(apiKey, &http.Client{})

	// Create Params
	params := url.Values{}

	// Set the search value
	params.Set("search", host.URL)

	// Get Monitors
	r, err := uptimerobotClient.GetMonitors(params)
	if err != nil {
		return err
	}
	// Cast Content
	resp := r.Content().(uptimerobot.UptimeRobotResponse)

	// Set values to create or update
	params = url.Values{}
	// @TODO: Remove url to se a bad formated error
	params.Set("url", host.URL)
	params.Set("friendly_name", host.FriendlyName)
	// @TODO: Do something for type
	params.Set("type", "1")

	//Setting AlertContacts
	//TODO TEST: https://play.golang.org/p/scrV3OHtNXd
	var alertContacts []string
	for _, c := range monitor.Spec.AlertContacts {
		alertContacts = append(alertContacts, strings.Join([]string{c.ID, c.Threshold, c.Recurrence}, "_"))
	}
	if len(alertContacts) > 0 {
		params.Set("alert_contacts", strings.Join(alertContacts, "-"))
	}

	// Reconcile
	// Oh god this piece of code
	log.Info("reconciling", "params", params)
	log.Info("monitors count", "monitors", len(resp.Monitors))
	if len(resp.Monitors) == 1 {
		m := resp.Monitors[0]
		params.Set("id", strconv.Itoa(m.Id))
		log.Info("updating monitor")
		r, err = uptimerobotClient.EditMonitor(params)
	} else {
		log.Info("creating monitor")
		r, err = uptimerobotClient.NewMonitor(params)
	}

	if err != nil {
		// @TODO: Get clear error here
		log.Error(err, "reconciliation failed")
		return err
	}
	resp = r.Content().(uptimerobot.UptimeRobotResponse)
	log.Info("reconciliation done", "resp", resp)

	err = reconcileStatusPage(apiKey, strconv.Itoa(resp.Monitor.Id), monitor)
	if err != nil {
		log.Error(err, "reconciliation status page failed")
		return
	}

	log.Info("reconciliation status done")
	return
}

func reconcileStatusPage(apiKey string, monitorID string, monitor *monitorv1.UptimeRobot) (err error) {
	uptimerobotClient := uptimerobot.NewClient(apiKey, &http.Client{})

	// Search for a existing page
	statusPage, err := uptimerobotClient.SearchStatusPage(monitor.Spec.StatusPage.FriendlyName)
	if err != nil {
		return
	}

	// Build list of monitor ids exitent + new
	var monitorIds []string

	// Add new monitorID
	monitorIds = append(monitorIds, monitorID)

	// Add the ones that exist
	for _, v := range statusPage.MonitorList() {
		monitorIds = appendIfMissing(monitorIds, strconv.Itoa(v))
	}

	// Build params
	params := make(url.Values)
	params.Set("custom_domain", monitor.Spec.StatusPage.URL)
	params.Set("friendly_name", monitor.Spec.StatusPage.FriendlyName)
	params.Set("monitors", strings.Join(monitorIds, "-"))

	if statusPage.FriendlyName != "" {
		log.Info("updating status page")
		params.Set("id", strconv.Itoa(statusPage.Id))
		_, err = uptimerobotClient.EditPublicStatusPage(params)
	} else {
		log.Info("creating status page")
		_, err = uptimerobotClient.NewPublicStatusPage(params)
	}

	return
}

// Delete deletes a uptimerobot monitor
func (a *Actuator) Delete(monitor *monitorv1.UptimeRobot) error {
	a.recorder.Event(monitor, corev1.EventTypeWarning, "tks-uptimerobot", "monitoractuator Delete invoked")

	for _, h := range monitor.Spec.Hosts {
		err := delete(&h, monitor, a.apiKey)
		if err != nil {
			return err
		}
	}
	return nil
}

func delete(host *monitorv1.UptimeRobotHosts, monitor *monitorv1.UptimeRobot, apiKey string) (err error) {
	// @TODO: Where to put this
	// Make client
	uptimerobotClient := uptimerobot.NewClient(apiKey, &http.Client{})

	// Create Search Params
	params := url.Values{}
	params.Set("search", host.URL)

	// Get Monitors
	r, err := uptimerobotClient.GetMonitors(params)
	if err != nil {
		return
	}
	monitorsLen := len(r.Content().(uptimerobot.UptimeRobotResponse).Monitors)

	// Do nothing if no monitors
	if monitorsLen != 1 {
		log.Info("quantity of monitors not 1, skiping", "monitorsLen", monitorsLen)
		return
	}

	// Cast Monitor
	m := r.Content().(uptimerobot.UptimeRobotResponse).Monitors[0]

	// Set values to create or update
	params = url.Values{}
	params.Set("id", strconv.Itoa(m.Id))

	// Reconcile
	// Oh god this piece of code
	log.Info("deleting monitor")
	_, err = uptimerobotClient.DeleteMonitor(params)
	if err != nil {
		log.Error(err, "deletion failed")
		return
	}

	// Delete status page
	err = deleteStatusPage(apiKey, strconv.Itoa(m.Id), monitor)
	if err != nil {
		log.Error(err, "status page deletion failed")
		return
	}

	log.Info("deletion done")
	return
}

func deleteStatusPage(apiKey string, monitorID string, monitor *monitorv1.UptimeRobot) (err error) {
	uptimerobotClient := uptimerobot.NewClient(apiKey, &http.Client{})

	// Search for a existing page
	statusPage, err := uptimerobotClient.SearchStatusPage(monitor.Spec.StatusPage.FriendlyName)
	if err != nil {
		return err
	}

	// Build params
	params := make(url.Values)

	if statusPage.FriendlyName != "" {
		log.Info("updating status page")
		params.Set("id", strconv.Itoa(statusPage.Id))
		_, err = uptimerobotClient.DeletePublicStatusPage(params)
	}

	return
}

// UTILS (God forgive this bad people)

func appendIfMissing(slice []string, i string) []string {
	for _, a := range slice {
		if a == i {
			return slice
		}
	}
	return append(slice, i)
}
