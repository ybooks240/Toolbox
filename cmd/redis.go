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

	"github.com/spf13/cobra"
)

// redisCmd represents the redis command
var redisCmd = &cobra.Command{
	Use:   "redis",
	Short: "使用该命令连接redis",
	Long:  `使用该命令连接redis，可以连接多种redis模式，单节点模式、sentinel模式以及cluster模式`,
	Example: `ToolBox redis --mode standalone --address <IP>:<PORT> -p <password>
	ToolBox redis --mode sentinel --address <IP>:<PORT>,<IP>:<PORT>,<IP>:<PORT> -p <password>
	ToolBox redis --mode cluster --address <IP>:<PORT>,<IP>:<PORT>,<IP>:<PORT> -p <password>`,
	Aliases:   []string{"standalone", "sentinel", "cluster"},
	ValidArgs: []string{"mode", "address", "cluster"},

	ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		_, err := cmd.Flags().GetString("mode")
		if err != nil {
			return nil, cobra.ShellCompDirectiveError
		}

		// if

		return args, cobra.ShellCompDirectiveDefault
	},
	Run: func(cmd *cobra.Command, args []string) {

		fmt.Println("你选择的redis模式是：", redis.Mode)
		fmt.Println("将要连接到redis的地址是：", redis.Address)
		// sr := tb.NewRedis()
		// tbredis.Add(sr)
	},
}

type opt struct {
	Mode     string
	Address  []string
	UserName string
	Password string
}

var redis opt

func init() {
	rootCmd.AddCommand(redisCmd)

	redisCmd.Flags().StringVarP(&redis.Mode, "mode", "m", "standalone", "指定redis类型")
	redisCmd.Flags().StringArrayVar(&redis.Address, "address", []string{}, "指定iP列表")
	redisCmd.Flags().StringVarP(&redis.UserName, "username", "u", "", "指定用户名")
	redisCmd.Flags().StringVarP(&redis.Password, "password", "p", "", "指定密码")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// redisCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// redisCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
