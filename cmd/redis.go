/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"log"

	"github.com/astaxie/beego/logs"
	"github.com/go-redis/redis/v8"
	tbRedis "github.com/ybooks240/ToolBox/pkg/redis"

	"github.com/spf13/cobra"
)

type opt struct {
	Mode     string
	Address  []string
	UserName string
	Password string
	DB       int
}

func redisError(err error, k interface{}) {
	if err == redis.Nil {
		logs.Warning("没有该key：", k)
		return
	}

	logs.Critical("出现了错误:", err)
}

var instance opt

// redisCmd represents the redis command
var redisCmd = &cobra.Command{
	Use:   "redis",
	Short: "使用该命令连接redis",
	Long:  `使用该命令连接redis，可以连接多种redis模式，单节点模式、sentinel模式以及cluster模式`,
	Example: `ToolBox redis --mode standalone --address <IP>:<PORT> -p <password>
	ToolBox redis --mode sentinel --address <IP>:<PORT>  --address  <IP>:<PORT>  --address <IP>:<PORT> -p <password>
	ToolBox redis --mode cluster --address <IP>:<PORT>  --address <IP>:<PORT>  --address <IP>:<PORT> --address <IP>:<PORT>  --address <IP>:<PORT>  --address <IP>:<PORT>  -p <password>`,
	// Aliases:   []string{"standalone", "sentinel", "cluster"},
	ValidArgs: []string{"standalone", "sentinel", "cluster"},
	// ValidArgs: []string{"mode", "address", "cluster"},

	Run: func(cmd *cobra.Command, args []string) {

		if len(args) < 2 {
			log.Fatal("需要输出参数：get or set")
			return
		}
		if len(args) < 3 {
			args = append(args, "")
		}
		operator := tbRedis.Operator{
			Opt: args[0],
			K:   args[1],
			V:   args[2],
		}

		logs.Info("你选择的redis模式是：%s\n", instance.Mode)
		logs.Info("将要执行的操作是%s,%s", operator.Opt, operator.K)
		logs.Info("将要连接到redis的地址是：%v\n", instance.Address)

		switch instance.Mode {
		case "standalone":
			sr := tbRedis.StandaloneRedis{
				Address:  instance.Address,
				Password: instance.Password,
				DB:       0,
			}
			result, err := sr.SetAndGet(operator)
			if err != nil {
				redisError(err, operator.K)
			}
			logs.Info("你%s的结果是：%s", operator.Opt, result)
		case "sentinel":
			sr := tbRedis.SentinelRedis{
				MasterName: "mymaster",
				Address:    instance.Address,
				UserName:   instance.UserName,
				PassWord:   instance.Password,
				DB:         instance.DB,
			}
			result, err := sr.SetAndGet(operator)
			if err != nil {
				redisError(err, operator.K)
			}
			logs.Info("你%s的结果是：%s", operator.Opt, result)
		case "cluster":
			cr := tbRedis.ClusterRedis{
				Address:  instance.Address,
				Password: instance.Password,
				Username: instance.UserName,
			}
			result, err := cr.SetAndGet(operator)
			if err != nil {
				redisError(err, operator.K)
			}
			logs.Info("你%s的结果是：%s", operator.Opt, result)
		default:
			cmd.Help()
			logs.Info("没有这个模式%s\n", instance.Mode)
		}
	},
}

func init() {
	rootCmd.AddCommand(redisCmd)

	redisCmd.Flags().StringVarP(&instance.Mode, "mode", "m", "standalone", "指定redis类型")
	// redisCmd.MarkFlagRequired("mode")
	redisCmd.Flags().StringArrayVar(&instance.Address, "address", []string{}, "指定iP列表")
	redisCmd.MarkFlagRequired("address")
	redisCmd.Flags().StringVarP(&instance.UserName, "username", "u", "", "指定用户名")
	redisCmd.Flags().StringVarP(&instance.Password, "password", "p", "", "指定密码")
}
