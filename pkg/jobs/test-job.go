package job

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/hibiken/asynq"
	"go.uber.org/zap"
	"ych/vgo/internal/global"
)

const (
	TestJob = "test:job"
)

// TestJobPayload 任务参数
type TestJobPayload struct {
	UserID     int
	TemplateID string
}

// NewTestJob 创建一个新的测试任务
func NewTestJob(userID int, tmplID string) (*asynq.Task, error) {
	payload, err := json.Marshal(TestJobPayload{UserID: userID, TemplateID: tmplID})
	if err != nil {
		return nil, err
	}
	return asynq.NewTask(TestJob, payload), nil
}

// HandleTestJobTask 处理测试任务
func HandleTestJobTask(ctx context.Context, t *asynq.Task) error {
	var p TestJobPayload
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		global.Logger.Error("test:job-json.Unmarshal failed: %v", zap.Any("task-err", err), zap.Any("task-asynq", asynq.SkipRetry))
	}
	fmt.Println(fmt.Sprintf("test:job to User: user_id=%d, template_id=%s", p.UserID, p.TemplateID))
	return nil
}

// 使用示例：
//func Index(ctx *gin.Context) {
//client := queue.NewRedisClient()
//defer queue.CloseRedisClient(client)
//task, err := job.NewTestJob(42, "666666")
//if err != nil {
//global.Logger.Error(fmt.Sprintf("could not create task: %v", err))
//}
//info, err := client.Enqueue(task)
//if err != nil {
//global.Logger.Error(fmt.Sprintf("could not enqueue task: %v", err))
//}
//global.Logger.Info(fmt.Sprintf("enqueued task: id=%s queue=%s", info.ID, info.Queue))
//}
