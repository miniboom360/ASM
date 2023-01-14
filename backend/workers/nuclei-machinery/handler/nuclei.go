package handler

import (
	"fmt"
	uuid2 "github.com/google/uuid"
	"io/ioutil"
	"os"
	"os/exec"
)

const (
	nuclei_win_test_path = "E:\\tmp\\nuclei_2.8.6_windows_amd64\\nuclei.exe"
)

func NucleiScan(domains []string, tags string) ([]byte, error) {
	// nuclei -duc -tags cve -severity low,medium,high,critical -type http -l targets.txt -json -stats -stats-interval 60 -o target.json
	// write domains to targets.txt
	target_file, err := writeTargetsToFile(domains)
	if err != nil {
		return nil, err
	}

	f, err := ioutil.ReadFile(target_file)
	if err != nil {
		fmt.Println("read fail", err)
	}
	fmt.Println(string(f))

	result_file, err := execNucleiCVE(target_file, tags)
	if err != nil {
		return nil, err
	}
	fmt.Println(result_file)
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
	result_file := fmt.Sprintf("D:\\code\\ASM\\backend\\workers\\nuclei-machinery\\tests\\%s.json", uuid.String())
	command := fmt.Sprintf(" -duc -tags %s -severity low,medium,high,critical -type http -l %s -json -stats -stats-interval 60 -o %s",
		tags, target_file, result_file)
	out1, err := exec.Command(nuclei_win_test_path, command).Output()
	if err != nil {
		fmt.Printf("%s", err)
		return "", err
	} else {
		fmt.Println("exec nuclei success!")
		output := string(out1[:])
		fmt.Println(output)
	}
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
