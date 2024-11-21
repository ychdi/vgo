package api

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"ych/vgo/app/notice/model"
	"ych/vgo/internal/global"
	"ych/vgo/pkg/helper"
	"ych/vgo/pkg/response"
)

// Index 列表
func Index(ctx *gin.Context) {
	var res []model.Notice
	var total int64
	pageNo, err := strconv.Atoi(ctx.DefaultQuery("pageNum", "1"))
	pageSize, err := strconv.Atoi(ctx.DefaultQuery("pageSize", "10"))
	if helper.HandleErr(ctx, err, "参数无效") {
		return
	}
	if err := global.DbCon.Model(&model.Notice{}).Count(&total).Error; helper.HandleErr(ctx, err, "数据库查询失败") {
		return
	}
	if err := global.DbCon.Order("id desc").Offset((pageNo - 1) * pageSize).Limit(pageSize).Find(&res).Error; helper.HandleErr(ctx, err, "数据库查询失败") {
		return
	}
	totalPages := (int(total) + pageSize - 1) / pageSize
	response.Success(ctx, "成功", gin.H{
		"pageNum":  pageNo,
		"total":    totalPages,
		"pageSize": pageSize,
		"list":     res,
	}, nil)
}

// Show 详情
func Show(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		response.Fail(ctx, "无效的ID参数", nil)
		return
	}
	var model model.Notice
	if err := global.DbCon.First(&model, id).Error; helper.HandleErr(ctx, err, "") {
		return
	}
	response.Success(ctx, "成功", model, nil)
}

func RegisterNoticeApiRoutes() {
	global.ApiRouter.GET("/notices", Index)
	global.ApiRouter.GET("/notices/:id", Show)
}
