package monitor_test

import (
	"testing"

	monitorv1 "github.com/cloud104/tks-uptimerobot-controller/pkg/apis/monitors/v1"
	monitoractuator "github.com/cloud104/tks-uptimerobot-controller/pkg/client/uptimerobot/actuators/monitor"
	mocks "github.com/cloud104/tks-uptimerobot-controller/pkg/client/uptimerobot/actuators/monitor/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	corev1 "k8s.io/api/core/v1"
)

//go:generate mockgen -package=mocks -destination=mocks/monitoractuator_mock.go k8s.io/client-go/tools/record EventRecorder
func TestRencocile(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	key := ""
	r := mocks.NewMockEventRecorder(ctrl)

	// Stubs
	monitor := &monitorv1.UptimeRobot{}

	r.EXPECT().
		Event(
			gomock.Eq(monitor),
			gomock.Eq(corev1.EventTypeWarning),
			gomock.Eq("tks-uptimerobot"),
			gomock.Eq("monitoractuator Reconcile invoked"),
		)

	actuator, err := monitoractuator.NewActuator(key, r)
	assert.Nil(t, err)

	// @TODO
	actuator.Reconcile(monitor)
}
