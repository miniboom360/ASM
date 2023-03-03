package activitys

import (
	"backend/app"
	"bufio"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	uuid2 "github.com/google/uuid"
	"os"
	"os/exec"
)

const HttpxPath = "httpx"

// httpx -status-code -content-length -title -tech-detect -cdn -ip -follow-host-redirects -random-agent -t num
//
//	-json -o {output_file} -l {targets.txt}
func HttpxScan(ctx context.Context, hr app.HttpxReq) ([]*app.HttpXData, error) {
	hxd := make([]*app.HttpXData, 0)

	if len(hr.Targets) == 0 {
		return nil, errors.New("[Httpx Input Is Null]")
	}
	targets_file, err := WriteTargetsToFile(hr.Targets)
	if err != nil {
		return nil, err
	}

	output_file := uuid2.New().String() + ".json"
	//
	if hr.ThreadsNum == 0 {
		hr.ThreadsNum = 1
	}
	httpx_command := HttpxPath + fmt.Sprintf(" -status-code -content-length -title -tech-detect -cdn -ip -follow-host-redirects -random-agent -t %d "+
		" -json -o %s -l %s", hr.ThreadsNum, output_file, targets_file)

	cmd := exec.Command("bash", "-c", httpx_command)
	_, err = cmd.CombinedOutput()
	if err != nil {
		panic(err)
		return nil, err
	}

	readFile, err := os.Open(output_file)
	if err != nil {
		panic(err)
		return nil, err
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		var httpx app.HttpXData

		err = json.Unmarshal(fileScanner.Bytes(), &httpx)
		if err != nil {
			panic(err)
			return nil, err
		}

		hxd = append(hxd, &httpx)
	}

	readFile.Close()
	deleteFile(targets_file, output_file)
	return hxd, nil

}
