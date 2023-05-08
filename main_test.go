package main

import (
	"fmt"
	"openai/internal/config"
	"openai/internal/service/openai"
	"testing"
	"time"
)

func Test(t *testing.T) {
	var in string
	for {
		if _, err := fmt.Scanf("%s", &in); err != nil {
		} else {
			r := openai.Query("001", in, time.Second*time.Duration(config.Wechat.Timeout))
			fmt.Println(r)
		}
	}
}
