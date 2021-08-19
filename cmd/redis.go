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
	"github.com/ybooks240/ToolBox/pkg/tbredis"
)

// redisCmd represents the redis command
var tb tbredis.StandaloneRedis
var redisCmd = &cobra.Command{
	Use:   "redis",
	Short: "will to connect redis",
	Long:  `will to connect redis or test`,
	Run: func(cmd *cobra.Command, args []string) {
		for i := range args {
			fmt.Println(i)
		}
		fmt.Println("redis called")
		// sr := tb.NewRedis()
		// tbredis.Add(sr)
	},
}

func init() {

	rootCmd.AddCommand(redisCmd)

	redisCmd.Flags().StringP("mode", "m", "", "use --mode or -m to choose redis conn mode")
	redisCmd.Flags().StringVar(&tb.IP, "ip", "", "ip address for redis")
	redisCmd.Flags().StringVar(&tb.Port, "port", "", "ip address for redis")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// redisCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// redisCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
