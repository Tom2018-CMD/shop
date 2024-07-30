package models

import (
	"fmt"
)

const CAPTCHA = "captcha:"

type RedisStore struct {
}

// 实现设置 captcha 的方法
func (r RedisStore) Set(id string, value string) error {
	key := CAPTCHA + id
	err := CacheDb.Set(key, value, 60)
	return err
}

// 实现获取 captcha 的方法
func (r RedisStore) Get(id string, clear bool) string {
	key := CAPTCHA + id
	var val string
	res := CacheDb.Get(key, &val)
	fmt.Println(res)
	if !res {
		return ""
	}
	if clear {
		r := CacheDb.Del(key)
		if r == false {
			fmt.Println(err)
			return ""
		}
	}
	return val
}

// 实现验证 captcha 的方法
func (r RedisStore) Verify(id, answer string, clear bool) bool {
	v := r.Get(id, clear)
	return v == answer
}
