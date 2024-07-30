package admin

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
	"shop/models"
	pbRbac "shop/proto/rbacLogin"
)

type LoginController struct {
	BaseController
}

func (con LoginController) Index(c *gin.Context) {
	//fmt.Println(models.Md5("123456"))
	c.HTML(http.StatusOK, "admin/login/login.html", gin.H{})
}

func (con LoginController) DoLogin(c *gin.Context) {

	username := c.PostForm("username")
	password := c.PostForm("password")
	captchaId := c.PostForm("captchaId")
	verifyValue := c.PostForm("verifyValue")
	if verifyValue != "" && models.VerifyCaptcha(captchaId, verifyValue) {

		password = models.Md5(password)
		//userinfoList := []models.Manager{}
		//models.DB.Where("username = ? and password = ?", username, password).Find(&userinfoList)
		rbacSrv := pbRbac.NewRbacLoginService("rbac", models.MicroRbacClient)
		rsp, err := rbacSrv.Login(context.Background(), &pbRbac.LoginRequest{
			Username: username,
			Password: password,
		})
		if err != nil {
			return
		}
		if rsp.IsLogin {
			//设置session
			session := sessions.Default(c)
			userinfoSlice, _ := json.Marshal(rsp.Userlist)
			session.Set("userinfo", string(userinfoSlice))
			session.Save()
			con.Success(c, "登陆成功", "/admin/")
		} else {
			con.Error(c, "用户名或者密码错误", "/admin/login")
		}

	} else {
		con.Error(c, "验证码为空或验证码错误", "/admin/login")
	}
}

func (con LoginController) Captcha(c *gin.Context) {
	id, b64s, answer, err := models.MakeCaptcha(40, 100, 2)

	fmt.Println("answer:", answer)
	if err != nil {
		fmt.Println(err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"captchaId":    id,
		"captchaImage": b64s,
	})
}

func (con LoginController) LoginOut(c *gin.Context) {
	session := sessions.Default(c)
	session.Delete("userinfo")
	session.Save()
	con.Success(c, "退出登陆成功 ", "/admin/login")
}
