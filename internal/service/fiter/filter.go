package fiter

import (
	"openai/internal/config"
)

func Check(text string) string {
	for k, v := range config.Wechat.Keyword_Reg_List {
		if k.MatchString(text) {
			return v
		}
	}
	return ""
}
