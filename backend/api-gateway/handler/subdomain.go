package handler

import (
	"api-gateway/module"
	"common"
	"encoding/json"
	"fmt"
	"github.com/RichardKnop/machinery/v2/tasks"
	"github.com/google/uuid"
	"time"
)

func FindSubDomain(domain, taskid string) (string, error) {

	if taskid != "" {
		common.SetTaskStatusByTaskId(taskid, "SUBDOMAIN_SCANNING")
	}

	var task = tasks.Signature{
		Name: "ScanOneDomain",
		Args: []tasks.Arg{
			{
				Name:  "domain",
				Type:  "string",
				Value: domain,
			},
		},
	}
	content, err := common.MchClient("oneforall", task, false)
	if err != nil {
		return "", err
	}
	var item []*common.Subdomains

	err = json.Unmarshal(content, &item)
	if err != nil {
		panic(err)
		return "", err
	}

	for _, v := range item {
		if v.UId == "" {
			v.UId = uuid.New().String()
		}
	}

	if taskid == "" {
		taskid = uuid.New().String()
	}

	loc, _ := time.LoadLocation("Asia/Shanghai")
	time := time.Now().In(loc).String()
	for _, d := range item {
		d.FirstTime = time
		d.TaskId = taskid
		d.Domain = domain
	}
	if err := module.AddSubDomainItems(item); err != nil {
		panic(err)
		return "", err
	}

	common.SetTaskStatusByTaskId(taskid, "SUBDOMAIN_COMPLETE")
	fmt.Println("子域名扫描结束\n")
	return taskid, nil
}
