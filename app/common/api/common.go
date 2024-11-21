package api

import (
	"github.com/gin-gonic/gin"
	"ych/vgo/pkg/response"
)

func Test(ctx *gin.Context) {
	//global.Logger.Info("Request handled", zap.Any("data", "test"))
	//global.RedisCon.Set(ctx, "test", "test", 0)

	//response.Success(ctx, "哈哈哈", gin.H{
	//	"test": snow.SnowflakeService().Generate(),
	//})

	//// 生成token
	//res, err := auth.GenAdminToken(ctx, 1, []string{"哈哈"}, 1)
	//if err != nil {
	//	response.Fail(ctx, "失败", nil)
	//	return
	//}
	//response.Success(ctx, "登录成功", res)

	//client := queue.NewRedisClient()
	//defer queue.CloseRedisClient(client)
	//task, err := job.NewTestJob(42, "666666")
	//if err != nil {
	//	global.Logger.Error(fmt.Sprintf("could not create task: %v", err))
	//}
	//info, err := client.Enqueue(task)
	//if err != nil {
	//	global.Logger.Error(fmt.Sprintf("could not enqueue task: %v", err))
	//}
	//global.Logger.Info(fmt.Sprintf("enqueued task: id=%s queue=%s", info.ID, info.Queue))

	//err := global.DbCon.AutoMigrate(&model.notice{})
	//if err != nil {
	//	return
	//}
	response.Success(ctx, "vgo!!!", nil)
	return
}
