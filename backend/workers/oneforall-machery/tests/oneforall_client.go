package main

import (
	"common"
	"encoding/json"
	"fmt"
	"github.com/RichardKnop/machinery/v2/tasks"
	"log"
	"oneforall-machery/handler"
)

func main() {
	var task = tasks.Signature{
		Name: "ScanOneDomain",
		Args: []tasks.Arg{
			{
				Name:  "domain",
				Type:  "string",
				Value: "chinapnr.com",
			},
		},
	}
	content, err := common.MchClient("oneforall", task, false)
	if err != nil {
		log.Fatal(err)
		return
	}
	var item []handler.OneforAllItem

	err = json.Unmarshal(content, &item)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println(item)
}
