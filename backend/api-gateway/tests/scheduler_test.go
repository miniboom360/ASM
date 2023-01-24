package tests

import (
	"api-gateway/handler"
	"testing"
	"time"
)

func TestScheduler(t *testing.T) {
	//org_name, plan, scan_policy string, domains []string
	go handler.InitScheduler()
	time.Sleep(time.Second * 3)
	domains := make([]string, 0)
	domains = append(domains, "chinapnr.com", "adapay.com")
	handler.AddTask("huifu", "now#everyday#8:00", "default", domains)
	//handler.AddTask1("huifu", "now#everyday#8:00", "default", domains)
}
