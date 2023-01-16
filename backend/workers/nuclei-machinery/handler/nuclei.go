package handler

import (
	"bufio"
	"bytes"
	"common"
	"encoding/json"
	"fmt"
	uuid2 "github.com/google/uuid"
	"io"
	"log"
	"os"
	"os/exec"
)

const (
	nuclei_win_test_path = "E:\\tmp\\nuclei_2.8.6_windows_amd64\\nuclei.exe"
)

func NucleiScan(domains []string, tags string) ([]byte, error) {
	target_file, err := writeTargetsToFile(domains)
	if err != nil {
		return nil, err
	}
	result_file, err := execNucleiCVE(target_file, tags)
	if err != nil {
		return nil, err
	}
	vulns, err := readVulnFromFile(result_file)
	if err != nil {
		return nil, err
	}

	// deleteFile
	//deleteFile(target_file, result_file)
	return json.Marshal(vulns)
}

func readVulnFromFile(result_file string) ([]*common.NucleiVuln, error) {
	vulns := make([]*common.NucleiVuln, 0)
	file, err := os.OpenFile(result_file, os.O_RDWR, 0666)
	if err != nil {
		fmt.Println("Open file error!", err)
		return nil, err
	}
	defer file.Close()

	buf := bufio.NewReader(file)
	for {
		vuln := new(common.NucleiVuln)
		line, err := buf.ReadBytes('\n')
		if err != nil {
			if err == io.EOF {
				fmt.Println("File read ok!")
				break
			} else {
				fmt.Println("Read file error!", err)
				return nil, err
			}
		}
		err = json.Unmarshal(line, vuln)
		if err != nil {
			return nil, err
		}
		vulns = append(vulns, vuln)
	}
	return vulns, nil
}

func writeTargetsToFile(domains []string) (string, error) {
	uuid := uuid2.New()
	target_file := fmt.Sprintf("%s.txt", uuid.String())
	f, err := os.OpenFile(target_file, os.O_RDWR|os.O_CREATE, 0600)
	defer f.Close()
	if err != nil {
		fmt.Println(err.Error())
		return "", err
	}

	for _, v := range domains {
		v := v + "\n"
		_, err = f.Write([]byte(v))
		if err != nil {
			return "", err
		}
	}
	return target_file, nil
}

func execNucleiCVE(target_file, tags string) (string, error) {
	uuid := uuid2.New()
	result_file := fmt.Sprintf("%s.json", uuid.String())
	cmd := exec.Command(nuclei_win_test_path, "-l", target_file, "-tags", tags, "-stats", "-stats-interval", "60", "-json", "-o", result_file)
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		log.Fatal(err.Error(), stderr.String())
	}
	//outStr, errStr := string(out.Bytes()), string(stderr.Bytes())
	//fmt.Printf("out:\n%s\nerr:\n%s\n", outStr, errStr)

	return result_file, nil
}

func deleteFile(filepaths ...string) {
	for _, v := range filepaths {
		out2, err := exec.Command("rm", "-rf", v).Output()
		if err != nil {
			fmt.Printf("%s", err)
			continue
		} else {
			fmt.Println("删除文件命令执行成功")
			output := string(out2[:])
			fmt.Println(output)
		}
	}
}
