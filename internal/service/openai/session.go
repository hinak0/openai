package openai

import (
	"context"
	"encoding/json"
	"log"
	"openai/internal/config"

	"github.com/go-redis/redis/v8"
)

var (
	client *redis.Client
	ctx    context.Context
)

func init() {
	if config.Session.Type != "redis" {
		panic("目前只支持redis存储!")
	}
	ctx = context.Background()
	client = redis.NewClient(&redis.Options{
		Addr:     config.Session.Addr,
		Password: config.Session.Password,
		DB:       config.Session.Database,
	})

	_, err := client.Ping(ctx).Result()
	if err != nil {
		log.Println("database error : ", err)
	}
}

func getHistory(uid string) (history []requestMessageItem, err error) {
	val, e := client.Get(ctx, uid).Result()
	if e != nil {
		return nil, e
	}
	err = json.Unmarshal([]byte(val), &history)
	log.Println("get history : ", uid, history)
	return
}

func setHistory(uid string, history []requestMessageItem) error {
	log.Println("set history : ", uid, history)
	j, _ := json.Marshal(&history)
	err := client.Set(ctx, uid, j, 0).Err()
	return err
}
