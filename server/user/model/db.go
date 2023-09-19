package model

import (
	"context"

	"github.com/redis/go-redis/v9"
)

func CheckUsernameExists(ctx context.Context, client *redis.ClusterClient, username string) (bool, error) {
	// 查询 Redis 中是否存在该用户名
	exists, err := client.Exists(ctx, "username_"+username).Result()
	if err != nil {
		return false, err
	}
	return exists == 1, nil
}

func RegisterUser(ctx context.Context, client *redis.ClusterClient, username string) error {
	// 存储用户名到 Redis 中
	_, err := client.Set(ctx, "username_"+username, "registered", 0).Result()
	return err
}
