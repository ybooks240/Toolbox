package tbRedis

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)

// 单节点操作
type tbRedis interface {
	NewClient() *redis.Client
	SetAndGet() (result interface{}, err error)
}

type Operator struct {
	Opt string
	K   string
	V   string
}

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
	case "set":
		get := rdb.Do(ctx, opt.Opt, opt.K, opt.V)
		if err := get.Err(); err != nil {
			if err != nil {
				return nil, err
			}
			//TODO 外面再recover
			panic(err)
		}
		return get.Val().(string), nil
	case "get":
		get := rdb.Do(ctx, opt.Opt, opt.K)
		return get.Val().(string), nil
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

	// rdb := redis.NewFailoverClient(&redis.FailoverOptions{
	// 	MasterName: "mymaster",
	// 	SentinelAddrs: []string{
	// 		"172.16.123.137:26379",
	// 		"172.16.123.138:26379",
	// 		"172.16.123.139:26379",
	// 	},
	// 	Username: "redis_cncp",
	// 	Password: "wXGwskVXi2vCBSld",
	// 	DB:       0,
	// })
	rdb := redis.NewFailoverClient(&redis.FailoverOptions{
		MasterName:    sr.MasterName,
		SentinelAddrs: sr.Address,
		Username:      sr.UserName,
		Password:      sr.PassWord,
		DB:            sr.DB,
	})
	return rdb
}

func (sr SentinelRedis) GET() {
	rdb := sr.NewClient()
	defer rdb.Close()
	opt := "set"
	k := "test2"
	v := "testtes"
	var ctx = context.Background()
	// rdb.Ping(ctx)
	get := rdb.Do(ctx, opt, k, v)
	fmt.Println(get.Val())
}

func (sr SentinelRedis) SetAndGet(opt Operator) (result interface{}, err error) {
	rdb := sr.NewClient()
	defer rdb.Close()
	var (
		ctx = context.Background()
	)
	switch opt.Opt {
	case "set":
		get := rdb.Do(ctx, opt.Opt, opt.K, opt.V)
		if err := get.Err(); err != nil {
			if err == redis.Nil {
				return nil, err
			}
			//TODO 外面再recover
			panic(err)
		}
		return get.Val().(string), nil
	case "get":
		get := rdb.Do(ctx, opt.Opt, opt.K)
		return get.Val().(string), nil
	default:
		return nil, fmt.Errorf("不支持的方法:%s", opt.Opt)
	}
}

// // TODO 集群模式
// func (sr cluster) NewClient1 {
// 	rdb := redis.NewClusterClient(&redis.ClusterOptions{
// 		Addrs: []string{":7000", ":7001", ":7002", ":7003", ":7004", ":7005"},

// 		// To route commands by latency or randomly, enable one of the following.
// 		//RouteByLatency: true,
// 		//RouteRandomly: true,
// 	})
// }
