package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	addr, err := net.ResolveUDPAddr("udp", "0.0.0.0:10001")
	if err != nil {
		fmt.Println("failed to resolve address", err)
		os.Exit(1)
	}

	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		fmt.Println("failed to listen:", err)
		os.Exit(1)
	}
	defer conn.Close()
	fmt.Println("listening on", addr)

	buf := make([]byte, 1024)
	for {
		n, a, err := conn.ReadFromUDP(buf)
		if err != nil {
			fmt.Println("Failed to read: ", err)
			continue
		}
		fmt.Printf("Received %s from %s\n", string(buf[0:n]), a)
	}
}
