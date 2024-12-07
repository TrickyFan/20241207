package dao

import (
	"context"
	"errors"
	"fmt"
	"github.com/redis/go-redis/v9"
)

type RedisDao struct {
	redis *redis.Client
}

func InitRedisDao() *RedisDao {
	return &RedisDao{
		redis: redis.NewClient(nil),
	}
}

func getLockUserNameKey(userName string) string {
	return fmt.Sprintf("user_name_lock_%s", userName)
}

func (dao *RedisDao) LockUserName(ctx context.Context, userName string) (err error) {
	lockKey := getLockUserNameKey(userName)
	cmd := dao.redis.Do(ctx, "set -nxp", lockKey, 5)
	err = cmd.Err()
	if err != nil {
		return
	}
	lock, err := cmd.Bool()
	if err != nil {
		return
	}
	if !lock {
		err = errors.New("lock error")
		return
	}
	return
}
