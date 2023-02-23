package workflows

import (
  "backend/app/activitys"
  "go.temporal.io/sdk/workflow"
  "time"
)

func ScanTaskWorkFlow(ctx workflow.Context, domains []string) (string, error) {
  options := workflow.ActivityOptions{
    // 这是什么意思？StartToCloseTimeout意思是5秒之内就timeout就关？
    StartToCloseTimeout: time.Minute * 10,
  }

  ctx = workflow.WithActivityOptions(ctx, options)

  var result string
  err := workflow.ExecuteActivity(ctx, activitys.SearchSubDomain, domains).Get(ctx, &result)
  return result, err

}
