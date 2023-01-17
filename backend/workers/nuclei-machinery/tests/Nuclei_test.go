package tests

import (
	"common"
	"encoding/json"
	"fmt"
	"log"
	"nuclei-machinery/handler"
	"testing"
)

func TestNucleiScan(t *testing.T) {
	ss := make([]string, 0)
	ss = append(ss, "http://106.75.13.27:8080")
	ss = append(ss, "http://106.75.13.27:80")
	vs, err := handler.NucleiScan(ss, "thinkphp")
	vulns := make([]*common.Nucleivulns, 0)
	err = json.Unmarshal(vs, &vulns)
	if err != nil {
		log.Fatalln(err.Error())
	}
	fmt.Printf("test pass, vulns result is %+v", vulns)
}
