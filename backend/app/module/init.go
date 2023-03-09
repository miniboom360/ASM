package module

import (
	"backend/app"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
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

	// 增加一个用户表
	AddInitOneAdmin()
	return nil
}

func AddInitOneAdmin() error {
	users, err := GetUsersInfo()
	if err != nil {
		if err.Error() != app.DB_TABLE_NOT_EXIST {
			return err
		}
	}

	if len(users) != 0 {
		return nil
	}
	// 如果没有的话，就创建一个
	data := make([]*app.User, 0)
	user := new(app.User)
	user.UserId = "1"
	user.Username = "miniboom"
	user.RealName = "LiYang"
	user.Desc = "Hacker"
	user.Token = uuid.New().String()
	user.Password = "miniboom"
	user.HomePath = "/dashboard/analysis"
	user.Roles = make([]*app.Role, 0)
	role := new(app.Role)
	role.RoleName = "Super Admin"
	role.Value = "super"

	user.Roles = append(user.Roles, role)

	data = append(data, user)

	if err := AddInitUser(data); err != nil {
		panic(err)
		return err
	}
	return nil
}
