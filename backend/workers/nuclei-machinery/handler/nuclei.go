package handler

import (
	"bytes"
	"fmt"
	uuid2 "github.com/google/uuid"
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
	execNucleiCVE(target_file, "thinkphp")
	// deleteFile
	//deleteFile(target_file)

	return nil, nil
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
	//nuclei -duc -tags cve -severity low,medium,high,critical -type http -l targets.txt -json -stats -stats-interval 60 -o target.json
	uuid := uuid2.New()
	//nuclei.exe -duc -tags thinkphp -severity low,medium,high,critical -type http -l .\946aba5e-2e2d-4da8-b5f5-2fbd67bc2ca6.txt -json -stats -stats-interval 60 -o .\target1.json
	result_file := fmt.Sprintf("%s.json", uuid.String())
	//command := fmt.Sprintf(" -duc -tags %s -severity low,medium,high,critical -type http -u %s -json -stats -stats-interval 60 -o %s",
	//	tags, target_domain, result_file)
	//command := fmt.Sprintf(" -duc -tags %s -u %s -json -stats -stats-interval 60 -o %s",
	//	tags, target_domain, result_file)
	//command := fmt.Sprintf(" -duc -tags %s -u %s",
	//	tags, target_domain)
	//out1, err := exec.Command(nuclei_win_test_path, command).Output()

	//cmd := exec.Command(nuclei_win_test_path, "-l", target_file, "-tags", tags, "-stats", "-stats-interval", "60", "-json", "-o", result_file)
	cmd := exec.Command(nuclei_win_test_path, "-l", target_file, "-stats", "-stats-interval", "60", "-json", "-o", result_file)
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		log.Fatal(err.Error(), stderr.String())
	}
	outStr, errStr := string(out.Bytes()), string(stderr.Bytes())
	fmt.Printf("out:\n%s\nerr:\n%s\n", outStr, errStr)

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
