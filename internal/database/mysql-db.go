package database

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
	"ych/vgo/internal/global"
)

// ConnectMysql 连接mysql数据库
func ConnectMysql() *gorm.DB {
	dbConf := global.Config.DbConf
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True", dbConf.Username, dbConf.Password, dbConf.Hostname, dbConf.HostPort, dbConf.Database)
	var err error
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true, // 禁用外键约束
	})
	if err != nil {
		fmt.Printf("数据库连接错误: %v\n", err)
		return nil
	}
	sqlDB, err := db.DB()
	if err != nil {
		fmt.Printf("获取数据库实例错误: %v\n", err)
		return nil
	}
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(500)
	sqlDB.SetConnMaxLifetime(time.Hour)
	return db
}
