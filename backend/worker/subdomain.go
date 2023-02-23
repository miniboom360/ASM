package main

import (
  "backend/app"
  "backend/app/activitys"
  "backend/app/workflows"
  "go.temporal.io/sdk/client"
  "go.temporal.io/sdk/worker"
  "log"
)

func main() {
  c, err := client.Dial(client.Options{HostPort: "106.75.13.27:7233"})
  if err != nil {
    log.Fatalln("unable to create Temporal client", err)
  }
  defer c.Close()

  w := worker.New(c, app.ScanTaskQueue, worker.Options{})
  w.RegisterWorkflow(workflows.ScanTaskWorkFlow)
  w.RegisterActivity(activitys.SearchSubDomain)

  err = w.Run(worker.InterruptCh())
  if err != nil {
    log.Fatalln("unable to start Worker", err)
  }

}
