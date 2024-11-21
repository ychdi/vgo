package api

import (
	"ych/vgo/app/article/model"
	"ych/vgo/app/common/backend"
	"ych/vgo/internal/global"
)

func RegisterArticleRoutes() {
	articleHandler := backend.NewCRUDHandler(&model.Article{}, nil)
	global.ApiRouter.GET("/articles", articleHandler.Index)
	global.ApiRouter.GET("/articles/:id", articleHandler.Show)
}
