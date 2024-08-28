package System

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"vgo/core/db"
	"vgo/core/response"
	Model "vgo/model"
)

// GetBingBackgroundImage 获取必应每日背景图
func GetBingBackgroundImage(ctx *gin.Context) {
	isAbroad := false
	var url string
	if isAbroad {
		url = "https://bing.com/th?id=OHR.TateishiPark_ZH-CN9903501398_1920x1080.jpg&rf=LaDigue_1920x1080.jpg&pid=hp"
	} else {
		url = "https://cn.bing.com/th?id=OHR.TateishiPark_ZH-CN9903501398_1920x1080.jpg&rf=LaDigue_1920x1080.jpg&pid=hp"
	}
	response.Success(ctx, "成功", map[string]string{"url": url}, nil)
}

// getBingImageURL 联网获取必应每日背景图URL
func getBingImageURL(isAbroad bool, defaultURL string) (string, error) {
	var domain string
	if isAbroad {
		domain = "https://cn.bing.com"
	} else {
		domain = "https://bing.com"
	}
	resp, err := http.Get(domain + "/HPImageArchive.aspx?format=js&idx=0&n=1")
	if err != nil {
		return defaultURL, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return defaultURL, err
	}
	var content struct {
		Images []struct {
			URL string `json:"url"`
		} `json:"images"`
	}
	if err := json.Unmarshal(body, &content); err != nil {
		return defaultURL, err
	}
	if len(content.Images) > 0 && content.Images[0].URL != "" {
		return domain + content.Images[0].URL, nil
	}
	return defaultURL, nil
}

// GetInfo 获取管理员权限和菜单
func GetInfo(ctx *gin.Context) {
	userID := ctx.GetString("userID")
	// 无限极分类菜单结构
	var menus []Model.Menu
	db.Con().Order("sort desc").Find(&menus)
	menuTree := Model.BuildMenuTree(menus, 0)

	// 管理员信息
	var adminUser Model.AdminUser
	db.Con().First(&adminUser, "id = ?", userID)

	response.Success(ctx, "成功", gin.H{
		"codes":   []string{"*"},
		"roles":   []string{"superAdmin"},
		"routers": menuTree,
		"user":    adminUser,
	}, nil)
}
