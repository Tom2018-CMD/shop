package admin

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"shop/models"
	pbRole "shop/proto/rbacRole"
	"strings"
)

type RoleController struct {
	BaseController
}

func (con RoleController) Index(c *gin.Context) {
	rbacSrv := pbRole.NewRbacRoleService("rbac", models.MicroRbacClient)
	rsp, err := rbacSrv.RoleGet(context.Background(), &pbRole.RoleGetRequest{})
	if err != nil {
		return
	}
	//roleList := []models.Role{}
	//models.DB.Find(&roleList)
	c.HTML(http.StatusOK, "admin/role/index.html", gin.H{
		"roleList": rsp.RoleList,
	})
}

func (con RoleController) Add(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/role/add.html", gin.H{})
}

func (con RoleController) DoAdd(c *gin.Context) {

	title := strings.Trim(c.PostForm("title"), " ")
	description := strings.Trim(c.PostForm("description"), " ")
	if title == "" {
		con.Error(c, "角色标题不能为空", "/admin/role/add")
		return
	}
	rbacSrv := pbRole.NewRbacRoleService("rbac", models.MicroRbacClient)
	rsp, _ := rbacSrv.RoleAdd(context.Background(), &pbRole.RoleAddRequest{
		Title:       title,
		Description: description,
		Status:      1,
		AddTime:     models.GetUnix(),
	})
	if !rsp.Success {
		con.Error(c, "增加角色失败", "/admin/role/add")
	} else {
		con.Success(c, "增加角色成功", "/admin/role")
	}
}

func (con RoleController) Edit(c *gin.Context) {
	id, err := models.Int(c.Query("id"))
	if err != nil {
		con.Error(c, "传入数据错误", "/admin/role")
	} else {

		rbacSrv := pbRole.NewRbacRoleService("rbac", models.MicroRbacClient)
		rsp, err := rbacSrv.RoleGet(context.Background(), &pbRole.RoleGetRequest{
			Id: int64(id),
		})
		if err != nil {
			return
		}
		c.HTML(http.StatusOK, "admin/role/edit.html", gin.H{
			"role": rsp.RoleList[0],
		})
	}

}
func (con RoleController) DoEdit(c *gin.Context) {
	id, err1 := models.Int(c.PostForm("id"))
	if err1 != nil {
		con.Error(c, "传入数据错误", "/admin/role")
		return
	}
	title := strings.Trim(c.PostForm("title"), " ")
	description := strings.Trim(c.PostForm("description"), " ")
	if title == "" {
		con.Error(c, "角色标题不能为空", "/admin/role/edit")
	}

	rbacSrv := pbRole.NewRbacRoleService("rbac", models.MicroRbacClient)
	rsp, _ := rbacSrv.RoleEdit(context.Background(), &pbRole.RoleEditRequest{
		Id:          int64(id),
		Title:       title,
		Description: description,
		Status:      1,
		AddTime:     models.GetUnix(),
	})

	if !rsp.Success {
		con.Error(c, "修改数据失败", "/admin/role/edit?id="+models.String(id))
	} else {
		con.Success(c, "修改数据成功", "/admin/role/edit?id="+models.String(id))
	}
}

func (con RoleController) Delete(c *gin.Context) {
	id, err1 := models.Int(c.Query("id"))
	if err1 != nil {
		fmt.Println("err1", err1)
		con.Error(c, "传入数据错误", "/admin/role")
	} else {
		rbacSrv := pbRole.NewRbacRoleService("rbac", models.MicroRbacClient)
		rsp, _ := rbacSrv.RoleDelete(context.Background(), &pbRole.RoleDeleteRequest{
			Id: int64(id),
		})
		if rsp.Success {
			con.Success(c, "删除数据成功", "/admin/role")
		} else {
			con.Error(c, "删除数据错误", "/admin/role")
		}
	}
}

func (con RoleController) Auth(c *gin.Context) {
	roleId, err := models.Int(c.Query("id"))
	if err != nil {
		con.Error(c, "传入数据错误", "/admin/role")
		return
	}
	rbacSrv := pbRole.NewRbacRoleService("rbac", models.MicroRbacClient)
	rsp, _ := rbacSrv.RoleAuth(context.Background(), &pbRole.RoleAuthRequest{
		RoleId: int64(roleId),
	})
	//accessList := []models.Access{}
	//models.DB.Where("module_id = ?", 0).Preload("AccessItem").Find(&accessList)
	//
	//roleAccess := []models.RoleAccess{}
	//models.DB.Where("role_id = ?", roleId).Find(&roleAccess)
	//
	//roleAccessMap := map[int]int{}
	//for _, v := range roleAccess {
	//	roleAccessMap[v.AccessId] = v.AccessId
	//}
	//
	//for i := 0; i < len(accessList); i++ {
	//	if _, ok := roleAccessMap[accessList[i].Id]; ok {
	//		accessList[i].Checked = true
	//	}
	//	for j := 0; j < len(accessList[i].AccessItem); j++ {
	//		if _, ok := roleAccessMap[accessList[i].AccessItem[j].Id]; ok {
	//			accessList[i].AccessItem[j].Checked = true
	//		}
	//	}
	//}

	c.HTML(http.StatusOK, "admin/role/auth.html", gin.H{
		"roleId":     roleId,
		"accessList": rsp.AccessList,
	})
}
func (con RoleController) DoAuth(c *gin.Context) {
	roleId, err1 := models.Int(c.PostForm("role_id"))
	if err1 != nil {
		fmt.Println(err1)
		con.Error(c, "传入数据错误", "/admin/role")
		return
	}
	accessIds := c.PostFormArray("access_node[]")
	//roleAccess := models.RoleAccess{}
	//models.DB.Where("role_id = ?", roleId).Delete(&roleAccess)
	//for _, v := range accessIds {
	//	roleAccess.RoleId = roleId
	//	roleAccess.AccessId, _ = models.Int(v)
	//	models.DB.Create(&roleAccess)
	//}

	rbacSrv := pbRole.NewRbacRoleService("rbac", models.MicroRbacClient)
	rsp, _ := rbacSrv.RoleDoAuth(context.Background(), &pbRole.RoleDoAuthRequest{
		RoleId:    int64(roleId),
		AccessIds: accessIds,
	})
	fmt.Println(roleId)
	fmt.Println(accessIds)
	if rsp.Success {
		con.Success(c, "授权成功", "/admin/role")
	} else {
		con.Error(c, rsp.Message, "/admin/role")
	}

}
