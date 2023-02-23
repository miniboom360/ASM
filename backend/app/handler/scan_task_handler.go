package handler

import (
  "backend/app"
  "backend/app/workflows"
  "context"
  "fmt"
  "github.com/gin-gonic/gin"
  "go.temporal.io/sdk/client"
  "log"
  "net/http"
  "time"
)

type ScanTask struct {
  Domains  []string `json:"domains"`
  ScanName string   `json:"scan_name"`
}

func ScanTaskHandler(g *gin.Context) {

  req := ScanTask{}

  g.BindJSON(&req)
  if req.ScanName == "" {
    g.String(http.StatusOK, "扫描名不能为空")
    return
  }

  workflowID := req.ScanName + "-" + fmt.Sprintf("%d", time.Now().Unix())
  c, err := client.Dial(client.Options{HostPort: "106.75.13.27:7233"})
  if err != nil {
    log.Fatalln("unable to create Temporal client", err)
  }
  defer c.Close()

  options := client.StartWorkflowOptions{
    ID:        workflowID,
    TaskQueue: app.ScanTaskQueue,
  }

  // name := "World"
  we, err := c.ExecuteWorkflow(context.Background(), options, workflows.ScanTaskWorkFlow, req.Domains)
  if err != nil {
    log.Fatalln("unable to complete Workflow", err)
  }

  var greeting string
  err = we.Get(context.Background(), &greeting)
  if err != nil {
    log.Fatalln("unable to get Workflow result", err)
  }
  printResults(greeting, we.GetID(), we.GetRunID())

}

func printResults(greeting string, workflowID, runID string) {
  fmt.Printf("\nWorkflowID: %s RunID: %s\n", workflowID, runID)
  fmt.Printf("\n%s\n\n", greeting)
}
