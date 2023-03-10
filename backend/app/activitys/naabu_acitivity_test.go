package activitys

import (
	"backend/app"
	"context"
	"fmt"
	"testing"
)

func TestNaabu(t *testing.T) {
	targets := make([]string, 0)
	targets = append(targets, "106.75.13.27", "remote.cloudpnr.com")
	aps := app.PortScanReq{
		Tag:     "top-1000",
		Targets: targets,
	}

	nds, err := NaabuScan(context.Background(), aps)
	if err != nil {
		fmt.Printf("有错误 %+v\n", err)
		return
	}

	for _, nd := range nds {
		fmt.Printf("返回数据结果:%+v\n", nd)
	}
}
