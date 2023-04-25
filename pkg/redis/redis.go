package redis

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

func SaveKey(ctx context.Context, key string, value interface{}, expire time.Duration) (*redis.StatusCmd, error) {
	rdb := GetRedis()
	redisDB, err := rdb.GetConnection()
	if err != nil {
		return nil, err
	}
	defer rdb.Release(redisDB)

	return redisDB.Set(ctx, key, value, expire), nil

}

func SaveOnTable(ctx context.Context, key string, value ...interface{}) (*redis.IntCmd, error) {
	rdb := GetRedis()
	redisDB, err := rdb.GetConnection()
	if err != nil {
		return nil, err
	}
	defer rdb.Release(redisDB)

	return redisDB.HSet(ctx, key, value), nil

}

func GetFromTable(ctx context.Context, key string, field string) (*redis.StringCmd, error) {
	rdb := GetRedis()
	redisDB, err := rdb.GetConnection()
	if err != nil {
		return nil, err
	}

	defer rdb.Release(redisDB)

	return redisDB.HGet(ctx, key, field), nil
}

func GetValue(ctx context.Context, key string) (*redis.StringCmd, error) {
	rdb := GetRedis()
	redisDB, err := rdb.GetConnection()
	if err != nil {
		return nil, err
	}

	defer rdb.Release(redisDB)

	return redisDB.Get(ctx, key), nil
}

func GetTTL(ctx context.Context, key string) (*redis.DurationCmd, error) {
	rdb := GetRedis()
	redisDB, err := rdb.GetConnection()
	if err != nil {
		return nil, err
	}

	defer rdb.Release(redisDB)

	return redisDB.TTL(ctx, key), nil
}

func GetAll(ctx context.Context, key string) (*redis.StringStringMapCmd, error) {
	rdb := GetRedis()
	redisDB, err := rdb.GetConnection()
	if err != nil {
		return nil, err
	}

	defer rdb.Release(redisDB)
	return redisDB.HGetAll(ctx, key), nil
}

func DeleteKeys(ctx context.Context, key ...string) (*redis.IntCmd, error) {
	rdb := GetRedis()
	redisDB, err := rdb.GetConnection()
	if err != nil {
		return nil, err
	}
	defer rdb.Release(redisDB)
	return redisDB.Del(ctx, key...), nil
}

func DeleteFromTable(ctx context.Context, key string, field ...string) (*redis.IntCmd, error) {
	rdb := GetRedis()
	redisDB, err := rdb.GetConnection()
	if err != nil {
		return nil, err
	}

	defer rdb.Release(redisDB)
	return redisDB.HDel(ctx, key, field...), nil
}
