package tbRedis

import (
	"fmt"

	"github.com/gomodule/redigo/redis"
)

type StandaloneRedis struct {
	Address []string
}

type tbRedis interface {
	NewRedisConn() redis.Conn
	Add(conn redis.Conn) error
}

func (sr StandaloneRedis) NewRedisConn() redis.Conn {
	address := sr.Address[0]
	conn, err := redis.Dial("tcp", address)
	if err != nil {
		fmt.Printf("connect redis error：%s", err)
		return nil
	}
	// defer conn.Close()
	return conn
}

// 测试针对值的增删改查

func (sr StandaloneRedis) Add(opt, k, v string) error {
	conn := sr.NewRedisConn()
	defer conn.Close()
	r, err := conn.Do(opt, k, v)
	if err != nil {
		fmt.Printf("redis 写入数据错误：%s", err)
		return err
	}
	// defer conn.Close()
	fmt.Println("redis写入值成功:", r)
	return nil
}
