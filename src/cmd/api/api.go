/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

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
package api

import (
	"context"
	"fmt"
	"gin-vue-admin/app/router"
	"gin-vue-admin/tool"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/spf13/cobra"
)

// APICmd command
var (
	environment string
	APICmd      = &cobra.Command{
		Use:     "api",
		Short:   "start api server",
		Example: "gin-vue-admin api -c config/config.yaml",
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println("start API server!")
			startHTTP()
			return nil
		},
	}
)

func init() {
	// APICmd.AddCommand(APICmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// api/apiCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// api/apiCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	APICmd.Flags().StringVarP(&environment, "env", "e", "prod", "default local, program running environment dev|local|prod")
}

//startHttp 启动http服务
func startHTTP() error {
	var host = tool.GetConfig("app.host")

	router := router.Initrouter()
	srv := &http.Server{
		Addr:    host,
		Handler: router,
	}

	go func() {
		// 服务连接,暂时只支持http
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal("listen: ", err)
		}
	}()

	fmt.Println("server run start.")
	// 等待中断信号以优雅地关闭服务器（设置 5 秒的超时时间）
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	fmt.Printf("%s Shutdown Server ... \r\n", tool.GetCurrentTimeStr())

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	fmt.Println("Server exiting")

	return nil
}
