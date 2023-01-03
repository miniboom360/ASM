package main

import (
	"common"
	"github.com/urfave/cli"
	"gobuster/handler"
	"log"
	"os"
)

func main() {
	var queueName = "Gobuster"
	var mTasks = map[string]interface{}{
		"Dir": handler.GobusterDir,
	}

	var workerApp *cli.App
	workerApp = cli.NewApp()
	workerApp.Name = "asm Gobuster"
	workerApp.Usage = "Gobuster machinery worker"
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
	//handler_gobuster()
}
