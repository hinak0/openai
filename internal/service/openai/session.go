package openai

import (
	"context"
	"encoding/json"
	"openai/internal/config"
	"openai/internal/util"

	"github.com/go-redis/redis/v8"
)

var (
	client *redis.Client
	ctx    context.Context
)

func init() {
	if config.Session.Type != "redis" {
		util.Logger.Fatal("目前只支持redis存储!")
	}
	ctx = context.Background()
	client = redis.NewClient(&redis.Options{
		Addr:     config.Session.Addr,
		Password: config.Session.Password,
		DB:       config.Session.Database,
	})

	_, err := client.Ping(ctx).Result()
	if err != nil {
		util.Logger.Println("数据库错误 : ", err)
	}
}

func getHistory(uid string) (history []requestMessageItem, err error) {
	val, e := client.Get(ctx, uid).Result()
	if e != nil {
		return nil, e
	}
	err = json.Unmarshal([]byte(val), &history)
	if config.Debug {
		util.Logger.Printf("获取 %s 聊天记录 %s: ", uid, history)
	}
	return
}

func setHistory(uid string, history []requestMessageItem) error {
	if config.Debug {
		util.Logger.Printf("存储 %s 聊天记录 %s: ", uid, history)
	}
	j, _ := json.Marshal(&history)
	err := client.Set(ctx, uid, j, 0).Err()
	return err
}
