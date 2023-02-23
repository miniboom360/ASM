package handler

import (
	"api-gateway/module"
	"common"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/robfig/cron/v3"
	"log"
	"strings"
)

const MaxTaskCount = 10000

var (
	c            *cron.Cron
	taskRunQueue chan *TaskReq
)

type TaskReq struct {
	ScanOption ScanOptions `json:"scan_option"`
	Plan       string      `json:"plan"`
	Org_name   string      `json:"org_name"`
	Domains    []string    `json:"domains"`
	ScanPolicy string      `json:"scan_policy"`
	TaskId     string      `json:"task_id"`
}

type ScanOptions struct {
	// this is nessnary, even though not choose this option, it still work
	Subdomain  bool `json:"subdomain"`
	NucleiScan bool `json:"nuclei_scan"`
}

// plan values: now#everyday#8:00
// everyweek#8:00
// todo:wait for record the cron in redis, make it still work when the programer is reboot
// todo:more choose about time
func AddTask(req *TaskReq) (string, error) {
	result, err := module.CheckWrongTask(req.Org_name, req.Domains)
	if err != nil {
		return "", err
	}
	if !result {
		return "", errors.New("there are the same domain in different task")
	}

	task_id := uuid.New().String()
	req.TaskId = task_id
	// 写入数据库
	items := make([]*common.TaskItem, 0)
	item := new(common.TaskItem)
	item.OrganizationName = req.Org_name
	item.Domains = req.Domains
	item.ScanPolice = req.ScanPolicy
	item.Period = req.Plan
	eid, err := addTask(req.Plan)
	if err != nil {
		return "", err
	}
	item.EntryId = eid
	item.Staus = "processing"
	item.TaskId = task_id
	items = append(items, item)
	if err := module.AddTaskItem(items); err != nil {
		panic(err)
		return "", err
	}
	// 把数据发给任务管理器运行
	SendRunTask(req)
	return task_id, nil
}

// wait for learning more about time period rule in cron.Cron.
// plan value : daily#default
func addTask(plan string) (int, error) {
	ps := strings.SplitN(plan, "#", -1)
	if len(ps) < 2 {
		return 0, errors.New("please enter the right params")
	}

	spec := "@every"
	if ps[0] == "daily" {
		spec += " 24h"
	} else if ps[0] == "week" {
		spec += " 168h"
	}

	switch ps[1] {
	case "default":
	default:
		break

	}

	// no matter how often it is executed, it must be executed now
	eid, err := c.AddFunc(spec, func() {
		fmt.Println(spec)
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

func SendRunTask(tr *TaskReq) {
	taskRunQueue <- tr
}

// go开个协程再挂起
// 在这里进行任务模块，将解析任务发到chan里，后台进行解析
func InitScheduler() {
	// 检查任务是否完成
	log.Println("Starting TaskManager...")
	taskRunQueue = make(chan *TaskReq, MaxTaskCount)
	chan1 := make(chan int)
	chan2 := make(chan int)
	c = cron.New()
	c.Start()

	for {
		select {
		case trq := <-taskRunQueue:
			// some times, MarshTaskOpt not be executed, although some message had send to taskRunQueue Already
			MarshTaskOpt(trq)
		// chan虽然准备好了接收操作，但是由于无数据写入，所以次case处于阻塞
		case <-chan1:
			fmt.Println("chan1 ready.")
		// chan2同1一样，也处于阻塞
		case <-chan2:
			fmt.Println("chan2 ready.")
		}
	}

}

/*
	task type:default ...
	Customize the process logic, or just fill and
	display the data according to the configuration?

	I prefect to use the way of data gather
*/
// func Default

func MarshTaskOpt(tr *TaskReq) {
	if tr.ScanOption.Subdomain {
		for _, v := range tr.Domains {
			fmt.Printf("已进入MarshTaskOpt %v\n", v)

			// TODO：假如所有的任务都完成了，要用什么逻辑能够验证呢
			// 比如这个任务包含了子域名和nuclei功能，你怎么知道何时能够set task complete呢？
			// 而且任务之间有串联的情况，子域名结束之后，再进行扫描你怎么做呢？
			// 目前最简单的，可能就是发现了有新增对象，就填入nuclei进行扫描，是异步进行的。
			// https://github.com/gocraft/work使用这个来看看，可以让任务paused、start、delete and display process
			go FindSubDomain(v, tr.TaskId)
		}
	}
}

// 检查任务状态
