package module

import (
	"common"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"

	"io/ioutil"
	"xorm.io/xorm"
)

var engine *xorm.Engine

// var configpath = "D:\\code\\ASM\\backend\\api-gateway\\scripts\\config.json"

var configpath = "/Users/liyang/tools/asm/ASM/backend/api-gateway/scripts/config.json"

type config struct {
	Mysql MysqlConf
}
type MysqlConf struct {
	Address  string
	Port     int
	Password string
}

func ReadConf() (*config, error) {
	c := new(config)
	data, err := ioutil.ReadFile(configpath)
	if err != nil {
		return nil, err
	}
	// 读取的数据为json格式，需要进行解码
	err = json.Unmarshal(data, c)
	if err != nil {
		return nil, err
	}
	return c, nil
}

func InitMysql() error {
	var err error
	c, err := ReadConf()
	if err != nil {
		return err
	}

	dsn := fmt.Sprintf("root:%s@(%s:%d)/asm?charset=utf8mb4&parseTime=True&loc=Local", c.Mysql.Password,
		c.Mysql.Address, c.Mysql.Port)
	engine, err = xorm.NewEngine("mysql", dsn)
	if err != nil {
		return err
	}

	// ss, err := engine.DBMetas()
	// if err != nil {
	// 	fmt.Println(err.Error())
	// 	return err
	// }
	// for _, s := range ss {
	// 	fmt.Printf("s = %+v\n", s)
	// }

	return nil
}

func AddSubDomainItems(data []*common.Subdomains) error {
	// var affected int64

	res, err := engine.IsTableExist(common.Subdomains{})
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	if !res {
		err = engine.CreateTables(common.Subdomains{})
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
