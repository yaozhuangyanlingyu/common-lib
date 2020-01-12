package mysql

import(
	"fmt"
	"log"
	"github.com/jinzhu/gorm"
	"github.com/astaxie/beego/config"
)

// conn数据库连接
var(
	conn *gorm.DB
)

/**
 * 连接mysql数据库
 */
func ConnectDB() {
	if conn != nil {
		return
	}
	c, err := gorm.Open("mysql", buildConnString())
	if err != nil {
		log.Fatalf("MySQL Connect Error %v", err)
	}

	c.DB().SetMaxIdleConns(10)
	c.DB().SetMaxOpenConns(100)
	conn = c
}

/**
 * 获取MySQL数据库
 */
func GetDB() *gorm.DB {
	if conn == nil {
		log.Print("mysql gone away")
		return nil
	}

	return conn
}

/**
 * 生产连接字符串
 */
func buildConnString() string {
	iniConf, err := config.NewConfig("ini", "config/dev.ini")
	if err != nil {
		log.Print("parse dev.ini error：%v", err)
		return ""
	}

	return fmt.Sprintf(
		"%s:%s@(%s:%s)/%s?charset=%s&parseTime=%s&loc=%s",
		iniConf.String("aplumDB::user"),
		iniConf.String("aplumDB::password"),
		iniConf.String("aplumDB::host"),
		iniConf.String("aplumDB::port"),
		iniConf.String("aplumDB::db"),
		iniConf.String("aplumDB::charset"),
		iniConf.String("aplumDB::parsetime"),
		iniConf.String("aplumDB::loc"),
	)
}

