package config

import (
	"openai/internal/util"
	"os"

	"github.com/spf13/viper"
)

var (
	Debug bool

	Http struct {
		Addr  string
		Port  string
		Proxy string
	}

	Session struct {
		Enable   bool
		Type     string
		Addr     string
		Password string
		Database int
		Track    int
	}

	OpenAI struct {
		Key string

		Params struct {
			Api         string
			Model       string
			Prompt      string
			Temperature float32
			MaxTokens   uint16
		}

		MaxQuestionLength int
	}

	Wechat struct {
		Token        string
		Timeout      int
		SubscribeMsg string
		Keyword      map[string]string
	}
)

func init() {

	// 读取配置
	viper.SetConfigFile("./config.yaml")
	viper.SetConfigType("yaml")

	err := viper.ReadInConfig()
	if err != nil {
		util.Logger.Println("解析配置文件config.yaml失败:", err.Error())
		os.Exit(0)
	}

	viper.UnmarshalKey("debug", &Debug)
	viper.UnmarshalKey("http", &Http)
	viper.UnmarshalKey("session", &Session)
	viper.UnmarshalKey("openai", &OpenAI)
	viper.UnmarshalKey("wechat", &Wechat)

	if OpenAI.Key == "" {
		util.Logger.Println("OpenAI的Key不能为空")
		os.Exit(0)
	}

	if Wechat.Token == "" {
		util.Logger.Println("未设置公众号token，公众号功能不可用")
	}

	if Wechat.Timeout < 3 || Wechat.Timeout > 13 {
		Wechat.Timeout = 8
	}
}
