package tencent_cos

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/tencentyun/cos-go-sdk-v5"
	"mime/multipart"
	"net/http"
	"net/url"
	"path/filepath"
	Upload "ych/vgo/app/upload/common"
	"ych/vgo/internal/global"
	"ych/vgo/pkg/response"
)

func Run(ctx *gin.Context) {
	form, err := ctx.MultipartForm()
	if err != nil {
		fmt.Println("文件解析失败:", err)
		return
	}
	files := form.File["file"]
	var fileUrl string
	for _, file := range files {
		if fileUrl, err = UploadFile(file); err != nil {
			response.Fail(ctx, err.Error(), nil)
			return
		}
	}
	response.Success(ctx, "Cos文件上传成功", gin.H{
		"fileUrl": fileUrl,
	}, nil)
}

// UploadFile 上传文件到腾讯云对象存储
func UploadFile(file *multipart.FileHeader) (signedURL string, err error) {
	config := global.Config.CosConf
	u, _ := url.Parse(config.Endpoint)
	b := &cos.BaseURL{BucketURL: u}
	client := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  config.SecretID,
			SecretKey: config.SecretKey,
		},
	})
	localfile, err := file.Open()
	if err != nil {
		return "", fmt.Errorf("文件打开失败: %v", err)
	}
	defer localfile.Close()
	fileName := file.Filename
	fileType := filepath.Ext(fileName)
	if fileType == "" {
		return "", fmt.Errorf("文件解析失败: 文件没有扩展名")
	}
	// 生成文件保存的key
	key := Upload.RandomFileName() + fileType
	// 上传文件
	_, err = client.Object.Put(context.Background(), key, localfile, nil)
	if err != nil {
		return "", fmt.Errorf("文件上传失败: %v", err)
	}
	// 拼装文件的URL
	signedURL = fmt.Sprintf("%s/%s", u.String(), key)
	return signedURL, nil
}
