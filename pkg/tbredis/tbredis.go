package tbredis

import (
	"fmt"

	"github.com/gomodule/redigo/redis"
)

type StandaloneRedis struct {
	IP   string
	Port string
}

func (sr StandaloneRedis) NewRedis() redis.Conn {
	address := sr.IP + ":" + sr.Port
	conn, err := redis.Dial("tcp", address)
	if err != nil {
		fmt.Printf("connect redis error：%s", err)
		return nil
	}
	defer conn.Close()
	return conn
}

// 测试针对值的增删改查

func Add(conn redis.Conn) {
	r, err := conn.Do("set", "cncpresult", "this is cncp auto test")
	if err != nil {
		fmt.Printf("redis 写入数据错误：%s", err)
	}
	fmt.Println("redis写入值成功:", r)
}
