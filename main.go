package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
	"html/template"
	"shop/models"
	"shop/routers"
	"time"
)

type UserInfo struct {
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
}
type Article struct {
	Title   string `xml:"title" json:"title"`
	Content string `xml:"content"json:"content"`
}

func main() {

	r := gin.Default()
	//r.GET()
	//配置gin允许跨域请求
	//r.Use(cors.Default())
	f := cors.New(cors.Config{
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Cookie"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
		//AllowAllOrigins:  true,
		AllowOrigins: []string{"http://localhost:8081"},
	})
	r.Use(f)
	r.SetFuncMap(template.FuncMap{
		"UnixToTime": models.UnixToTime,
		"Str2Html":   models.Str2Html,
		"FormatImg":  models.FormatImg,
		"Sub":        models.Sub,
		"Substr":     models.Substr,
		"FormatAttr": models.FormatAttr,
		"Mul":        models.Mul,
	})
	r.LoadHTMLGlob("templates/**/**/*")
	r.Static("/static", "./static")

	//store := cookie.NewStore([]byte("secret"))
	//r.Use(sessions.Sessions("mysession", store))

	//使用中间件，因为每个context对应的上下文不同，每次来路由都需要执行这个session初始化，是针对不同的context
	store, _ := redis.NewStore(10, "tcp", "localhost:6379", "", []byte("secret"))
	r.Use(sessions.Sessions("mysession", store))

	routers.AdminRoutersInit(r)

	routers.ApiRoutersInit(r)

	routers.DefaultRoutersInit(r)

	r.Run()
}
