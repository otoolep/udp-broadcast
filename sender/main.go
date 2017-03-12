package main

import (
	"fmt"
	"net"
	"os"
	"strconv"
	"time"
)

func main() {
	sa, err := net.ResolveUDPAddr("udp", "255.255.255.255:10001")
	if err != nil {
		fmt.Println("failed to resolve server address:", err)
		os.Exit(1)
	}

	la, err := net.ResolveUDPAddr("udp", "127.0.0.1:0")
	if err != nil {
		fmt.Println("failed to resolve local address:", err)
		os.Exit(1)
	}

	conn, err := net.DialUDP("udp", la, sa)
	if err != nil {
		fmt.Println("failed to dial:", err)
		os.Exit(1)
	}
	defer conn.Close()

	i := 0
	for {
		msg := strconv.Itoa(i)
		i++
		buf := []byte(msg)
		_, err := conn.Write(buf)
		if err != nil {
			fmt.Println(msg, err)
		}
		time.Sleep(time.Second * 1)
	}
}
