package dao

import (
	"context"
	"crypto/md5"
	"database/sql"
	"encoding/base64"
	"fmt"
)

const (
	_insertNewUser = "insert into user_table_%s (user_name, pwd) values (?, ?)"
)

type UserDao struct {
	db *sql.DB
}

func InitUserDao() {

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

func userTableNum(userName string) string {
	md5Hash := md5.New()
	md5Str := string(md5Hash.Sum([]byte(userName)))
	return md5Str[-2:]
}
