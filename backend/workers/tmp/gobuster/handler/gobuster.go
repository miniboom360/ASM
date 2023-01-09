package handler

import (
	"bufio"
	"common"
	"encoding/json"
	"fmt"
	uuid2 "github.com/google/uuid"
	"io"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

const (
	wordlist_path = "wordlist.txt"
	gobuster_path = "./gobuster"

	wordlist_test_path = "D:\\code\\asm-demo\\backend\\workers\\gobuster\\scripts\\wordlist.txt"
	gobuster_test_path = "D:\\code\\asm-demo\\backend\\workers\\gobuster\\scripts\\gobuster.exe"
)

// ./gobuster dir -u https://huifu.com  -t 50 -w wordlist.txt -o china.txt
func GobusterDir(domain string) ([]byte, error) {
	uuid := uuid2.New()
	resultfile := fmt.Sprintf("%s.txt", uuid.String())
	out, err := exec.Command(gobuster_path, "dir", "-u", domain,
		"-t", "50", "-w", wordlist_path, "-o", resultfile).Output()
	if err != nil {
		fmt.Printf("%s", err)
		return nil, err
	} else {
		fmt.Println("Command Successfully Executed")
		output := string(out[:])
		fmt.Println(output)
	}

	content, err := Handler_gobuster(resultfile, domain)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	// 删除这个文件
	// 删除oneforall的多余日志和文件
	// out2, err := exec.Command("rm", "-rf", test_resultfile).Output()
	// docker里在开启
	out2, err := exec.Command("rm", "-rf", resultfile).Output()
	if err != nil {
		fmt.Printf("%s", err)
		return nil, err
	} else {
		fmt.Println("删除文件命令执行成功")
		output := string(out2[:])
		fmt.Println(output)
	}

	bs, err := json.Marshal(content)
	if err != nil {
		return nil, err
	}

	return bs, nil
}

func Handler_gobuster(resultfile, domain string) (*common.DirItems, error) {
	di := new(common.DirItems)
	fi, err := os.Open(resultfile)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return nil, err
	}
	defer fi.Close()
	br := bufio.NewReader(fi)
	for {
		a, _, c := br.ReadLine()
		if c == io.EOF {
			break
		}
		s := string(a)
		ss := strings.Split(s, "[--> ")
		if len(ss) > 1 {
			item := new(common.DirItem)
			ts := strings.Replace(ss[0], " ", "", -1)
			statess := strings.Split(ts, ")")[0]
			code := strings.Split(statess, ":")[1]
			dir := strings.Split(ss[1], "]")[0]

			Size_s := strings.Split(ss[0], "Size:")[1]
			Size := strings.Replace(strings.Split(Size_s, "]")[0], " ", "", -1)

			size, err := strconv.Atoi(Size)
			if err != nil {
				return nil, err
			}

			item.Code = code
			item.Dir = dir
			item.Size = size
			di.Dir = append(di.Dir, item)
			// di.Dir = append(di.Dir, sss[0])
		} else if len(ss) == 1 {
			it := new(common.DirItem)
			ts := strings.Replace(ss[0], " ", "", -1)

			Size_s := strings.Split(ts, "Size:")[1]
			Size := strings.Split(Size_s, "]")[0]

			size, err := strconv.Atoi(Size)
			if err != nil {
				return nil, err
			}

			// fmt.Println(Size_s, Size)

			tss := strings.Split(ts, "[")
			// tss[0] = /404(Status:200)
			tsss := strings.Split(tss[0], "(")
			dir := domain + tsss[0]
			it.Dir = dir
			it.Size = size

			codes := strings.Split(tsss[1], ":")[1]
			code := strings.Replace(codes, ")", "", -1)
			it.Code = code
			di.Dir = append(di.Dir, it)
		}
	}
	return di, nil
}
