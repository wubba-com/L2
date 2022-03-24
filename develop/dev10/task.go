package main

import "fmt"

func r(nums []int) chan int{
	out := make(chan int)
	go func(nums []int) {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}(nums)

	return out
}

func r2(c chan int) chan int{
	out := make(chan int)
	go func(out, c chan int) {
		for n := range c {
			out <- n*2
		}
		close(out)
	}(out, c)

	return out
}

func main()  {
	test := []int{1, 1, 1, 1, 1}
	out := r(test)
	out2 := r2(out)
	var sum int
	for n := range out2 {
		sum += n
	}
	fmt.Println(sum)
}
