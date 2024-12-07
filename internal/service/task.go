package service

import (
	"context"

	pb "translate/api/translate/v1"
)

type TaskService struct {
	pb.UnimplementedTaskServer
}

func NewTaskService() *TaskService {
	return &TaskService{}
}

func (s *TaskService) CreateTask(ctx context.Context, req *pb.CreateTaskRequest) (*pb.CreateTaskResponse, error) {
	return &pb.CreateTaskResponse{}, nil
}
func (s *TaskService) ExecuteTranslationTask(ctx context.Context, req *pb.ExecuteTranslationTaskRequest) (*pb.ExecuteTranslationTaskRequestResponse, error) {
	return &pb.ExecuteTranslationTaskRequestResponse{}, nil
}
func (s *TaskService) DownloadTranslatedContent(ctx context.Context, req *pb.DownloadTranslatedContentRequest) (*pb.DownloadTranslatedContentResponse, error) {
	return &pb.DownloadTranslatedContentResponse{}, nil
}
func (s *TaskService) GetTaskDetail(ctx context.Context, req *pb.GetTaskDetailRequest) (*pb.GetTaskDetailResponse, error) {
	return &pb.GetTaskDetailResponse{}, nil
}
