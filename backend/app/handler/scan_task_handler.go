package handler

import (
	"backend/app"
	"backend/app/module"
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
	go handlerScanLogic(workflowID, req)
	g.String(http.StatusOK, fmt.Sprintf(`{"taskid":"%s"}`, workflowID))
}

// todo:添加数据写入功能，以及加入扫描逻辑
func handlerScanLogic(workflowID string, req ScanTask) {
	c, err := client.Dial(client.Options{HostPort: "106.75.13.27:7233"})
	if err != nil {
		log.Fatalln("unable to create Temporal client", err)
	}
	//defer c.Close()

	options := client.StartWorkflowOptions{
		ID:        workflowID,
		TaskQueue: app.ScanTaskQueue,
	}

	// name := "World"
	we, err := c.ExecuteWorkflow(context.Background(), options, workflows.ScanTaskWorkFlow, req.Domains)
	if err != nil {
		log.Fatalln("unable to complete Workflow", err)
	}

	result := make([]*app.SubdomainS, 0)

	err = we.Get(context.Background(), &result)
	if err != nil {
		log.Fatalln("unable to get Workflow result", err)
	}
	//判断是否是c被defer造成的panic，但是为什么写入数据库的时候就出现这个问题？
	//还有写入数据库，对这个c好像并没有要求呀
	// 将数据写入mysql
	module.AddSubDomainItems(result)
	//printResults(result, we.GetID(), we.GetRunID())
}
