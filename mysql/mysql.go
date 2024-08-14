package mysql

import (
	"fmt"
	"framework/nacos"
	"gopkg.in/yaml.v3"
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

func CreateMySQLClient(handler func(db *gorm.DB) error) error {
	var conf Config
	config, err := nacos.GetNacosConfig()
	if err != nil {
		return err
	}

	err = yaml.Unmarshal([]byte(config), &conf)
	if err != nil {
		return nil
	}

	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local",
		conf.Mysql.User,
		conf.Mysql.Pass,
		conf.Mysql.Host,
		conf.Mysql.Port,
		conf.Mysql.Dbname,
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	return handler(db)
}

func CreatedTxClient(handle func(db *gorm.DB) error) error {
	return CreateMySQLClient(func(db *gorm.DB) error {
		tx := db.Begin()

		err := handle(tx)
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}

		return err
	})
}
