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

const (
	NaabuPath = "naabu"
	NmapPath  = "nmap"
)

func NaabuScan(ctx context.Context, nr app.PortScanReq) (map[string]*app.NaabuData, error) {
	// nds := make([]*app.NaabuData, 0)
	if nr.Targets == nil || len(nr.Targets) == 0 {
		return nil, errors.New("please input right params")
	}
	nds := make(map[string]*app.NaabuData, 0)
	targets_file, err := WriteTargetsToFile(nr.Targets)
	if err != nil {
		return nil, err
	}

	output_file := uuid2.New().String() + ".json"
	naabu_command := NaabuPath + fmt.Sprintf(" -list %s -json -o %s -exclude-cdn ", targets_file, output_file)

	if nr.Tag == "top-100" {
		naabu_command += " -top-ports 100 "
	}
	if nr.Tag == "full" {
		naabu_command += " -p -"
	}
	if nr.Tag == "top-1000" {
		naabu_command += " -top-ports 1000 "
	}

	cmd := exec.Command("bash", "-c", naabu_command)
	_, err = cmd.CombinedOutput()
	if err != nil {
		str_err := err.Error()
		fmt.Println(str_err)
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
		var naabu app.NaabuData
		err = json.Unmarshal(fileScanner.Bytes(), &naabu)
		if err != nil {
			panic(err)
			return nil, err
		}

		if naabu.Host == "" {
			nds[naabu.IP] = &naabu
		} else {
			nds[naabu.Host] = &naabu
		}

		// nds = append(nds, &naabu)
	}

	readFile.Close()

	deleteFile(targets_file, output_file)
	return nds, nil
}
