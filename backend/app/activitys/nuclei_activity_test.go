package activitys

import (
	"backend/app"
	"context"
	"fmt"
	"testing"
)

func TestNuclei(t *testing.T) {
	domains := make([]string, 0)
	domains = append(domains, "172.31.13.198:8080")
	// 没有指定端口，就没有扫描出漏洞，所以还是得先爆破端口，不是吗
	// domains = append(domains, "172.31.13.198")
	nr := app.NucleiReq{
		Domains: domains,
		Tags:    "thinkphp",
	}
	r, err := NucleiScan(context.Background(), nr)
	if err != nil {
		panic(err)
	}
	for _, v := range r {
		fmt.Printf("response is %+v\n", v)
	}

}
