package module

import (
	"backend/app"
	"errors"
	"strings"
)

func AddInitUser(data []*app.User) error {
	// var affected int64
	res, err := engine.IsTableExist(app.User{})
	if err != nil {
		panic(err)
		return err
	}
	if !res {
		err = engine.CreateTables(app.User{})
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

func GetUsersInfo() ([]*app.User, error) {

	users := make([]*app.User, 0)

	err := engine.Table("user").Find(&users)
	if err != nil {
		if strings.Contains(err.Error(), "Table 'ASM.user' doesn't exist") {
			return nil, errors.New(app.DB_TABLE_NOT_EXIST)
		} else {
			return nil, err
		}
	}
	// fmt.Println(result)
	return users, nil
}
