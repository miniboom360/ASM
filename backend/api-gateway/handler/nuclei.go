package handler

import (
	"api-gateway/module"
	"common"
	"encoding/json"
	"github.com/RichardKnop/machinery/v2/tasks"
	"github.com/google/uuid"
)

func NucleiScanByTags(tags, task_id string, domains []string) (string, error) {
	if task_id == "" {
		task_id = uuid.New().String()
	}

	var task = tasks.Signature{
		Name: "NucleiTagsScan",
		Args: []tasks.Arg{
			{
				Name:  "domains",
				Type:  "[]string",
				Value: domains,
			},
			{
				Name:  "tags",
				Type:  "string",
				Value: tags,
			},
		},
	}
	content, err := common.MchClient("nuclei", task, false)
	if err != nil {
		return "", err
	}

	var vulns []*common.Nucleivulns
	err = json.Unmarshal(content, &vulns)
	if err != nil {
		return "", err
	}

	// loc, _ := time.LoadLocation("Asia/Shanghai")
	// time := time.Now().In(loc).String()
	for _, v := range vulns {
		v.TaskId = task_id
	}
	// 写入module
	if err := module.AddNucleiItems(vulns); err != nil {
		return "", err
	}
	return task_id, nil
}
