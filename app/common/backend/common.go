package backend

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"reflect"
	"strconv"
	"ych/vgo/internal/global"
	"ych/vgo/pkg/helper"
	"ych/vgo/pkg/response"
)

type ValidationRules struct {
	Create map[string]map[string]string `json:"Create"`
	Update map[string]map[string]string `json:"Update"`
}

type CRUDHandler struct {
	Model        interface{}      // 模型指针，用于创建实例
	Validate     *ValidationRules // 验证规则
	IndexWith    func(ctx *gin.Context, db *gorm.DB) *gorm.DB
	BeforeCreate func(ctx *gin.Context, model interface{}) error
	AfterCreate  func(ctx *gin.Context, model interface{}) error
	BeforeUpdate func(ctx *gin.Context, model interface{}) error
	AfterUpdate  func(ctx *gin.Context, model interface{}) error
	ShowWith     func(ctx *gin.Context, db *gorm.DB) *gorm.DB
	BeforeDelete func(ctx *gin.Context, model interface{}) error
	AfterDelete  func(ctx *gin.Context, model interface{}) error
}

func NewCRUDHandler(model interface{}, validate *ValidationRules) *CRUDHandler {
	return &CRUDHandler{
		Model:    model,
		Validate: validate,
	}
}

// Index 通用列表查询
func (h *CRUDHandler) Index(ctx *gin.Context) {
	pageNo, err := strconv.Atoi(ctx.DefaultQuery("pageNum", "1"))
	pageSize, err := strconv.Atoi(ctx.DefaultQuery("pageSize", "10"))
	if helper.HandleErr(ctx, err, "参数无效") {
		return
	}

	// 创建一个指向模型切片的指针
	modelSlicePtr := reflect.New(reflect.SliceOf(reflect.TypeOf(h.Model).Elem())).Interface()

	var total int64
	db := global.DbCon.Model(h.Model)

	if h.IndexWith != nil {
		db = h.IndexWith(ctx, db)
	}

	if err := db.Count(&total).Error; helper.HandleErr(ctx, err, "数据库查询失败") {
		return
	}

	if err := db.Order("id desc").Offset((pageNo - 1) * pageSize).Limit(pageSize).Find(modelSlicePtr).Error; helper.HandleErr(ctx, err, "数据库查询失败") {
		return
	}

	modelSlice := reflect.ValueOf(modelSlicePtr).Elem().Interface()

	totalPages := (int(total) + pageSize - 1) / pageSize
	response.Success(ctx, "成功", gin.H{
		"pageNum":  pageNo,
		"total":    totalPages,
		"pageSize": pageSize,
		"list":     modelSlice,
	}, nil)
}

// Create 通用创建
func (h *CRUDHandler) Create(ctx *gin.Context) {
	model := reflect.New(reflect.TypeOf(h.Model).Elem()).Interface()
	if h.Validate != nil && h.Validate.Create != nil {
		if !helper.BindAndValidate(ctx, model, h.Validate.Create) {
			return
		}
	} else {
		if err := helper.BindJSON(ctx, model); helper.HandleErr(ctx, err, "参数错误") {
			return
		}
	}
	if h.BeforeCreate != nil {
		if err := h.BeforeCreate(ctx, model); err != nil {
			response.Fail(ctx, "BeforeCreate-hook-failed", err.Error())
			return
		}
	}
	if err := global.DbCon.Create(model).Error; helper.HandleErr(ctx, err, "创建失败") {
		return
	}
	if h.AfterCreate != nil {
		if err := h.AfterCreate(ctx, model); err != nil {
			response.Fail(ctx, "AfterCreate-hook-failed", err.Error())
			return
		}
	}
	response.Success(ctx, "成功", model, nil)
}

// Update 通用编辑
func (h *CRUDHandler) Update(ctx *gin.Context) {
	model := reflect.New(reflect.TypeOf(h.Model).Elem()).Interface()
	if h.Validate != nil && h.Validate.Update != nil {
		if !helper.BindAndValidate(ctx, model, h.Validate.Update) {
			return
		}
	} else {
		if err := helper.BindJSON(ctx, model); helper.HandleErr(ctx, err, "参数错误") {
			return
		}
	}
	idField := reflect.ValueOf(model).Elem().FieldByName("ID")
	if !idField.IsValid() || idField.IsZero() {
		response.Fail(ctx, "记录 ID 无效", nil)
		return
	}

	existing := reflect.New(reflect.TypeOf(h.Model).Elem()).Interface()
	if err := global.DbCon.First(existing, idField.Interface()).Error; helper.HandleErr(ctx, err, "记录不存在") {
		return
	}

	if h.BeforeUpdate != nil {
		if err := h.BeforeUpdate(ctx, model); err != nil {
			response.Fail(ctx, "BeforeUpdate-hook-failed", err.Error())
			return
		}
	}

	if err := global.DbCon.Model(existing).Updates(model).Error; helper.HandleErr(ctx, err, "更新失败") {
		return
	}

	if h.AfterUpdate != nil {
		if err := h.AfterUpdate(ctx, model); err != nil {
			response.Fail(ctx, "AfterUpdate-hook-failed", err.Error())
			return
		}
	}

	response.Success(ctx, "成功", model, nil)
}

// Show 通用详情
func (h *CRUDHandler) Show(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		response.Fail(ctx, "无效的ID参数", nil)
		return
	}

	// 创建模型实例
	model := reflect.New(reflect.TypeOf(h.Model).Elem()).Interface()
	db := global.DbCon

	if h.ShowWith != nil {
		db = h.ShowWith(ctx, db)
	}
	// 查询数据库
	if err := db.First(model, id).Error; helper.HandleErr(ctx, err, "查询失败") {
		return
	}

	// 返回成功响应
	response.Success(ctx, "成功", model, nil)
}

// Delete 通用删除
func (h *CRUDHandler) Delete(ctx *gin.Context) {
	idStrings := ctx.QueryArray("id[]")
	var ids []int64
	for _, idStr := range idStrings {
		var id int64
		if _, err := fmt.Sscan(idStr, &id); err == nil {
			ids = append(ids, id)
		}
	}
	if len(ids) == 0 {
		response.Fail(ctx, "参数错误", "未提供有效的ID", nil)
		return
	}
	modelSlicePtr := reflect.New(reflect.SliceOf(reflect.TypeOf(h.Model).Elem())).Interface()
	if err := global.DbCon.Where("id in (?)", ids).Find(modelSlicePtr).Error; helper.HandleErr(ctx, err, "查询失败") {
		return
	}
	modelSlice := reflect.ValueOf(modelSlicePtr).Elem().Interface()
	if h.BeforeDelete != nil {
		sliceValue := reflect.ValueOf(modelSlice)
		for i := 0; i < sliceValue.Len(); i++ {
			model := sliceValue.Index(i).Interface()
			if err := h.BeforeDelete(ctx, model); err != nil {
				response.Fail(ctx, "BeforeDelete-hook-failed", err.Error())
				return
			}
		}
	}
	if err := global.DbCon.Where("id in (?)", ids).Delete(h.Model).Error; err != nil {
		response.Fail(ctx, "删除失败", err.Error())
		return
	}

	if h.AfterDelete != nil {
		sliceValue := reflect.ValueOf(modelSlice)
		for i := 0; i < sliceValue.Len(); i++ {
			model := sliceValue.Index(i).Interface()
			if err := h.AfterDelete(ctx, model); err != nil {
				response.Fail(ctx, "AfterDelete-hook-failed", err.Error())
				return
			}
		}
	}

	response.Success(ctx, "成功", nil, nil)
}
