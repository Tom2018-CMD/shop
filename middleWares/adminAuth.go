package middleWares

import (
	"encoding/json"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"gopkg.in/ini.v1"
	"net/http"
	"os"
	"shop/models"
	"strings"
)

func InitAdminAuthMiddleWare(c *gin.Context) {
	fmt.Println("进入中间件")
	//excludeAuthPath("aaa")

	pathname := strings.Split(c.Request.URL.String(), "?")[0]
	//fmt.Println(pathname)

	session := sessions.Default(c)
	userinfo := session.Get("userinfo")

	//fmt.Println(userinfo)
	userinfoStr, ok := userinfo.(string)
	if ok {
		var userinfoStruct []models.Manager
		err := json.Unmarshal([]byte(userinfoStr), &userinfoStruct)
		if err != nil || !(len(userinfoStruct) > 0 && userinfoStruct[0].Username != "") {
			if pathname != "/admin/login" && pathname != "/admin/doLogin" && pathname != "/admin/captcha" {
				c.Redirect(http.StatusFound, "/admin/login")
			}
		} else {
			//
			urlPath := strings.Replace(pathname, "/admin/", "", 1)
			if userinfoStruct[0].IsSuper == 0 && !excludeAuthPath("/"+urlPath) {

				roleAccess := []models.RoleAccess{}
				models.DB.Where("role_id = ?", userinfoStruct[0].RoleId).Find(&roleAccess)
				//放入一个map方便比较
				roleAccessMap := make(map[int]int)
				for _, v := range roleAccess {
					roleAccessMap[v.AccessId] = v.AccessId
				}

				//获取url权限的id
				access := models.Access{}
				models.DB.Where("url = ?", urlPath).Find(&access)

				if _, ok := roleAccessMap[access.Id]; !ok {
					c.String(http.StatusOK, "没有权限")
					c.Abort()
				}
			}
		}
	} else {
		if pathname != "/admin/login" && pathname != "/admin/doLogin" && pathname != "/admin/captcha" {
			c.Redirect(http.StatusFound, "/admin/login")
		}
	}

}

// 排除权限判断的方法
func excludeAuthPath(urlPath string) bool {
	cfg, err := ini.Load("./conf/app.ini")
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		os.Exit(1)
	}
	excludeAuthPath := cfg.Section("").Key("excludeAuthPath").String()

	excludeAuthPathSlice := strings.Split(excludeAuthPath, ",")

	for _, v := range excludeAuthPathSlice {
		if v == urlPath {
			return true
		}
	}
	return false
}
