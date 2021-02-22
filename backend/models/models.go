package models

import (
	"exam/pkg/setting"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Model struct {
	CreateTime int       `json:"create_time"`
	UpdateTime time.Time `json:"update_time"`
}

func init() {
	var err error
	db, err = gorm.Open(setting.DatabaseSetting.Type,
		fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
			setting.DatabaseSetting.User,
			setting.DatabaseSetting.Password,
			setting.DatabaseSetting.Host,
			setting.DatabaseSetting.Name))
	if err != nil {
		log.Println(err)
	}

	// 前缀
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return setting.DatabaseSetting.TablePrefix + defaultTableName
	}

	// 默认把你的struct中的大写字母转换为小写并加上“s”,
	// 所以可以加上db.SingularTable(true) 让grom转义struct名字的时候不用加上s
	db.SingularTable(true)

	// 启用Logger，显示详细日志
	db.LogMode(true)

	// SetMaxIdleConns用于设置闲置的连接数。
	db.DB().SetMaxIdleConns(10)

	// SetMaxOpenConns用于设置最大打开的连接数，默认值为0表示不限制。
	db.DB().SetMaxOpenConns(100)
}

func CloseDb() {
	db.Close()
}
