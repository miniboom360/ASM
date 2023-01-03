package handler

import (
	"fmt"
	"io/ioutil"
	"log"
	"os/exec"
	"time"
)

const (
	// 测试环境
	test_pythonPath          = "/usr/local/Homebrew/bin/python3.10"
	test_oneforallPath       = "/Users/liyang/tools/asm/asm-demo/workers/oneforall-machery/scripts/python/oneforall.py"
	test_oneforallresultPath = "/Users/liyang/tools/asm/asm-demo/workers/oneforall-machery/scripts/python/results"

	test_resultfile = "/Users/liyang/tools/asm/asm-demo/workers/oneforall-machery/adapay2022.json"

	pythonPath    = "python3"
	oneforallPath = "/python/oneforall.py"
	// resultfile          = "/adapay2022.json"
	oneforallresultPath = "/oneforall-machery/scripts/python/results"
)

type OneforAllItem struct {
	ID        int    `json:"id"`
	Alive     int    `json:"alive"`
	Request   int    `json:"request"`
	Resolve   int    `json:"resolve"`
	URL       string `json:"url"`
	Subdomain string `json:"subdomain"`
	Level     int    `json:"level"`
	Cname     string `json:"cname"`
	IP        string `json:"ip"`
	// Public    int    `json:"public"`
	Cdn    int    `json:"cdn"`
	Port   int    `json:"port"`
	Status int    `json:"status"`
	Reason string `json:"reason"`
	Title  string `json:"title"`
	Banner string `json:"banner"`
	Cidr   string `json:"cidr"`
	Asn    string `json:"asn"`
	Org    string `json:"org"`
	Addr   string `json:"addr"`
	Isp    string `json:"isp"`
	Source string `json:"source"`
}

func GetOneSubDomain(domain string) ([]byte, error) {
	log.Printf("received search subdomains request:%s", domain)

	// out, err := exec.Command(test_pythonPath, test_oneforallPath, "--target", domain,
	// 	"--fmt", "json", "--path", test_resultfile, "run").Output()

	resultfile := fmt.Sprintf("/%s_%s.json", domain, time.Now().Format("2006-01-02 15:04:05"))
	out, err := exec.Command(pythonPath, oneforallPath, "--target", domain,
		"--fmt", "json", "--path", resultfile, "run").Output()
	if err != nil {
		fmt.Printf("%s", err)
		return nil, err
	} else {
		fmt.Println("Command Successfully Executed")
		output := string(out[:])
		fmt.Println(output)
	}

	// 删除oneforall的多余日志和文件
	out1, err := exec.Command("rm", "-rf", oneforallresultPath).Output()
	// out1, err := exec.Command("rm", "-rf", test_oneforallresultPath).Output()
	if err != nil {
		fmt.Printf("%s", err)
		return nil, err
	} else {
		fmt.Println("删除文件命令执行成功")
		output := string(out1[:])
		fmt.Println(output)
	}

	// 结果文件

	// content, err := ioutil.ReadFile(test_resultfile)
	content, err := ioutil.ReadFile(resultfile)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	// 删除这个文件
	// 删除oneforall的多余日志和文件
	// out2, err := exec.Command("rm", "-rf", test_resultfile).Output()
	out2, err := exec.Command("rm", "-rf", resultfile).Output()
	if err != nil {
		fmt.Printf("%s", err)
		return nil, err
	} else {
		fmt.Println("删除文件命令执行成功")
		output := string(out2[:])
		fmt.Println(output)
	}
	return content, nil
}
