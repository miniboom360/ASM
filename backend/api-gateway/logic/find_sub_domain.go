package logic

import (
	"api-gateway/module"
	"common"
	"encoding/json"
	"fmt"
	"github.com/RichardKnop/machinery/v2/tasks"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"log"
	"net/http"
	"time"
)

func Find_sub_domain(c *gin.Context) {
	// *HostScanResult
	domain := c.Query("domain")
	if domain == "" {
		c.String(http.StatusOK, "域名不能为空")
		return
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
		log.Fatal(err)
		return
	}
	var item []*common.SubDomainItems

	err = json.Unmarshal(content, &item)
	if err != nil {
		log.Fatal(err)
		return
	}

	task_id := uuid.New().String()

	loc, _ := time.LoadLocation("Asia/Shanghai")
	time := time.Now().In(loc).String()
	for _, d := range item {
		d.FirstTime = time
		d.TaskId = task_id
		d.Domain = domain
	}
	//写入module
	if err := module.AddSubDomainItems(item); err != nil {
		c.JSON(http.StatusConflict, `{"result":"存储数据库错误"}`)
		return
	}

	r := fmt.Sprintf(`{"task_id":%s}`, task_id)
	c.JSON(http.StatusOK, r)
	return
}
