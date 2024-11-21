package upload

import (
	"fmt"
	"math/rand"
	"time"
)

// RandomFileName 生成随机文件名
func RandomFileName() string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	randomNum := r.Intn(10000)
	nano := time.Now().UnixNano()
	fileName := fmt.Sprintf("%d_%d", nano, randomNum)
	return fileName

}
