package fiter

import (
	"openai/internal/config"
	"regexp"
)

var reg_map_list map[*regexp.Regexp]string

func init() {
	reg_map_list = compileReg(config.Wechat.Keyword)
}

func Check(text string) string {
	for k, v := range reg_map_list {
		if k.MatchString(text) {
			return v
		}
	}
	return ""
}

func compileReg(m map[string]string) map[*regexp.Regexp]string {
	compiled := make(map[*regexp.Regexp]string)
	for k, v := range m {
		r := regexp.MustCompile(k)
		compiled[r] = v
	}
	return compiled
}
