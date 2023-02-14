package module

import (
	"common"
	"fmt"
)

func AddTaskItem(data []*common.TaskItem) error {

	res, err := engine.IsTableExist(common.TaskItem{})
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	if !res {
		err = engine.CreateTables(common.TaskItem{})
		if err != nil {
			panic(err)
			return err
		}

	}

	_, err = engine.Insert(&data)
	if err != nil {
		panic(err)
		return err
	}
	return nil
}

func CheckWrongTask(org_name string, domains []string) (bool, error) {
	data := make([]*common.TaskItem, 0)
	// 要增加缓存了，不然这样太慢了
	engine.Table("task_item").Select("*").Find(&data)
	// engine.Where("domains = ?", domains).Get(&data)
	for _, d := range data {
		if IsIdentical(domains, d.Domains) && org_name != d.OrganizationName {
			return false, nil
		}
	}
	// 获取所有任务domains列表
	return true, nil
}

// true有重复的元素，false无相同元素
func IsIdentical(str1 []string, str2 []string) (t bool) {
	t = false
	if len(str1) == 0 || len(str2) == 0 {
		return
	}
	map1, map2 := make(map[string]int), make(map[string]int)
	for i := 0; i < len(str1); i++ {
		map1[str1[i]] = i
	}
	for i := 0; i < len(str2); i++ {
		map2[str2[i]] = i
	}
	for k, _ := range map1 {
		if _, ok := map2[k]; ok {
			t = true
		}
	}
	return
}
