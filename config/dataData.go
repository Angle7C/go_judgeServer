package config

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
)

type MysqlConfig struct {
	Url      string `yaml:"Url"`
	UserName string `yaml:"UserName"`
	Password string `yaml:"PassWord"`
	DataName string `yaml:"DataName"`
}

var (
	db *gorm.DB
)

func (mysqlConfig MysqlConfig) Init() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=true", mysqlConfig.UserName, mysqlConfig.Password, mysqlConfig.Url, mysqlConfig.DataName)
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatalf("链接数据库失败 %v", err.Error())
	} else {
	}
	log.Println("数据库链接成功")
}
func GetDb() *gorm.DB {
	return db
}
