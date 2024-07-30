package models

import (
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

//func init() {
//
//	cfg, err := ini.Load("./conf/app.ini")
//	if err != nil {
//		fmt.Printf("Fail to read file: %v", err)
//		os.Exit(1)
//	}
//	ip := cfg.Section("mysql").Key("ip").String()
//	port := cfg.Section("mysql").Key("port").String()
//	user := cfg.Section("mysql").Key("user").String()
//	password := cfg.Section("mysql").Key("password").String()
//	database := cfg.Section("mysql").Key("database").String()
//
//	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local", user, password, ip, port, database)
//	//dsn := "root:123456@tcp(192.168.229.128:3306)/gin?charset=utf8mb4&parseTime=True&loc=Local"
//	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
//
//	if err != nil {
//		fmt.Println(err)
//	}
//	fmt.Printf("Db地址为%p\n", DB)
//}
