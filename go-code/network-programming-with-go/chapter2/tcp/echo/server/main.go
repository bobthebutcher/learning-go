package main

import (
	"flag"
	"fmt"
	"net"
	"os"
)

func main() {
	var addr string
	flag.StringVar(&addr, "e", ":4040", "service address endpoint")
	flag.Parse()

	laddr, err := net.ResolveTCPAddr("tcp", addr)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	l, err := net.ListenTCP("tcp", laddr)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer l.Close()
	fmt.Println("listenting at (tcp)", laddr.String())

	for {
		conn, err := l.AcceptTCP()
		if err != nil {
			fmt.Println("failed to accept conn:", err)
			conn.Close()
			continue
		}
		fmt.Println("connected to:", conn.RemoteAddr())

		go handleConnection(conn)
	}
}

func handleConnection(conn *net.TCPConn) {
	defer conn.Close()

	buf := make([]byte, 1024)

	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println(err)
		return
	}

	w, err := conn.Write(buf[:n])
	if err != nil {
		fmt.Println("failed to write to client: ", err)
		return
	}
	if w != n {
		fmt.Println("warning: not all data sent to client")
		return
	}
}
