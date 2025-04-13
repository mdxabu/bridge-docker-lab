package main

import (
	"fmt"
	"net"
)

func main() {
	addr := net.UDPAddr{
		Port: 9090,
		IP:   net.ParseIP("0.0.0.0"),
	}
	conn, err := net.ListenUDP("udp", &addr)
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	buf := make([]byte, 1024)
	for {
		n, remote, _ := conn.ReadFromUDP(buf)
		fmt.Printf("[server] Received: %s from %s\n", string(buf[:n]), remote)
	}
}
