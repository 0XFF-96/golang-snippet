package eof

import (
"bufio"
"fmt"
"net"
"testing"
)
// 1.Name Resolution
// 2.
func TestDial(t *testing.T){
	l, err := net.Listen("new", "127:")
	if err != nil {
		t.Error(err)
	}
	t.Log(l)
}


func TestConectToServer(t *testing.T){
	conn, err := net.Dial("tcp", "golang.org:80")
	if err != nil {
		// handle error
	}
	fmt.Fprintf(conn, "GET / HTTP/1.0\r\n\r\n")
	status, err := bufio.NewReader(conn).ReadString('\n')
	t.Log(status)
}

func TestServer(t *testing.T){
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		// handle error
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			// handle error
		}
		t.Log(conn)
		// go handleConnection(conn)
	}
}


// 在 Listen 有一个策略模式的实现

// first returns the first address which satisfies strategy, or if
// none do, then the first address of any kind.
/*
func (addrs addrList) first(strategy func(Addr) bool) Addr {
	for _, addr := range addrs {
		if strategy(addr) {
			return addr
		}
	}
	return addrs[0]
}
*/
