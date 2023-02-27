package workflows

import (
	"backend/app"
	"backend/app/activitys"
	"go.temporal.io/sdk/workflow"
	"time"
)

// start处给到workid，剩下的调用多个activity都没有关系，都是一个workflow
func ScanTaskWorkFlow(ctx workflow.Context, sti app.ScanTaskItem) ([]*app.SubdomainS, error) {
	options := workflow.ActivityOptions{
		// 这是什么意思？StartToCloseTimeout意思是5秒之内就timeout就关？
		StartToCloseTimeout: time.Minute * 10,
	}

	ctx = workflow.WithActivityOptions(ctx, options)

	result := make([]*app.SubdomainS, 0)
	err := workflow.ExecuteActivity(ctx, activitys.SearchSubDomain, sti).Get(ctx, &result)

	// test param tags is info leak

	return result, err

}
