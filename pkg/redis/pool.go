package redis

import (
	"context"
	"github.com/dozheiny/backend-template/config"

	"github.com/go-redis/redis/v8"
)

var redisConnection *Redis

type Redis struct {
	ctx       context.Context
	redisPool chan *redis.Client
}

func (pool *Redis) GetConnection() (*redis.Client, error) {
	if pool.redisPool == nil {
		if err := pool.initialize(); err != nil {
			return nil, err
		}
	}
	address, err := config.Get("ASANBILIT_REDIS_ADDR")
	if err != nil {
		return nil, err
	}

	if len(pool.redisPool) == 0 {
		client := redis.NewClient(&redis.Options{
			Addr:     address,
			Password: "",
			DB:       1,
		})
		pool.redisPool <- client
	}
	con := <-pool.redisPool
	if _, err := con.Ping(pool.ctx).Result(); err != nil {
		return nil, err
	}

	return con, nil
}

func GetRedis() *Redis {
	if redisConnection == nil {
		redisConnection = new(Redis)
	}
	return redisConnection
}

func (pool *Redis) Release(con *redis.Client) {
	if len(pool.redisPool) > 500 {
		_ = con.Close()
	} else {
		pool.redisPool <- con
	}
}

func (pool *Redis) initialize() error {

	address, err := config.Get("ASANBILIT_REDIS_ADDR")
	if err != nil {
		return err
	}

	pool.ctx = context.Background()
	pool.redisPool = make(chan *redis.Client, 1000)
	for range [4]int{} {
		client := redis.NewClient(&redis.Options{
			Addr:     address,
			Password: "",
			DB:       1,
			PoolSize: 2,
		})
		pool.redisPool <- client
	}

	return nil
}
