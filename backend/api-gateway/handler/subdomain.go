package handler

import (
	"api-gateway/module"
	"common"
	"encoding/json"
	"github.com/RichardKnop/machinery/v2/tasks"

	"github.com/google/uuid"
	"time"
)

func FindSubDomain(domain, taskid string) (string, error) {
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

	// 纠结于表的数据，现在这样做行不行？会不会有什么问题？其实是想最佳方案，但是现在想不到
	// 问题还是遇到了在解决在改变吧。
	err = json.Unmarshal(content, &item)
	if err != nil {
		return "", err
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
	// 写入module
	if err := module.AddSubDomainItems(item); err != nil {
		return "", err
	}
	return taskid, nil
}
