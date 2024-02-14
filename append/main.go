package main

import (
	"fmt"
	"time"
)

func main() {

	start := time.Now()
	s := make([]int, 0, 10000000)
	for i := 0; i < 10000000; i++ {
		s = append(s, i)
		fmt.Println(s)
	}

	fmt.Println("소요시간 : ", time.Since(start))

}
