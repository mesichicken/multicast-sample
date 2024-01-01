package main

import (
	"fmt"
	"net"
	"time"
)

const interval = 10 * time.Second

func main() {
	fmt.Println("Start tick server at 224.0.0.1:9999")
	conn, err := net.Dial("udp", "224.0.0.1:9999")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	start := time.Now()
	wait := start.Truncate(interval).Add(interval).Sub(start) // 10秒ごとに送信するための待ち時間を計算
	time.Sleep(wait)
	ticker := time.Tick(interval) // 10秒ごとにチャネルに値が送信される
	for now := range ticker {
		conn.Write([]byte(now.String()))
		fmt.Println("Tick: ", now.String())
	}
}
