package main

import (
	"fmt"
	"net"
)

func main() {
	fmt.Println("Listen tick server at 224.0.0.1:9999")
	address, err := net.ResolveUDPAddr("udp", "224.0.0.1:9999") // マルチキャスト用のアドレスを作成
	if err != nil {
		panic(err)
	}
	listener, err := net.ListenMulticastUDP("udp", nil, address) // マルチキャスト用のUDPソケットを作成
	defer listener.Close()

	buffer := make([]byte, 1500)

	for {
		length, remoteAddress, err := listener.ReadFromUDP(buffer) // マルチキャストで送信されてきたデータを読み込む
		if err != nil {
			panic(err)
		}
		fmt.Printf("Server %v\n", remoteAddress)
		fmt.Printf("Now %s\n", string(buffer[:length]))
	}
}
