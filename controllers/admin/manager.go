package admin

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"shop/models"
	pbManager "shop/proto/rbacManager"
	pbRole "shop/proto/rbacRole"
	"strings"
)

type ManagerController struct {
	BaseController
}

func (con ManagerController) Index(c *gin.Context) {

	//managerList := []models.Manager{}
	//models.DB.Preload("Role").Find(&managerList)
	managerSrv := pbManager.NewRbacManagerService("rbac", models.MicroRbacClient)
	res, err := managerSrv.ManagerGet(context.Background(), &pbManager.ManagerGetRequest{})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%v", res.ManagerList)
	c.HTML(http.StatusOK, "admin/manager/index.html", gin.H{
		"managerList": res.ManagerList,
	})
}

func (con ManagerController) Add(c *gin.Context) {
	//
	rbacSrv := pbRole.NewRbacRoleService("rbac", models.MicroRbacClient)
	rsp, err := rbacSrv.RoleGet(context.Background(), &pbRole.RoleGetRequest{})
	if err != nil {
		return
	}
	c.HTML(http.StatusOK, "admin/manager/add.html", gin.H{
		"roleList": rsp.RoleList,
	})
}

func (con ManagerController) DoAdd(c *gin.Context) {
	roleId, err := models.Int(c.PostForm("role_id"))
	if err != nil {
		con.Error(c, "传入数据错误", "/admin/manager/add")
		return
	}
	username := strings.Trim(c.PostForm("username"), " ")
	password := strings.Trim(c.PostForm("password"), " ")
	email := strings.Trim(c.PostForm("email"), " ")
	mobile := strings.Trim(c.PostForm("mobile"), " ")

	if len(username) < 2 || len(password) < 6 {
		con.Error(c, "用户名或者密码长度不合法", "/admin/manager/add")
		return
	}

	managerSrv := pbManager.NewRbacManagerService("rbac", models.MicroRbacClient)
	res, err := managerSrv.ManagerGet(context.Background(), &pbManager.ManagerGetRequest{
		Username: username,
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	if len(res.ManagerList) > 0 {
		con.Error(c, "管理员已经存在", "/admin/manager/add")
		return
	}

	//manager := models.Manager{
	//	Username: username,
	//	Password: models.Md5(password),
	//	Email:    email,
	//	Mobile:   mobile,
	//	RoleId:   roleId,
	//	Status:   1,
	//	AddTime:  int(models.GetUnix()),
	//}
	resManagerAdd, err := managerSrv.ManagerAdd(context.Background(), &pbManager.ManagerAddRequest{
		Username: username,
		Password: models.Md5(password),
		Mobile:   mobile,
		Email:    email,
		Status:   1,
		RoleId:   int64(roleId),
		AddTime:  models.GetUnix(),
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	if !resManagerAdd.Success {
		con.Error(c, "增加管理员失败", "/admin/manager/add")
		return
	}
	con.Success(c, "增加管理员成功", "/admin/manager")
}

func (con ManagerController) Edit(c *gin.Context) {

	id, err := models.Int(c.Query("id"))
	if err != nil {
		con.Error(c, "传入数据错误", "/admin/manager")
		return
	}

	managerSrv := pbManager.NewRbacManagerService("rbac", models.MicroRbacClient)
	res, _ := managerSrv.ManagerGet(context.Background(), &pbManager.ManagerGetRequest{
		Id: int64(id),
	})

	//roleList := []models.Role{}
	//models.DB.Find(&roleList)
	RoleSrv := pbRole.NewRbacRoleService("rbac", models.MicroRbacClient)
	rsp, err := RoleSrv.RoleGet(context.Background(), &pbRole.RoleGetRequest{})
	if err != nil {
		return
	}
	c.HTML(http.StatusOK, "admin/manager/edit.html", gin.H{
		"manager":  res.ManagerList[0],
		"roleList": rsp.RoleList,
	})
}

func (con ManagerController) DoEdit(c *gin.Context) {
	id, err1 := models.Int(c.PostForm("id"))
	if err1 != nil {
		con.Error(c, "传入数据错误", "/admin/manager")
		return
	}
	roleId, err1 := models.Int(c.PostForm("role_id"))
	if err1 != nil {
		con.Error(c, "传入数据错误", "/admin/manager")
		return
	}
	username := strings.Trim(c.PostForm("username"), " ")
	password := strings.Trim(c.PostForm("password"), " ")
	email := strings.Trim(c.PostForm("email"), " ")
	mobile := strings.Trim(c.PostForm("mobile"), " ")

	if len(mobile) > 11 {
		con.Error(c, "电话长度不合法", "/admin/manager/edit?id="+models.String(id))
		return
	}

	//manager := models.Manager{Id: id}
	//models.DB.Find(&manager)
	//manager.Username = username
	//manager.Email = email
	//manager.Mobile = mobile
	//manager.RoleId = roleId

	if password != "" {
		if len(password) < 6 {
			con.Error(c, "密码长度不合法", "/admin/manager/edit?id="+models.String(id))
			return
		}
		password = models.Md5(password)
	}
	fmt.Println("password", password)
	managerSrv := pbManager.NewRbacManagerService("rbac", models.MicroRbacClient)
	rsp, _ := managerSrv.ManagerEdit(context.Background(), &pbManager.ManagerEditRequest{
		Id:       int64(id),
		Username: username,
		Password: password,
		Mobile:   mobile,
		Email:    email,
		RoleId:   int64(roleId),
	})
	//err2 := models.DB.Save(&manager).Error
	if !rsp.Success {
		con.Error(c, "修改管理员失败", "/admin/manager/edit?id="+models.String(id))
		return
	}
	con.Success(c, "修改管理员成功", "/admin/manager")
}

func (con ManagerController) Delete(c *gin.Context) {
	id, err1 := models.Int(c.Query("id"))
	if err1 != nil {
		fmt.Println("err1", err1)
		con.Error(c, "传入数据错误", "/admin/manager")
	} else {
		managerSrv := pbManager.NewRbacManagerService("rbac", models.MicroRbacClient)
		rsp, _ := managerSrv.ManagerDelete(context.Background(), &pbManager.ManagerDeleteRequest{
			Id: int64(id),
		})
		if rsp.Success {
			con.Success(c, "删除数据成功", "/admin/manager")
		} else {
			con.Error(c, "传入数据错误", "/admin/manager")
		}
	}
}
