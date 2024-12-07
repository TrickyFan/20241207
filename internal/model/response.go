package model

type BodyResponse struct {
	Code int32
	Data interface{}
	Msg  string
}

type UserRegisterResponse struct {
}

type UserLoginResponse struct {
	Token string
}

type CreateTaskResponse struct {
	TaskId      string
	VerifyToken string
}

type GetTaskDetailResponse struct {
	TaskId string
	Status int32
}

type DownloadTranslatedContentResponse struct {
	TaskId  string
	Content string
}
