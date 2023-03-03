package activitys

import (
	"bufio"
	"fmt"
	uuid2 "github.com/google/uuid"
	"os"
)

func WriteTargetsToFile(targets []string) (file_name string, err error) {
	targets_file := fmt.Sprintf("%s.txt", uuid2.New())
	file, err := os.OpenFile(targets_file, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return "", err
	}
	// 及时关闭file句柄
	defer file.Close()
	write := bufio.NewWriter(file)
	for _, t := range targets {
		write.WriteString(t + "\n")
	}
	// Flush将缓存的文件真正写入到文件中
	write.Flush()

	return targets_file, nil
}
