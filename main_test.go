package main

import (
	"fmt"
	"openai/internal/service/fiter"
	"testing"
)

func Test(t *testing.T) {
	res := fiter.Check("keyword")
	if len(res) != 0 {
		fmt.Println(res)
		return
	}
	fmt.Println("no keyword match")
}
