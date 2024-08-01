package mysql

import (
	"fmt"
	"github.com/lzzyuhyy/framework/nacos"
	"gopkg.in/yaml.v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Config struct {
	Mysql `yaml:"mysql"`
}

type Mysql struct {
	Host   string `yaml:"host"`
	Port   int64  `yaml:"port"`
	User   string `yaml:"user"`
	Pass   string `yaml:"pass"`
	Dbname string `yaml:"dbname"`
}

var Conf Config

// 提供链接mysql功能的组件---并在使用完之后断开链接
func CreateMySQLClient(handler func(db *gorm.DB) error) error { // 闭包接收用户的sql操作

	config, err := nacos.GetConfig()
	if err != nil {
		return err
	}
	err = yaml.Unmarshal([]byte(config), &Conf)

	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local",
		Conf.Mysql.User,
		Conf.Mysql.Pass,
		Conf.Mysql.Host,
		Conf.Mysql.Port,
		Conf.Mysql.Dbname,
	)
	// 建立链接
	cli, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		return err
	}
	// 关闭链接
	db, err := cli.DB()
	if err != nil {
		return err
	}
	defer func() {
		db.Close()
	}()

	// 执行操作
	return handler(cli)
}

// 具备事务功能的mysql链接
func CreateTXClient(handler func(db *gorm.DB) error) error {
	return CreateMySQLClient(func(db *gorm.DB) error {
		tx := db.Begin()

		err := handler(tx)
		defer func() {
			if err == nil {
				tx.Commit()
			} else {
				tx.Rollback()
			}
		}()
		return err
	})
}
