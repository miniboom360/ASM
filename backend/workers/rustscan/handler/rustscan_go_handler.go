package handler

import (
  "bufio"
  "common"
  "fmt"
  "io"
  "os"
  "regexp"
  "strconv"
  "strings"
)

// handle result write in txt from rustscan
// --ulimit 5000
func HandlerRustResult() ([]byte, error) {
  file_path := "/Users/liyang/tools/asm/ASM/backend/workers/rustscan/scripts/1.txt"
  fi, err := os.Open(file_path)
  if err != nil {
    fmt.Printf("Error: %s\n", err)
    return nil, err
  }
  defer fi.Close()
  br := bufio.NewReader(fi)

  count := 1
  service_count := 0
  service_flag := false
  vaild_count := 0
  output_count := 1
  rsits := new(common.RustScanItems)
  for {
    a, _, c := br.ReadLine()
    if c == io.EOF {
      break
    }
    s := string(a)
    s = delete_extra_space(s)
    st := strings.Split(s, " ")
    if len(st) == 4 && st[0] == "PORT" && st[1] == "STATE" && st[2] == "SERVICE" && st[3] == "REASON" {
      service_count = count
      service_flag = true
      fmt.Printf("find this, line count = %d\n", service_count)
    }
    if count > service_count && s != "" && service_flag {
      item := new(common.RustScanItem)

      ss := strings.Split(s, " ")
      status := ss[1]
      service := ss[2]
      reason := ss[3]
      sss := strings.Split(ss[0], "/")
      port, err := strconv.Atoi(sss[0])
      if err != nil {
        return nil, err
      }
      protocol := sss[1]

      item.Protocol = protocol
      item.Port = port
      item.Status = status
      item.Service = service
      item.Reason = reason
      rsits.Items = append(rsits.Items, item)
      output_count++
    }

    if count > service_count && s == "" && service_flag {
      vaild_count = count
      fmt.Printf("vaild count = %d\n", vaild_count)
      service_flag = false
      break
    }
    count++

  }

  fmt.Printf("output_count = %d\n", output_count)
  return nil, err
}

func delete_extra_space(s string) string {
  // 删除字符串中的多余空格，有多个空格时，仅保留一个空格
  s1 := strings.Replace(s, "	", " ", -1)      // 替换tab为空格
  regstr := "\\s{2,}"                          // 两个及两个以上空格的正则表达式
  reg, _ := regexp.Compile(regstr)             // 编译正则表达式
  s2 := make([]byte, len(s1))                  // 定义字符数组切片
  copy(s2, s1)                                 // 将字符串复制到切片
  spc_index := reg.FindStringIndex(string(s2)) // 在字符串中搜索
  for len(spc_index) > 0 {                     // 找到适配项
    s2 = append(s2[:spc_index[0]+1], s2[spc_index[1]:]...) // 删除多余空格
    spc_index = reg.FindStringIndex(string(s2))            // 继续在字符串中搜索
  }
  return string(s2)
}
