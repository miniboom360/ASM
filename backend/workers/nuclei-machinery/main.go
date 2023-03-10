package main

import (
	"common"
	"github.com/urfave/cli"
	"log"
	"nuclei-machinery/handler"
	"os"
)

func main() {
	var queueName = "nuclei"
	var mTasks = map[string]interface{}{
		"NucleiTagsScan": handler.NucleiScan,
	}

	var workerApp *cli.App
	workerApp = cli.NewApp()
	workerApp.Name = "asm nuclei"
	workerApp.Usage = "nuclei machinery worker"
	workerApp.Version = "0.0.0"

	log.Println("work app is:", workerApp.Name)
	// Set the CLI app commands
	workerApp.Commands = []cli.Command{
		{
			Name:  "worker",
			Usage: "launch machinery worker",
			Action: func(c *cli.Context) error {
				if err := common.MchmultipleWorker(queueName, mTasks); err != nil {
					return cli.NewExitError(err.Error(), 1)
				}
				return nil
			},
		},
	}

	// Run the CLI app
	_ = workerApp.Run(os.Args)
}
