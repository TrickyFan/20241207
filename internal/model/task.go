package model

type TaskModel struct {
	TaskId   int64
	UserName string
	Status   int32
}

const (
	TaskStatus_Init   = 1
	TaskStatus_Cancel = 0
	TaskStatus_Using  = 2
	TaskStatus_Done   = 3
)
