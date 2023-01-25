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
			fmt.Println(err.Error())
			return err
		}

	}

	_, err = engine.Insert(&data)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	return nil
}
