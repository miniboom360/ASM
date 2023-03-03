package workflows

import (
  "backend/app"
  "backend/app/activitys"
  "go.temporal.io/sdk/workflow"
  "time"
)

func ScanTaskWorkFlow(ctx workflow.Context, sti app.ScanTaskItem) ([]*app.SubdomainS, error) {
  options := workflow.ActivityOptions{
    // 这是什么意思？StartToCloseTimeout意思是5秒之内就timeout就关？
    StartToCloseTimeout: time.Minute * 10,
  }

  ctx = workflow.WithActivityOptions(ctx, options)

  subs := make([]*app.SubdomainS, 0)
  err := workflow.ExecuteActivity(ctx, activitys.SearchSubDomain, sti).Get(ctx, &result)

  // portscan
  if sti.ScanOption.PortScan {
    // 调用端口扫描
    for _, subdomain := range subs {
      // subdomain.

    }
  }

  return result, err

}
