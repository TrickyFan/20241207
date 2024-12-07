package service

import (
	"context"
	"errors"
	v1 "translate/api/v1"
	"translate/internal/dao"
	"translate/internal/model"
)

type TaskService struct {
	dao    *dao.TaskDao
	llmDao *dao.LLMDao
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
	// 增加对LLM API的调用 及 对关联关系的存储
	result, err := s.llmDao.Translate(ctx, req.Content)
	if err != nil {
		return
	}
	if result == nil || result.LLMId == "" {
		err = errors.New("请求失败")
		return
	}
	err = s.dao.UpdateTaskLLM(ctx, req.TaskId, userName, result.LLMId)
	if err != nil {
		return
	}
	return &v1.ExecuteTranslationTaskRequestResponse{}, nil
}
func (s *TaskService) DownloadTranslatedContent(ctx context.Context, req *v1.DownloadTranslatedContentRequest) (response *v1.DownloadTranslatedContentResponse, err error) {
	userName, ok := ctx.Value("user_name").(string)
	if !ok || userName == "" {
		err = errors.New("未登录")
		return
	}
	if req.TaskId == 0 {
		err = errors.New("taskId无效")
		return
	}
	// 增加与LLM关系的信息查找
	record, err := s.dao.GetTask(ctx, req.GetTaskId(), userName)
	if err != nil || record == nil || record.Status == model.TaskStatus_Cancel {
		err = errors.New("task无效")
		return
	}
	result, err := s.llmDao.GetResult(ctx, record.LLMId)
	if err != nil {
		return
	}
	response = &v1.DownloadTranslatedContentResponse{}
	response.Content = result.RespContent
	return response, nil
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
