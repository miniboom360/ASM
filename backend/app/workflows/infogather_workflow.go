package workflows

import (
	"backend/app"
	"backend/app/activitys"
	"go.temporal.io/sdk/workflow"
	"time"
)

func ScanTaskWorkFlow(ctx workflow.Context, sti app.ScanTaskItem) ([]*app.SubdomainS, error) {
	options := workflow.ActivityOptions{
		// 这是什么意思？StartToCloseTimeout意思是5秒之内就timeout就关？
		StartToCloseTimeout: time.Minute * 10,
	}

	ctx = workflow.WithActivityOptions(ctx, options)

	subs := make([]*app.SubdomainS, 0)
	err := workflow.ExecuteActivity(ctx, activitys.SearchSubDomain, sti).Get(ctx, &subs)
	if err != nil {
		return nil, err
	}
	// default模式开始
	// port scan && httpx

	// 目前是单个子域名，如何组织，这个数据结构该如何组织？
	for _, subdomain := range subs {
		subdomain.TaskId = sti.TaskId
		subdomain.Subdomains = make(map[string]*app.SubdomainItem, 0)
		// subdomain.Subdomains["1"].Nd = nil
		if subdomain.SubdomainsSclice == nil || len(subdomain.SubdomainsSclice) == 0 {
			continue
		}

		subdomain.SubdomainsSclice = append(subdomain.SubdomainsSclice, subdomain.MainDomain)

		// app.PortScanReq
		psq := app.PortScanReq{
			Targets: subdomain.SubdomainsSclice,
			Tag:     sti.ScanOpt.PortTag,
		}

		nds := make(map[string][]*app.NaabuData, 0)
		err := workflow.ExecuteActivity(ctx, activitys.NaabuScan, psq).Get(ctx, &nds)
		if err != nil {
			return nil, err
		}

		for t, nd := range nds {
			// subdomain.Subdomains[t]就是nil，怎么又能给它的成员赋值呢
			si := new(app.SubdomainItem)
			si.Nd = make([]*app.NaabuData, 0)
			si.Nd = nd
			subdomain.Subdomains[t] = si
			// subdomain.Subdomains[t].Nd = nd
		}

		// httpx
		hr := app.HttpxReq{
			Targets:    subdomain.SubdomainsSclice,
			ThreadsNum: sti.HttpxReq.ThreadsNum,
		}
		hxds := make([]*app.HttpXData, 0)
		err = workflow.ExecuteActivity(ctx, activitys.HttpxScan, hr).Get(ctx, &hxds)
		if err != nil {
			return nil, err
		}

		for _, hxd := range hxds {
			subdomain.Subdomains[hxd.Input].Hxd = hxd
			subdomain.Subdomains[hxd.Input].TaskId = sti.TaskId
			subdomain.Subdomains[hxd.Input].SubDomain = hxd.Input
		}

	}

	// TODO:portscan only
	// if sti.ScanOption.PortScanOnly {
	//  // 调用端口扫描，对每个子域的端口内容都进行数据填充
	//  for _, subdomain := range subs {
	//    // subdomain.
	//
	//  }
	// }

	return subs, err

}
