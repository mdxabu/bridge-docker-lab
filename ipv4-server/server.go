package main

import (
    "fmt"
    "net"
    "os"
)

func main() {
    addr := ":8080"
    ln, err := net.Listen("tcp4", addr)
    if err != nil {
        fmt.Println("Listen error:", err)
        os.Exit(1)
    }
    fmt.Println("IPv4 Server listening on", addr)
    for {
        conn, err := ln.Accept()
        if err != nil {
            fmt.Println("Accept error:", err)
            continue
        }
        fmt.Println("Connection from:", conn.RemoteAddr())
        conn.Write([]byte("Hello from IPv4 server!\n"))
        conn.Close()
    }
}
