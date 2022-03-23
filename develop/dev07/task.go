package main

import (
	"fmt"
	"reflect"
	"time"
)

func or (channels ...<- chan interface{}) <- chan interface{} {
	var s []reflect.SelectCase
	c := make(chan interface{})

	for _, ch := range channels {
		s = append(s, reflect.SelectCase{
			Dir: reflect.SelectRecv, // Если Dir — SelectRecv, случай представляет операцию получения
			Chan: reflect.ValueOf(ch), // Конкретный канал на чтение
		})
	}

	go func() {
		// Как и в обычном select, здесь будет блокировка до тех пор, пока не будет выполнено хотя бы одно
		// из чтений done-каналов
		reflect.Select(s)
		// После чтения из канала закроется канал, который возобновит работу go-main
		close(c)
	}()

	return c
}


func main()  {


	var sig = func  (after time.Duration) <- chan interface{} {
		c := make(chan interface{})
		go func() {
			defer close(c)
			time.Sleep(after)
		}()
		return c
	}
	start := time.Now()
	<-or (
		sig(2*time.Hour),
		sig(5*time.Minute),
		sig(2*time.Second),
		sig(1*time.Hour),
		sig(1*time.Minute),
	)

	fmt.Printf("fone after %v", time.Since(start))
}
