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
	"fmt"

	tbRedis "github.com/ybooks240/ToolBox/pkg/redis"

	"github.com/spf13/cobra"
)

type opt struct {
	Mode     string
	Address  []string
	UserName string
	Password string
}

var instance opt

// redisCmd represents the redis command
var redisCmd = &cobra.Command{
	Use:   "redis",
	Short: "使用该命令连接redis",
	Long:  `使用该命令连接redis，可以连接多种redis模式，单节点模式、sentinel模式以及cluster模式`,
	Example: `ToolBox redis --mode standalone --address <IP>:<PORT> -p <password>
	ToolBox redis --mode sentinel --address <IP>:<PORT>,<IP>:<PORT>,<IP>:<PORT> -p <password>
	ToolBox redis --mode cluster --address <IP>:<PORT>,<IP>:<PORT>,<IP>:<PORT> -p <password>`,
	// Aliases:   []string{"standalone", "sentinel", "cluster"},
	ValidArgs: []string{"standalone", "sentinel", "cluster"},
	// ValidArgs: []string{"mode", "address", "cluster"},

	Run: func(cmd *cobra.Command, args []string) {

		// fmt.Println("你选择的redis模式是：", instance.Mode)
		// fmt.Println("将要连接到redis的地址是：", instance.Address)

		switch instance.Mode {
		case "standalone":
			fmt.Printf("正在使用%s模式", instance.Mode)
			tb := tbRedis.StandaloneRedis{
				Address: instance.Address,
			}
			tb.Add("set", "username", "xiaoml")
		case "sentinel":
			fmt.Printf("正在使用%s模式,还没适配，敬请期待", instance.Mode)
		case "cluster":
			fmt.Printf("正在使用%s模式，还没适配，敬请期待", instance.Mode)
		default:
			cmd.Help()
			fmt.Printf("没有这个模式%s\n", instance.Mode)
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

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// redisCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// redisCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
