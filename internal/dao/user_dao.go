package dao

import (
	"context"
	"crypto/md5"
	"database/sql"
	"encoding/base64"
	"fmt"
	"translate/internal/model"
)

const (
	_insertNewUser    = "insert into user_table_%s (user_name, pwd) values (?, ?)"
	_selectByUserName = "select user_id, user_name, pwd from user_table_%s where user_name = ?"
)

type UserDao struct {
	db *sql.DB
}

func InitUserDao() *UserDao {
	return &UserDao{
		db: &sql.DB{},
	}
}

func (dao *UserDao) CreateUser(ctx context.Context, userName string, pwd string) (err error) {
	var pwdStorageBytes []byte
	base64.NewEncoding("").Encode(pwdStorageBytes, []byte(pwd))
	conn, err := dao.db.Conn(ctx)
	if err != nil {
		return
	}

	tableNum := userTableNum(userName)
	sqlStr := fmt.Sprintf(_insertNewUser, tableNum)
	_, err = conn.ExecContext(ctx, sqlStr, userName, string(pwdStorageBytes))
	if err != nil {
		return
	}
	return
}

func (dao *UserDao) GetUserInfo(ctx context.Context, userName string) (record *model.UserModel, err error) {
	tableNum := userTableNum(userName)
	sqlStr := fmt.Sprintf(_selectByUserName, tableNum)
	conn, err := dao.db.Conn(ctx)
	if err != nil {
		return
	}
	row := conn.QueryRowContext(ctx, sqlStr, userName)
	if row.Err() != nil {
		return
	}
	record = &model.UserModel{}

	err = row.Scan(record.UserId, record.UserName, record.Pwd)
	if err != nil {
		return
	}
	return
}

func userTableNum(userName string) string {
	md5Hash := md5.New()
	md5Str := string(md5Hash.Sum([]byte(userName)))
	return md5Str[-2:]
}
