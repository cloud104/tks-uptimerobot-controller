package uptimerobot

import (
	monitorv1 "github.com/cloud104/tks-uptimerobot-controller/pkg/apis/monitors/v1"
)

// Actuator controls uptimerobot api.
// All methods should be idempotent unless otherwise specified.
//go:generate mockgen -package=mocks -destination=mocks/actuator_mock.go -source=actuator.go Actuator
type Actuator interface {
	// Reconcile creates or applies updates to the cluster.
	Reconcile(*monitorv1.UptimeRobot) error
	// Delete the cluster.
	Delete(*monitorv1.UptimeRobot) error
}
