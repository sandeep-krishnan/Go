package main

import (
	"fmt"
	"net"
	"strings"
	"time"
)

func main() {
	listener, error := net.Listen("tcp", ":80")
	if error != nil {
		fmt.Println("Error ", error)
		return
	}

	defer listener.Close()

	for {
		conn, error := listener.Accept()
		if error != nil {
			fmt.Println("Error ", error)
			conn.Close()
			continue
		}

		go handleRequest(conn)
	}
}

func handleRequest(conn net.Conn) {
	defer conn.Close()
	for {
		conn.SetReadDeadline(time.Now().Add(time.Second * 10))
		data := make([]byte, 1024)
		n, error := conn.Read(data)

		if n == 0 || error != nil {
			fmt.Println("Error ", error)
			return
		}

		input := stripCtlAndExtFromBytes(string(data))
		fmt.Println(input)

		if input == "quit" {
			return
		}

		conn.Write([]byte(strings.ToUpper(input) + "\n"))

		conn.SetWriteDeadline(time.Now().Add(time.Second * 5))
	}
}

func stripCtlAndExtFromBytes(str string) string {
	b := make([]byte, len(str))
	var bl int
	for i := 0; i < len(str); i++ {
		c := str[i]
		if c >= 32 && c < 127 {
			b[bl] = c
			bl++
		}
	}
	return string(b[:bl])
}
