package tests

import (
	"nuclei-machinery/handler"
	"testing"
)

func TestNucleiScan(t *testing.T) {
	ss := make([]string, 0)
	// find vulu
	ss = append(ss, "http://106.75.13.27:8080")
	// can't find vuln
	//ss = append(ss, "https://baidu.com")
	//是不是没有结果，就是没有相对应的文件
	handler.NucleiScan(ss, "thinkphp")
}
