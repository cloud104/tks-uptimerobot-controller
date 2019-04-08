package options

import (
	"errors"

	"github.com/spf13/cobra"
)

// Singleton
var controllerOpts = &ControllerOptions{}

// ControllerOptions ...
type ControllerOptions struct {
	MetricsAddr    string
	UptimeRobotKey string
}

// GetControllerOptions ...
func GetControllerOptions() *ControllerOptions {
	return controllerOpts
}

// AddFlags ...
func (o *ControllerOptions) AddFlags(cmd *cobra.Command) {
	cmd.Flags().StringVar(&o.MetricsAddr, "metrics-addr", ":8080", "The address the metric endpoint binds to.")
	cmd.Flags().StringVar(&o.UptimeRobotKey, "uptime-robot-key", "", "The key to access uptime robot")
}

// Validate ...
func (o *ControllerOptions) Validate() error {
	if o.UptimeRobotKey == "" {
		return errors.New("uptime robot secret key not set")
	}

	return nil
}
