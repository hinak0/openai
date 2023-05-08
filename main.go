package main

import (
	"net/http"
	"openai/bootstrap"
	"openai/internal/config"
	"openai/internal/handler"
	"openai/internal/util"
)

func main() {
	r := bootstrap.New()

	// 微信消息处理
	r.POST("/", handler.ReceiveMsg)
	// 用于公众号自动验证
	r.GET("/", handler.WechatCheck)
	// 用于测试 curl "http://127.0.0.1:$PORT/test"
	r.GET("/test", handler.Test)

	util.Logger.Printf("程序开始监听 %s:%s \n", config.Http.Addr, config.Http.Port)
	if err := http.ListenAndServe(config.Http.Addr+":"+config.Http.Port, r); err != nil {
		util.Logger.Fatalln(err)
	}
}
