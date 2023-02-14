package common

import (
	"errors"
	"xorm.io/xorm"
)

var engine *xorm.Engine

func SetTaskStatusByTaskId(taskid, status string) (bool, error) {
	if engine == nil {
		return false, errors.New("engine is null")
	}

	sql := "update `task_item` set staus=? where task_id=?"
	res, err := engine.Exec(sql, status, taskid)

	if err != nil {
		return false, err
	}

	// fmt.Printf("影响了%+v行\n", res)
	return true, nil
}

func SetMysqlConn(e *xorm.Engine) {
	engine = e
}
