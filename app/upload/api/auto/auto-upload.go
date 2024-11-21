package auto

import (
	"fmt"
	"github.com/gabriel-vasile/mimetype"
	"github.com/gin-gonic/gin"
	"io"
	"ych/vgo/app/upload/api/ali-oss"
	"ych/vgo/app/upload/api/local"
	"ych/vgo/app/upload/api/tencent-cos"
	"ych/vgo/internal/global"
	"ych/vgo/pkg/response"
)

func Run(ctx *gin.Context) {
	file, err := ctx.FormFile("file")
	if err != nil {
		response.Fail(ctx, "获取文件失败", err.Error(), nil)
		return
	}

	uploadType := ctx.PostForm("type")
	if uploadType == "" {
		response.Fail(ctx, "上传类型标识不能为空", nil)
		return
	}
	defTypes := map[string]struct {
		size  int64
		mimes []string
	}{
		"any": {
			size:  50 * 1024 * 1024, // 50MB
			mimes: []string{"image/jpeg", "image/png", "image/jpg", "video/mp4"},
		},
		"img": {
			size:  10 * 1024 * 1024, // 10MB
			mimes: []string{"image/jpeg", "image/png", "image/jpg"},
		},
		"video": {
			size:  10 * 1024 * 1024, // 10MB
			mimes: []string{"video/mp4", "video/avi"},
		},
	}
	if _, exists := defTypes[uploadType]; !exists {
		response.Fail(ctx, "上传类型标识不支持", nil)
		return
	}
	defType := defTypes[uploadType]
	//检查文件大小
	if file.Size > defType.size {
		response.Fail(ctx, fmt.Sprintf("文件大小不能超过 %d MB", defType.size/(1024*1024)), nil)
		return
	}
	//读取文件
	src, err := file.Open()
	if err != nil {
		response.Fail(ctx, "无法打开文件", err.Error())
		return
	}
	defer src.Close()
	// 读取文件内容
	fileContent, err := io.ReadAll(src)
	if err != nil {
		response.Fail(ctx, "无法读取文件内容", err.Error())
		return
	}
	// 使用 mimetype 库获取文件的真实类型
	mime := mimetype.Detect(fileContent)
	fileType := mime.String()
	validMime := false
	for _, mimeType := range defType.mimes {
		if mimeType == fileType {
			validMime = true
			break
		}
	}
	if !validMime {
		response.Fail(ctx, "不支持的文件类型", nil)
		return
	}

	sysUploadType := global.Config.App.UploadType
	switch sysUploadType {
	case "local":
		respUrl, err := local.UploadFile("img", file)
		if err != nil {
			return
		}
		imgDomain := global.Config.App.FileDomain
		response.Success(ctx, "Local文件上传成功", gin.H{
			"fileUrl": imgDomain + respUrl,
		}, nil)
		break
	case "oss":
		respUrl, err := ali_oss.UploadFile(file)
		if err != nil {
			return
		}
		response.Success(ctx, "Oss文件上传成功", gin.H{
			"fileUrl": respUrl,
		}, nil)
		break
	case "cos":
		respUrl, err := tencent_cos.UploadFile(file)
		if err != nil {
			return
		}
		response.Success(ctx, "Cos文件上传成功", gin.H{
			"fileUrl": respUrl,
		}, nil)
		break
	}
}
