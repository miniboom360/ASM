package main

import (
	"common"
	"encoding/json"
	"fmt"
	"github.com/RichardKnop/machinery/v2/tasks"
	"log"
)

func main() {
	var task = tasks.Signature{
		Name: "Dir",
		Args: []tasks.Arg{
			{
				Name:  "domain",
				Type:  "string",
				Value: "https://buffered.io",
			},
		},
	}
	content, err := common.MchClient("Gobuster", task, false)
	if err != nil {
		log.Fatal(err)
		return
	}
	var item common.DirItems
	err = json.Unmarshal(content, &item)
	if err != nil {
		log.Fatal(err)
		return
	}
	for _, v := range item.Dir {
		fmt.Printf("%s code is %s\n", v.Dir, v.Code)
	}
	//fmt.Println(item)
}
