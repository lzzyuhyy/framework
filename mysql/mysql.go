package mysql

import (
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// 提供链接mysql功能的组件---并在使用完之后断开链接
func CreateMySQLClient(handler func(db *gorm.DB) error) error { // 闭包接收用户的sql操作
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local",
		viper.GetString("mysql.user"),
		viper.GetString("mysql.pass"),
		viper.GetString("mysql.host"),
		viper.GetString("mysql.port"),
		viper.GetString("mysql.dbname"),
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
