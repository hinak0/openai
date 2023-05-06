package fiter

import (
	"fmt"
	"openai/internal/config"
)

// var (
// 	instance = sieve.New()
// )

func init() {

	go func() {
		// // 加载预定义词典
		// arr := []
		// var builder strings.Builder
		// builder.Grow(20)

		// for _, v := range arr {
		// 	runes := strings.Split(v, " ")
		// 	for _, w := range runes {
		// 		i, _ := strconv.Atoi(w)
		// 		builder.WriteRune(rune(i))
		// 	}
		// 	arr = append(arr, builder.String())
		// 	builder.Reset()
		// }

		// instance.Add(arr)

		// // 加载你定义的词典
		// f, err := os.Open("./keyword.txt")
		// if err != nil {
		// 	return
		// }
		// arr = arr[:0]
		// br := bufio.NewReader(f)
		// for {
		// 	a, _, c := br.ReadLine()
		// 	if c == io.EOF {
		// 		break
		// 	}
		// 	arr = append(arr, string(a))
		// }

		// instance.Add(arr)
		fmt.Println("no fucking Sensitive words")
	}()

}

func Check(text string) string {
	for k, v := range config.Wechat.Keyword_reg {
		if k.MatchString(text) {
			return v
		}
	}
	return ""
}
