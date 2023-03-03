package activitys

import (
	"backend/app"
	"context"
	"fmt"
	"testing"
)

func TestHttpX(t *testing.T) {
	targets := make([]string, 0)
	targets = append(targets, "106.75.13.27", "remote.cloudpnr.com")
	hps := app.HttpxReq{
		Targets:    targets,
		ThreadsNum: 1,
	}

	hrs, err := HttpxScan(context.Background(), hps)
	if err != nil {
		fmt.Printf("有错误 %+v\n", err)
		return
	}

	for _, hr := range hrs {
		fmt.Printf("返回数据结果:%+v\n", hr)
	}
}
