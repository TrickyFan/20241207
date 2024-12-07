package service

import (
	"context"
	"errors"
	v1 "translate/api/v1"
	"translate/internal/dao"
	"translate/internal/model"
)

type TaskService struct {
	dao *dao.TaskDao
}

func NewTaskService() *TaskService {
	return &TaskService{
		dao: dao.InitTaskDao(),
	}
}

func (s *TaskService) CreateTask(ctx context.Context, req *v1.CreateTaskRequest) (response *v1.CreateTaskResponse, err error) {
	userName, ok := ctx.Value("user_name").(string)
	if !ok || userName == "" {
		err = errors.New("未登录")
		return
	}
	response = &v1.CreateTaskResponse{}
	task, err := s.dao.CreateTask(ctx, userName)
	if err != nil {
		return
	}
	response.TaskId = task.TaskId
	return response, nil
}

func (s *TaskService) ExecuteTranslationTask(ctx context.Context, req *v1.ExecuteTranslationTaskRequest) (response *v1.ExecuteTranslationTaskRequestResponse, err error) {
	userName, ok := ctx.Value("user_name").(string)
	if !ok || userName == "" {
		err = errors.New("未登录")
		return
	}
	if req.TaskId == 0 {
		err = errors.New("taskId无效")
		return
	}
	record, err := s.dao.GetTask(ctx, req.GetTaskId(), userName)
	if err != nil || record == nil || record.Status != model.TaskStatus_Using {
		err = errors.New("task无效")
		return
	}
	// TODO 增加对LLM API的调用 及 对关联关系的存储

	return &v1.ExecuteTranslationTaskRequestResponse{}, nil
}
func (s *TaskService) DownloadTranslatedContent(ctx context.Context, req *v1.DownloadTranslatedContentRequest) (response *v1.DownloadTranslatedContentResponse, err error) {
	userName, ok := ctx.Value("user_name").(string)
	if !ok || userName == "" {
		err = errors.New("未登录")
		return
	}
	//  TODO 增加与LLM关系的信息查找

	return &v1.DownloadTranslatedContentResponse{}, nil
}
func (s *TaskService) GetTaskDetail(ctx context.Context, req *v1.GetTaskDetailRequest) (response *v1.GetTaskDetailResponse, err error) {
	userName, ok := ctx.Value("user_name").(string)
	if !ok || userName == "" {
		err = errors.New("未登录")
		return
	}
	if req.TaskId == 0 {
		err = errors.New("taskId无效")
		return
	}
	record, err := s.dao.GetTask(ctx, req.GetTaskId(), userName)
	if err != nil || record == nil || record.Status != model.TaskStatus_Using {
		err = errors.New("task无效")
		return
	}
	return &v1.GetTaskDetailResponse{}, nil
}
