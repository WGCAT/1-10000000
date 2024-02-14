package main

import (
	"fmt"
	"time"
)

var num int

func for1(num int, c chan int) {
	for i := 0; i < 1000; i++ {
		num++
		fmt.Println(num)
		for j := 0; j < 1000; j++ {
			num++
			fmt.Println(num)

		}

	}
	c <- num
}

// func for2(num int, c chan int) {
// 	for j := 3000000; j < 6000000; j++ {
// 		num++
// 		fmt.Println(j)

// 	}
// 	c <- num

// }

// func for3(num int, c chan int) {
// 	for p := 6000000; p < 10000000; p++ {
// 		num++
// 		fmt.Println(p)

// 	}
// 	c <- num

// }
func main() {
	now := time.Now()

	c := make(chan int)

	num = 0

	for i := 0; i < 10; i++ {
		go for1(num, c)
		fmt.Println(<-c)
	}

	fmt.Println("마지막값 : ", 10000000)
	fmt.Println(time.Since(now))

	// time.Sleep(time.Millisecond * 100000000000)

}
