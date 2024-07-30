package models

import (
	"fmt"
	"github.com/redis/go-redis/v9"
	"gopkg.in/ini.v1"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

func init() {
	config, iniErr := ini.Load("./conf/app.ini")
	if iniErr != nil {
		fmt.Printf("Fail to read file: %v", iniErr)
		os.Exit(1)
	}

	redisIp := config.Section("redis").Key("ip").String()
	redisPort := config.Section("redis").Key("port").String()
	redisEnable, _ = config.Section("redis").Key("redisEnable").Bool()

	fmt.Println("是否使用redis", redisEnable)
	if redisEnable {
		//连接redis数据库
		rdbClient = redis.NewClient(&redis.Options{
			Addr:     redisIp + ":" + redisPort,
			Password: "", // no password set
			DB:       0,  // use default DB
		})

		_, err := rdbClient.Ping(ctx).Result()
		if err != nil {
			fmt.Println("redis数据库连接失败")
		} else {
			fmt.Println("redis数据库连接成功...")
		}
	}

	//if redisEnable {
	//	store = RedisStore{}
	//} else {
	//	store = base64Captcha.DefaultMemStore
	//}

	ip := config.Section("mysql").Key("ip").String()
	port := config.Section("mysql").Key("port").String()
	user := config.Section("mysql").Key("user").String()
	password := config.Section("mysql").Key("password").String()
	database := config.Section("mysql").Key("database").String()

	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local", user, password, ip, port, database)
	//dsn := "root:123456@tcp(192.168.229.128:3306)/gin?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Db地址为%p\n", DB)
}
