package handler

import (
  "api-gateway/module"
  "common"
  "encoding/json"
  "github.com/RichardKnop/machinery/v2/tasks"

  "github.com/google/uuid"
  "time"
)

func FindSubDomain(domain string) (string, error) {
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
    return "", err
  }

  task_id := uuid.New().String()

  loc, _ := time.LoadLocation("Asia/Shanghai")
  time := time.Now().In(loc).String()
  for _, d := range item {
    d.FirstTime = time
    d.TaskId = task_id
    d.Domain = domain
  }
  // 写入module
  if err := module.AddSubDomainItems(item); err != nil {
    return "", err
  }
  return task_id, nil
}
