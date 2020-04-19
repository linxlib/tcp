package main

import (
	"fmt"
	"github.com/linxlib/conv"
	"github.com/linxlib/tcp"
	"time"
)

func main() {
	// Server
	go tcp.NewServer("127.0.0.1:8999", func(conn *tcp.Connection) {
		defer conn.Close()
		for {
			data, err := conn.RecvPkg()
			if err != nil {
				fmt.Println(err)
				break
			}
			fmt.Println("receive:", data)
		}
	}).Run()

	time.Sleep(time.Second)

	// Client
	conn, err := tcp.NewConn("127.0.0.1:8999")
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	for i := 0; i < 10000; i++ {
		if err := conn.SendPkg([]byte(conv.String(i))); err != nil {
			fmt.Println(err)
		}
		time.Sleep(1 * time.Second)
	}

}
