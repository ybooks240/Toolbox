package tbRedis

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)

type tbRedis interface {
	NewClient() *redis.Client
	SetAndGet() (result interface{}, err error)
}

type Operator struct {
	Opt string
	K   string
	V   string
}

// 单节点操作
type StandaloneRedis struct {
	Address  []string
	Password string
	DB       int
}

func (sr StandaloneRedis) NewClient() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     sr.Address[0],
		Password: sr.Password,
		DB:       sr.DB,
	})
	return rdb
}

func (sr StandaloneRedis) SetAndGet(opt Operator) (result interface{}, err error) {
	rdb := sr.NewClient()
	defer rdb.Close()
	var (
		ctx = context.Background()
	)
	switch opt.Opt {
	case "get":
		get := rdb.Do(ctx, opt.Opt, opt.K)
		return get.Val(), get.Err()
	case "set":

		get := rdb.Do(ctx, opt.Opt, opt.K, opt.V)
		return get.Val(), get.Err()
	default:
		return nil, fmt.Errorf("不支持的方法:%s", opt.Opt)
	}

}

// 哨兵操作
type SentinelRedis struct {
	MasterName string
	Address    []string
	UserName   string
	PassWord   string
	DB         int
}

func (sr SentinelRedis) NewClient() *redis.Client {
	rdb := redis.NewFailoverClient(&redis.FailoverOptions{
		MasterName:    sr.MasterName,
		SentinelAddrs: sr.Address,
		Username:      sr.UserName,
		Password:      sr.PassWord,
		DB:            sr.DB,
	})
	return rdb
}

func (sr SentinelRedis) SetAndGet(opt Operator) (result interface{}, err error) {
	rdb := sr.NewClient()
	defer rdb.Close()
	var (
		ctx = context.Background()
	)
	switch opt.Opt {
	case "get":
		get := rdb.Do(ctx, opt.Opt, opt.K)
		return get.Val(), get.Err()
	case "set":
		get := rdb.Do(ctx, opt.Opt, opt.K, opt.V)
		return get.Val(), get.Err()
	default:
		return nil, fmt.Errorf("不支持的方法:%s", opt.Opt)
	}
}

// TODO 集群模式
type ClusterRedis struct {
	Address  []string
	Password string
	Username string
}

func (cr ClusterRedis) NewClient() *redis.ClusterClient {
	rdb := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs:    cr.Address,
		Username: cr.Username,
		Password: cr.Password,
	})
	return rdb
}

func (sr ClusterRedis) SetAndGet(opt Operator) (result interface{}, err error) {
	rdb := sr.NewClient()
	defer rdb.Close()
	var (
		ctx = context.Background()
	)
	switch opt.Opt {
	case "get":
		get := rdb.Get(ctx, opt.K)
		// get.Val()
		return get.Val(), get.Err()
	case "set":
		get := rdb.Do(ctx, opt.Opt, opt.K, opt.V)
		// if redis.Nil
		return get.Val(), get.Err()
	default:
		return nil, fmt.Errorf("不支持的方法:%s", opt.Opt)
	}
}
