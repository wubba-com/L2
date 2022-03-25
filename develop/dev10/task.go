package main

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/wubba-com/L2/develop/dev10/config"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const (
	_   = iota
	KB  = 1 << (10 * 1) // size read buffer
	req = "GET / HTTP/1.1 Host: seasonkrasoty.ru Keep-Alive: 300 Connection: keep-alive User-Agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/99.0.4844.82 Safari/537.36"
)

func main() {
	// init config
	cfg := config.Get()

	// connect
	conn, err := net.Dial("tcp", net.JoinHostPort(cfg.Host, cfg.Port))
	if err != nil {
		fmt.Println("[err] conn: ", err.Error())
		return
	}

	// close connect
	defer conn.Close()

	// init signals ctl+D, ctl+C
	gFulShotDown := make(chan os.Signal)
	signal.Notify(gFulShotDown, syscall.SIGQUIT, syscall.SIGINT)

	// connections server
	go Client(conn, cfg.Timeout)

	// listen close event
	select {
	case <-gFulShotDown:
		fmt.Printf("\nExit\n")
	}
}

func Client(conn net.Conn, timeout time.Duration) {

	for {
		var err error
		fmt.Print("msg: ")

		var input []byte
		scanner := bufio.NewScanner(os.Stdin)
		if scanner.Scan() {
			input = scanner.Bytes()
		}

		//  timeout на чтение и запись (если через T сек не будет ответа, запрос завершиться с ошибкой timeout-a)
		err = conn.SetDeadline(time.Now().Add(timeout))
		if err != nil {
			fmt.Println("[err write deadline]", err.Error())
			return
		}

		_, err = conn.Write(input)
		if err != nil {
			fmt.Println("[err write]", err.Error())
			continue
		}

		buff := make([]byte, KB*4)
		n, err := conn.Read(buff)
		if err != nil {
			fmt.Println("[err read]", errors.New("long request"))
			continue
		}

		fmt.Fprintf(os.Stdout, "server: %s\n", string(buff[:n]))
	}
}
