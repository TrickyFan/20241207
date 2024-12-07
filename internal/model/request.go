package model

type Headers struct {
	AppId       string
	Token       string
	ContentType string
}

type UserRegisterRequest struct {
	UserName string
	Pwd      string
}

type UserLoginRequest struct {
	UserName string
	Pwd      string
}

type TaskCreateRequest struct {
}

type ExecuteTranslationRequest struct {
	TaskId      string
	Content     string
	VerifyToken string
}

type GetTaskDetailRequest struct {
	TaskId string
}

type DownloadTranslatedContentRequest struct {
	TaskId      string
	VerifyToken string
}
