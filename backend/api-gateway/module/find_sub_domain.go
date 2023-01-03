package module

import (
	"common"
	"encoding/json"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"io/ioutil"
)

var mysql_db *gorm.DB
var configpath = "D:\\code\\asm-demo\\backend\\api-gateway\\scripts\\config.json"
var test_configpath = "/Users/liyang/tools/asm/ASM/backend/api-gateway/scripts/config.json"

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
	data, err := ioutil.ReadFile(test_configpath)
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
	c, err := ReadConf()
	if err != nil {
		return err
	}
	dsn := fmt.Sprintf("root:%s@tcp(%s:%d)/asm?charset=utf8mb4&parseTime=True&loc=Local", c.Mysql.Password,
		c.Mysql.Address, c.Mysql.Port)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	mysql_db = db
	return nil
}

func AddSubDomainItems(data []*common.Subdomains) error {
	// 增加发现时间
	result := mysql_db.Create(&data)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
