package workflows

import (
  "backend/app/activitys"
  "go.temporal.io/sdk/workflow"
  "time"
)

func GreetingWorkflow(ctx workflow.Context, name string) (string, error) {
  options := workflow.ActivityOptions{
    // 这是什么意思？StartToCloseTimeout意思是5秒之内就timeout就关？
    StartToCloseTimeout: time.Second * 5,
  }

  ctx = workflow.WithActivityOptions(ctx, options)

  var result string
  err := workflow.ExecuteActivity(ctx, activitys.ComposeGreeting, name).Get(ctx, &result)
  return result, err

}
