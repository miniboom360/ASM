package activitys

import (
	"bufio"
	"fmt"
	uuid2 "github.com/google/uuid"
	"log"
	"os"
	"os/exec"
)

const (
	NaabuPath = "naabu.exe"
	NmapPath  = "nmap"
)

func NaabuScan(targets []string, tag string) {
	targets_file := fmt.Sprintf("%s.txt", uuid2.New())
	//写入域名
	file, err := os.OpenFile(targets_file, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("文件打开失败", err)
	}
	//及时关闭file句柄
	defer file.Close()
	write := bufio.NewWriter(file)
	for _, t := range targets {
		write.WriteString(t + "\n")
	}
	//Flush将缓存的文件真正写入到文件中
	write.Flush()
	output_file := uuid2.New().String() + ".json"
	naabu_command := NaabuPath + fmt.Sprintf(" -list %s -json -o %s -exclude-cdn ", targets_file, output_file)

	if tag == "top-100" {
		naabu_command += " -top-ports 100 "
	}
	if tag == "full" {
		naabu_command += " -p -"
	}
	if tag == "top-1000" {
		naabu_command += " -top-ports 1000 "
	}
	//target_file := fmt.Sprintf("%s.xml", uuid2.New())
	//c := fmt.Sprintf(`echo %s | naabu -nmap-cli 'nmap -sV -oX nmap-output > '`, target, target_file)
	cmd := exec.Command("bash", "-c", naabu_command)
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("result =>", string(output))
}
