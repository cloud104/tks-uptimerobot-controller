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

package main

import (
	"os"

	cli "gopkg.in/urfave/cli.v1"
	altsrc "gopkg.in/urfave/cli.v1/altsrc"
	logf "sigs.k8s.io/controller-runtime/pkg/runtime/log"

	// @TODO: Yep, what is this?
	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
)

var (
	version = "0.0.0"
)

func main() {
	// Start logger
	logf.SetLogger(logf.ZapLogger(false))
	log := logf.Log.WithName("cli")

	app := cli.NewApp()
	app.Action = run
	app.Version = version
	flags := createFlags()
	app.Before = altsrc.InitInputSourceWithContext(flags, altsrc.NewYamlSourceFromFlagFunc("config"))
	app.Flags = flags

	err := app.Run(os.Args)
	if err != nil {
		log.Error(err, "cli boot failed")
	}
}

func run(c *cli.Context) {
	op := Operator{
		UptimeRobotKey: c.String("uptimerobot-api-key"),
		MetricsAddr:    c.String("metrics-addr"),
	}
	op.Exec()
}

func createFlags() []cli.Flag {
	return []cli.Flag{
		cli.StringFlag{
			Name:  "config, c",
			Usage: "Load configuration from `FILE`",
			Value: os.Getenv("HOME") + "/config.yaml",
		},
		altsrc.NewStringFlag(
			cli.StringFlag{
				Name:   "metrics-addr",
				Usage:  "Metrics Address",
				EnvVar: "METRICS_ADDR",
				Value:  ":8080",
			},
		),
		altsrc.NewStringFlag(
			cli.StringFlag{
				Name:   "uptimerobot-api-key",
				Usage:  "UptimeRobot Api Key",
				EnvVar: "UPTIMEROBOT_API_KEY",
			},
		),
	}
}
