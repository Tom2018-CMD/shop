package models

//
//import (
//	"context"
//	"github.com/redis/go-redis/v9"
//	"sync"
//)
//
//var (
//	RedisCoreCtx = context.Background()
//	RedisDb      *redis.Client
//	Wg           sync.WaitGroup
//)
//
//func init() {
//	RedisDb = redis.NewClient(&redis.Options{
//		Addr:     "127.0.0.1:6379",
//		Password: "", // no password set
//		DB:       0,  // use default DB
//	})
//	Wg.Add(3)
//	//go func() {
//	//	defer Wg.Done()
//	//	RedisDb.Set(RedisCoreCtx, "sdad", "sada", -2)
//	//}()
//	//go func() {
//	//	defer Wg.Done()
//	//	RedisDb.Set(RedisCoreCtx, "sdad1", "sada", -2)
//	//}()
//	//go func() {
//	//	defer Wg.Done()
//	//	RedisDb.Set(RedisCoreCtx, "sdad2", "sada", -2)
//	//	fmt.Println(123)
//	//}()
//	//Wg.Wait()
//	//_, err := RedisDb.Ping(RedisCoreCtx).Result()
//	//if err != nil {
//	//	println(err)
//	//}
//}
//
////type stu struct {
////	*Goods
////	GoodsAttr
////}
////
////func test() {
////	stu := stu{
////		&Goods{},
////		GoodsAttr{},
////	}
////}
