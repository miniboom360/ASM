package handler

import (
	"api-gateway/module"
	"common"
	"fmt"
	"github.com/google/uuid"
	"github.com/robfig/cron/v3"
	"log"
)

var (
	c *cron.Cron
)

// plan values: now#everyday#8:00
// everyweek#8:00
func AddTask(org_name, plan, scan_policy string, domains []string) (string, error) {
	task_id := uuid.New().String()
	item := new(common.TaskItem)
	item.OrganizationName = org_name
	item.Domains = domains
	item.ScanPolice = scan_policy
	eid, err := addTask()
	if err != nil {
		return "", err
	}
	item.EntryId = eid
	item.Staus = "进行中"
	item.TaskId = task_id

	if err := module.AddTaskItem(item); err != nil {
		return "", err
	}
	return task_id, nil
}

// wait for learning more about time period rule in cron.Cron.
func addTask() (int, error) {
	eid, err := c.AddFunc("@every 1s", func() {
		fmt.Println("tick every 1 second")
	})
	if err != nil {
		return -1, err
	}
	return int(eid), nil
}

// cancel some task is useless
func CancelTask(eid int) {
	c.Remove(cron.EntryID(eid))
}

// go开个协程再挂起
func InitScheduler() {
	log.Println("Starting...")
	chan1 := make(chan int)
	chan2 := make(chan int)
	c = cron.New()
	c.Start()

	select {
	// chan虽然准备好了接收操作，但是由于无数据写入，所以次case处于阻塞
	case <-chan1:
		fmt.Println("chan1 ready.")
	// chan2同1一样，也处于阻塞
	case <-chan2:
		fmt.Println("chan2 ready.")
	}
}
