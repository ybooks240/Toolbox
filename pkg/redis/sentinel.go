package tbRedis

import (
	"context"

	"github.com/go-redis/redis/v8"
)

type sentinelRedis struct {
	MasterName string
	Address    []string
	UserName   string
	PassWord   string
	DB         int
}

func (sr sentinelRedis) NewRedisCon() *redis.Client {
	sf := &redis.FailoverOptions{
		MasterName:    sr.MasterName,
		SentinelAddrs: sr.Address,
		Password:      sr.PassWord,
		DB:            sr.DB,
	}
	conn := redis.NewFailoverClient(sf)
	// defer conn.Close()
	return conn
}

type StandaloneRedis1 struct {
	Address  []string
	UserName string
	PassWord string
	DB       int
}

func (sr StandaloneRedis) testRedis() string {
	redisdb := redis.NewClient(&redis.Options{
		Addr:     sr.Address[0],
		Username: "",
		Password: "",
		DB:       1,
	})
	defer redisdb.Close()
	var ctx context.Context
	client := redisdb.ClusterInfo(ctx)
	name := client.FullName()
	return name
	// client.SetErr()
}
