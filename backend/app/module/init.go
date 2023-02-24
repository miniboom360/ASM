package module

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"xorm.io/xorm"
)

var engine *xorm.Engine

func InitMysql() error {
	var err error

	dsn := fmt.Sprintf("root:%s@(%s:%d)/ASM?charset=utf8mb4&parseTime=True&loc=Local", "miniboom",
		"106.75.13.27", 3306)
	engine, err = xorm.NewEngine("mysql", dsn)
	if err != nil {
		return err
	}
	return nil
}
