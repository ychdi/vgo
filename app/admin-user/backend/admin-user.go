package backend

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"strconv"
	"ych/vgo/app/admin-user/model"
	"ych/vgo/app/common/backend"
	"ych/vgo/internal/global"
	"ych/vgo/internal/pkg/middleware/auth"
	"ych/vgo/pkg/helper"
	"ych/vgo/pkg/response"
)

// Login 登录
func Login(ctx *gin.Context) {
	var user model.AdminUser
	rules := map[string]map[string]string{
		"UserName": {"required": "用户名不能为空"},
		"Password": {"required": "密码不能为空"},
	}
	if !helper.BindAndValidate(ctx, &user, rules) {
		return
	}
	// 查询用户数据
	var dbUser model.AdminUser
	if err := global.DbCon.Where("username = ?", user.UserName).First(&dbUser).Error; err != nil {
		response.Fail(ctx, "用户名或密码错误", "用户名或密码错误001", nil)
		return
	}
	// 验证密码
	if err := bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(user.Password)); err != nil {
		response.Fail(ctx, "用户名或密码错误", "用户名或密码错误002", nil)
		return
	}
	// 获取角色
	roles, err := global.Rbac.GetRolesForUser(strconv.FormatUint(dbUser.ID, 10))
	if err != nil {
		response.Fail(ctx, "获取用户角色失败", err.Error(), nil)
		return
	}
	// 生成 token
	res, err := auth.GenAdminToken(ctx, dbUser.ID, roles, dbUser.Super)
	if err != nil {
		response.Fail(ctx, "生成 token 失败", err.Error(), nil)
		return
	}
	response.Success(ctx, "登录成功", res)
}

// LogOut 退出登录
func LogOut(ctx *gin.Context) {
	userID := ctx.GetUint64("userID")
	err := auth.DelAdminToken(ctx, userID)
	if err != nil {
		return
	}
	response.Success(ctx, "退出成功", nil)
}

// SetRole 设置角色
func SetRole(ctx *gin.Context) {
	var codes struct {
		ID    uint64   `json:"id"`
		Roles []string `json:"roles"`
	}
	if err := helper.BindJSON(ctx, &codes); err != nil {
		response.Fail(ctx, "参数错误", err.Error(), nil)
		return
	}
	userID := codes.ID
	enforcer := global.Rbac
	_, err := enforcer.DeleteRolesForUser(strconv.FormatUint(userID, 10))
	if err != nil {
		response.Fail(ctx, err.Error(), nil)
		return
	}
	for _, role := range codes.Roles {
		_, err2 := enforcer.AddRoleForUser(strconv.FormatUint(userID, 10), role)
		if err2 != nil {
			response.Fail(ctx, err2.Error(), nil)
			return
		}
	}
	response.Success(ctx, "设置成功", nil, nil)
}

// GetRole 获取角色
func GetRole(ctx *gin.Context) {
	var codes struct {
		ID uint64 `json:"id"`
	}
	if err := helper.BindJSON(ctx, &codes); err != nil {
		response.Fail(ctx, "参数错误", err.Error(), nil)
		return
	}
	userID := codes.ID
	enforcer := global.Rbac
	roles, err := enforcer.GetRolesForUser(strconv.FormatUint(userID, 10))
	if err != nil {
		response.Fail(ctx, "获取错误", err.Error(), nil)
		return
	}
	response.Success(ctx, "获取成功", roles, nil)
}

func RegisterAdminUserRoutes() {
	global.Engine.Group("/backend").POST("/user/login", Login)
	handler := backend.NewCRUDHandler(&model.AdminUser{}, nil)

	handler.BeforeCreate = func(ctx *gin.Context, m interface{}) error {
		adminUser, ok := m.(*model.AdminUser)
		if !ok {
			err := fmt.Errorf("类型断言失败")
			response.Fail(ctx, "类型断言失败", err.Error(), nil)
			return err
		}
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(adminUser.Password), bcrypt.DefaultCost)
		if err != nil {
			response.Fail(ctx, "密码哈希失败", err.Error(), nil)
			return err
		}
		adminUser.Password = string(hashedPassword)
		fmt.Println("创建前")
		return nil
	}
	handler.AfterCreate = func(ctx *gin.Context, model interface{}) error {
		fmt.Println("创建后")
		return nil
	}
	handler.BeforeUpdate = func(ctx *gin.Context, model interface{}) error {
		fmt.Println("更新前")
		return nil
	}
	handler.AfterUpdate = func(ctx *gin.Context, model interface{}) error {
		fmt.Println("更新后")
		return nil
	}
	handler.BeforeDelete = func(ctx *gin.Context, model interface{}) error {
		fmt.Println("删除前")
		return nil
	}
	handler.AfterDelete = func(ctx *gin.Context, model interface{}) error {
		fmt.Println("删除后")
		return nil
	}

	global.BackendRouter.GET("/users", handler.Index)
	global.BackendRouter.POST("/users", handler.Create)
	global.BackendRouter.PUT("/users", handler.Update)
	global.BackendRouter.GET("/users/:id", handler.Show)
	global.BackendRouter.DELETE("/users", handler.Delete)

	global.BackendRouter.POST("/users/set/role", SetRole)
	global.BackendRouter.POST("/users/get/role", GetRole)
	global.BackendRouter.POST("/users/logout", LogOut)
}
