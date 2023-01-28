package tests

import (
	"api-gateway/module"
	"fmt"
	"log"
	"testing"
)

func TestTask(t *testing.T) {

	err := module.InitMysql()
	if err != nil {
		log.Fatalln(err.Error())
		return
	}

	org_name1 := "汇付天下"
	domains1 := make([]string, 0)
	domains1 = append(domains1, "chinapnr.com")
	result, _ := module.CheckWrongTask(org_name1, domains1)
	if result {
		fmt.Printf("不存在错误的任务,组织名:%v\n", org_name1)
	}

	org_name2 := "汇来米2"
	domains2 := make([]string, 0)
	domains2 = append(domains2, "huilaimi.com")
	result1, _ := module.CheckWrongTask(org_name2, domains2)
	if result1 {
		fmt.Printf("不存在错误的任务,组织名:%v\n", org_name2)
	} else {
		fmt.Printf("存在错误的任务,组织名:%v\n", org_name2)
	}
}
