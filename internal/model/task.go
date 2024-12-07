package model

type TaskModel struct {
	TaskId   int64
	UserName string
	Status   int32
	LLMId    string
}

const (
	TaskStatus_Init   = 1
	TaskStatus_Cancel = 0
	TaskStatus_Using  = 2
	TaskStatus_Done   = 3
)

type LLMResult struct {
	LLMId       string
	Status      int32
	ReqContent  string
	RespContent string
}
