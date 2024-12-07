package service

import (
	"context"
	"errors"
	v1 "translate/api/v1"
	"translate/internal/dao"
)

type UserService struct {
	dao   *dao.UserDao
	redis *dao.RedisDao
}

func NewUserService() *UserService {
	return &UserService{
		dao:   dao.InitUserDao(),
		redis: dao.InitRedisDao(),
	}
}

func (s *UserService) RegisterUser(ctx context.Context, req *v1.RegisterUserRequest) (response *v1.RegisterUserResponse, err error) {
	response = new(v1.RegisterUserResponse)
	record, err := s.dao.GetUserInfo(ctx, req.GetUserName())
	if err != nil {
		return
	}
	if record != nil && record.UserId != 0 {
		err = errors.New("该userName已占用")
		return
	}
	err = s.redis.LockUserName(ctx, req.GetUserName())
	if err != nil {
		err = errors.New("该userName已占用")
		return
	}
	err = s.dao.CreateUser(ctx, req.GetUserName(), req.GetPwd())
	if err != nil {
		return
	}
	return response, nil
}
func (s *UserService) UserLogin(ctx context.Context, req *v1.UserLoginRequest) (response *v1.UserLoginResponse, err error) {
	response = &v1.UserLoginResponse{}
	if req == nil || req.UserName == "" {
		err = errors.New("必要参数缺失")
		return
	}
	record, err := s.dao.GetUserInfo(ctx, req.GetUserName())
	if err != nil {
		return
	}
	if record == nil || record.UserId == 0 {
		err = errors.New("未找到账号")
		return
	}
	return response, nil
}
