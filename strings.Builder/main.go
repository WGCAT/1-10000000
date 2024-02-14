package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {

	start := time.Now()
	var builder strings.Builder

	for i := 0; i <= 10000000; i++ {
		builder.WriteString(" ")
		builder.WriteString(fmt.Sprintf("%d", i))
	}

	s := builder.String()
	fmt.Println(s)

	fmt.Println("소요시간 : ", time.Since(start))
}
