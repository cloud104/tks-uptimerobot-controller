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

	"github.com/urfave/cli/altsrc"
	cli "gopkg.in/urfave/cli.v1"

	// @TODO: Yep, what is this?
	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
)

var (
	version = "0.0.0"
)

func main() {
	app := cli.NewApp()
	app.Action = run
	app.Name = "UptimeRobot Operator"
	app.Version = version
	flags := createFlags()
	app.Before = altsrc.InitInputSourceWithContext(flags, altsrc.NewYamlSourceFromFlagFunc("config"))
	app.Flags = flags

	app.Run(os.Args)
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
