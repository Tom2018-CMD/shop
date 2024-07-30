package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type ApiController struct {
}

func (con ApiController) Index(c *gin.Context) {
	c.String(http.StatusOK, "我是一个api")
}

func (con ApiController) UserList(c *gin.Context) {
	c.String(http.StatusOK, "我是一个api--userList")
}

func (con ApiController) PList(c *gin.Context) {
	c.String(http.StatusOK, "我是一个api--Plist")
}
