package handler

import (
  "fmt"
  uuid2 "github.com/google/uuid"
  "os"
  "os/exec"
)

func CVE_NucleiScan(domains []string) ([]byte, error) {
  // nuclei -duc -tags cve -severity low,medium,high,critical -type http -l targets.txt -json -stats -stats-interval 60 -o target.json
  // write domains to targets.txt
  uuid := uuid2.New()
  target_file := fmt.Sprintf("%s.txt", uuid.String())
  f, err := os.OpenFile(target_file, os.O_RDWR|os.O_CREATE, 0600)
  defer f.Close()
  if err != nil {
    fmt.Println(err.Error())
    return nil, err
  }

  for _, v := range domains {
    _, err = f.Write([]byte(v))
    if err != nil {
      return nil, err
    }
  }

  // deleteFile
  deleteFile(target_file)

  return nil, nil
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
