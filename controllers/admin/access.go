package admin

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"shop/models"
	pbAccess "shop/proto/rbacAccess"
	"strings"
)

type AccessController struct {
	BaseController
}

func (con AccessController) Index(c *gin.Context) {

	accessSrv := pbAccess.NewRbacAccessService("rbac", models.MicroRbacClient)
	res, err := accessSrv.AccessGet(context.Background(), &pbAccess.AccessGetRequest{})
	if err != nil {
		return
	}
	//accessList := []models.Access{}
	//models.DB.Where("module_id = ?", 0).Preload("AccessItem").Find(&accessList)
	//fmt.Printf("%#v", accessList)
	c.HTML(http.StatusOK, "admin/access/index.html", gin.H{
		"accessList": res.AccessList,
	})
}

func (con AccessController) Add(c *gin.Context) {

	accessSrv := pbAccess.NewRbacAccessService("rbac", models.MicroRbacClient)
	res, err := accessSrv.AccessGet(context.Background(), &pbAccess.AccessGetRequest{})
	if err != nil {
		return
	}
	c.HTML(http.StatusOK, "admin/access/add.html", gin.H{
		"accessList": res.AccessList,
	})
}

func (con AccessController) DoAdd(c *gin.Context) {

	moduleName := strings.Trim(c.PostForm("module_name"), " ")
	accessType, err1 := models.Int(c.PostForm("type"))
	actionName := c.PostForm("action_name")
	url := c.PostForm("url")
	moduleId, err2 := models.Int(c.PostForm("module_id"))
	sort, err3 := models.Int(c.PostForm("sort"))
	status, err4 := models.Int(c.PostForm("status"))
	description := c.PostForm("description")
	if err1 != nil || err2 != nil || err3 != nil || err4 != nil {
		con.Error(c, "传入参数错误", "/admin/access/add")
		return
	}
	if moduleName == "" {
		con.Error(c, "模块名称不能为空", "/admin/access/add")
		return
	}
	accessSrv := pbAccess.NewRbacAccessService("rbac", models.MicroRbacClient)
	res, err := accessSrv.AccessAdd(context.Background(), &pbAccess.AccessAddRequest{
		ModuleName:  moduleName,
		ActionName:  actionName,
		Type:        int64(accessType),
		Url:         url,
		ModuleId:    int64(moduleId),
		Sort:        int64(sort),
		Description: description,
		Status:      int64(status),
		AddTime:     models.GetUnix(),
	})

	if !res.Success {
		fmt.Println(err)
		con.Error(c, "增加数据失败", "/admin/access/add")
		return
	}
	con.Success(c, "增加数据成功", "/admin/access/add")
}

func (con AccessController) Edit(c *gin.Context) {

	id, err1 := models.Int(c.Query("id"))
	if err1 != nil {
		con.Error(c, "参数错误", "/admin/access")
	}
	accessSrv := pbAccess.NewRbacAccessService("rbac", models.MicroRbacClient)
	access, _ := accessSrv.AccessGet(context.Background(), &pbAccess.AccessGetRequest{
		Id: int64(id),
	})

	accessList, _ := accessSrv.AccessGet(context.Background(), &pbAccess.AccessGetRequest{})
	c.HTML(http.StatusOK, "admin/access/edit.html", gin.H{
		"access":     access.AccessList[0],
		"accessList": accessList.AccessList,
	})
}

func (con AccessController) DoEdit(c *gin.Context) {

	id, err1 := models.Int(c.PostForm("id"))
	moduleName := strings.Trim(c.PostForm("module_name"), " ")
	accessType, err2 := models.Int(c.PostForm("type"))
	actionName := c.PostForm("action_name")
	url := c.PostForm("url")
	moduleId, err3 := models.Int(c.PostForm("module_id"))
	sort, err4 := models.Int(c.PostForm("sort"))
	status, err5 := models.Int(c.PostForm("status"))
	description := c.PostForm("description")
	if err1 != nil || err2 != nil || err3 != nil || err4 != nil || err5 != nil {
		con.Error(c, "传入参数错误", "/admin/access")
		return
	}
	if moduleName == "" {
		con.Error(c, "模块名称不能为空", "/admin/access/edit?id="+models.String(id))
		return
	}

	accessSrv := pbAccess.NewRbacAccessService("rbac", models.MicroRbacClient)
	res, _ := accessSrv.AccessEdit(context.Background(), &pbAccess.AccessEditRequest{
		Id:          int64(id),
		ModuleName:  moduleName,
		ActionName:  actionName,
		Type:        int64(accessType),
		Url:         url,
		ModuleId:    int64(moduleId),
		Sort:        int64(sort),
		Description: description,
		Status:      int64(status),
		AddTime:     models.GetUnix(),
	})

	if !res.Success {
		con.Error(c, "修改数据失败", "/admin/access/edit?id="+models.String(id))
	} else {
		con.Success(c, "修改数据成功", "/admin/access/edit?id="+models.String(id))
	}

}

func (con AccessController) Delete(c *gin.Context) {

	id, err1 := models.Int(c.Query("id"))
	if err1 != nil {
		con.Error(c, "传入数据错误", "/admin/access")
	} else {
		accessSrv := pbAccess.NewRbacAccessService("rbac", models.MicroRbacClient)
		res, _ := accessSrv.AccessDelete(context.Background(), &pbAccess.AccessDeleteRequest{
			Id: int64(id),
		})
		if res.Success {
			con.Success(c, res.Message, "/admin/access")
		} else {
			con.Error(c, res.Message, "/admin/access")
		}
	}
}
