package main

import (
	"fmt"
	"strings"
)

type T struct {
	name  string // オブジェクトの名前
	value int    // その値
}

func main() {
	sum := 0
	for i := 1; i <= 10; i++ {
		if i > 5 {
			s := "iが6以上です"
			fmt.Printf("%d\n", s)
			fmt.Printf("%v\n", strings.Fields(s))
		}
		sum += i
	}
	fmt.Printf("%d\n", sum)
}
