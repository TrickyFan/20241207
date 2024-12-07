package dao

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"translate/internal/model"
)

const (
	_insertATask = "insert into user_task_%s (user_name, status) values (?, ?)"
	_getTaskInfo = "select task_id, user_name, status from user_task_%s where user_name = ? and task_id = ?"
)

type TaskDao struct {
	db *sql.DB
}

func InitTaskDao() *TaskDao {
	return &TaskDao{
		db: &sql.DB{},
	}
}

func (dao *TaskDao) CreateTask(ctx context.Context, userName string) (record *model.TaskModel, err error) {
	if userName == "" {
		err = errors.New("参数错误")
		return
	}
	conn, err := dao.db.Conn(ctx)
	if err != nil {
		return
	}
	tableNum := userTableNum(userName)
	sqlStr := fmt.Sprintf(_insertATask, tableNum)
	result, err := conn.ExecContext(ctx, sqlStr, userName, model.TaskStatus_Init)
	if err != nil {
		return
	}
	taskId, err := result.LastInsertId()
	if err != nil {
		return
	}
	record = &model.TaskModel{
		TaskId:   taskId,
		UserName: userName,
		Status:   model.TaskStatus_Init,
	}
	return
}

func (dao *TaskDao) GetTask(ctx context.Context, taskId int64, userName string) (record *model.TaskModel, err error) {
	if userName == "" {
		err = errors.New("参数错误")
		return
	}
	conn, err := dao.db.Conn(ctx)
	if err != nil {
		return
	}
	tableNum := userTableNum(userName)
	sqlStr := fmt.Sprintf(_getTaskInfo, tableNum)
	row := conn.QueryRowContext(ctx, sqlStr, userName, taskId)
	if row.Err() != nil {
		return
	}
	record = &model.TaskModel{}
	err = row.Scan(record.TaskId, record.UserName, record.Status)
	if err != nil {
		return
	}
	return
}

func (dao *TaskDao) UpdateTaskStatus(ctx context.Context, taskId string, status string) (err error) {
	return
}
