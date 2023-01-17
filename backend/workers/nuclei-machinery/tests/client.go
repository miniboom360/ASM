package main

import (
	"common"
	"encoding/json"
	"fmt"
	"github.com/RichardKnop/machinery/v2/tasks"
	"log"
)

func main() {
	ss := make([]string, 0)
	ss = append(ss, "http://106.75.13.27:8080")
	ss = append(ss, "http://106.75.13.27:80")
	var task = tasks.Signature{
		Name: "NucleiTagsScan",
		Args: []tasks.Arg{
			{
				Name:  "domains",
				Type:  "[]string",
				Value: ss,
			},
			{
				Name:  "tags",
				Type:  "string",
				Value: "thinkphp",
			},
		},
	}
	content, err := common.MchClient("nuclei", task, false)
	if err != nil {
		log.Fatal(err)
		return
	}

	var vulns []*common.NucleiVuln
	err = json.Unmarshal(content, &vulns)
	if err != nil {
		log.Fatal(err)
		return
	}
	for _, v := range vulns {
		fmt.Printf("vulns = %+v", v)
	}

}
