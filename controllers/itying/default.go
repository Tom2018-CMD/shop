package itying

import (
	"fmt"
	"shop/models"
	"time"

	"github.com/gin-gonic/gin"
)

type DefaultController struct {
	BaseController
}

func (con DefaultController) Index(c *gin.Context) {

	//models.Cookie.Set(c, "username", "李四")
	//str := ""
	//var str string
	//
	//models.Cookie.Get(c, "username", &str)
	//
	//fmt.Println(str)

	timeStart := time.Now().UnixNano()
	////1、获取顶部导航
	//topNavList := []models.Nav{}
	//if hasTopNavList := models.CacheDb.Get("topNavList", &topNavList); !hasTopNavList {
	//	models.DB.Where("status=1 AND position=1").Find(&topNavList)
	//	err := models.CacheDb.Set("topNavList", topNavList, 60*60)
	//	if err != nil {
	//		fmt.Println(err)
	//		return
	//	}
	//}

	//2、获取轮播图数据
	focusList := []models.Focus{}
	if hasFocusList := models.CacheDb.Get("focusList", &focusList); !hasFocusList {
		models.DB.Where("status=1 AND focus_type=1").Find(&focusList)
		models.CacheDb.Set("focusList", focusList, 60*60)
	}

	////3、获取分类的数据
	//goodsCateList := []models.GoodsCate{}
	//
	//if hasGoodsCateList := models.CacheDb.Get("goodsCateList", &goodsCateList); !hasGoodsCateList {
	//	//https://gorm.io/zh_CN/docs/preload.html
	//	models.DB.Where("pid = 0 AND status=1").Order("sort DESC").Preload("GoodsCateItems", func(db *gorm.DB) *gorm.DB {
	//		return db.Where("goods_cate.status=1").Order("goods_cate.sort DESC")
	//	}).Find(&goodsCateList)
	//
	//	models.CacheDb.Set("goodsCateList", goodsCateList, 60*60)
	//}

	////4、获取中间导航
	//middleNavList := []models.Nav{}
	//if hasMiddleNavList := models.CacheDb.Get("middleNavList", &middleNavList); !hasMiddleNavList {
	//	models.DB.Where("status=1 AND position=2").Find(&middleNavList)
	//	for i := 0; i < len(middleNavList); i++ {
	//		relation := strings.ReplaceAll(middleNavList[i].Relation, "，", ",") //21，22,23,24
	//		relationIds := strings.Split(relation, ",")
	//		goodsList := []models.Goods{}
	//		models.DB.Where("id in ?", relationIds).Select("id,title,goods_img,price").Find(&goodsList)
	//		middleNavList[i].GoodsItems = goodsList
	//	}
	//	models.CacheDb.Set("middleNavList", middleNavList, 60*60)
	//}

	//手机
	phoneList := []models.Goods{}
	if hasPhoneList := models.CacheDb.Get("phoneList", &phoneList); !hasPhoneList {
		phoneList = models.GetGoodsByCategory(1, "best", 8)
		models.CacheDb.Set("phoneList", phoneList, 60*60)
	}

	//配件

	otherList := []models.Goods{}
	if hasOtherList := models.CacheDb.Get("otherList", &otherList); !hasOtherList {
		otherList = models.GetGoodsByCategory(9, "all", 1)
		models.CacheDb.Set("otherList", otherList, 60*60)
	}

	timeEnd := time.Now().UnixNano()

	fmt.Printf("执行时间：%v ms", (timeEnd-timeStart)/1000000)

	tpl := "itying/index/index.html"

	data := gin.H{
		"focusList": focusList,
		"phoneList": phoneList,
		"otherList": otherList,
	}

	con.Render(c, tpl, data)

	//c.HTML(http.StatusOK, "itying/index/index.html", gin.H{
	//	"topNavList":    topNavList,
	//	"focusList":     focusList,
	//	"goodsCateList": goodsCateList,
	//	"middleNavList": middleNavList,
	//	"phoneList":     phoneList,
	//	"otherList":     otherList,
	//})

	/*

		5000万

		50ms

		20   1s


		100ms

		20   2s


		缓存技术：文件缓存  内存缓存  数据库缓存   Redis


	*/

	// models.CacheDb.Set("username", "张三", 600)

	// var username string
	// if hasUsername := models.CacheDb.Get("username", &username); hasUsername {
	// 	fmt.Println(username)
	// }

}
