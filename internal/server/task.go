package server

import (
	"context"

	v1 "translate/api/translate/v1"
	"translate/interal/dao"
)

// TaskService
type TaskService struct {
	dao *dao.UserDao
}

// NewTaskService
func NewTaskService() *TaskService {
	return &TaskService{}
}

// CreateTask
func (s *TaskService) CreateTask(ctx context.Context, in *v1.CreateTaskRequest) (*v1.CreateTaskReponse, error) {

}
