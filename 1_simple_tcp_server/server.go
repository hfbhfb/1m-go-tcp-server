package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	_ "net/http/pprof"
	"sync"
	"syscall"
)

var (
	CurConnCount = 0
	mu           sync.Mutex
)

func main() {
	setLimit()

	ln, err := net.Listen("tcp", ":8972")
	if err != nil {
		panic(err)
	}

	go func() {
		if err := http.ListenAndServe(":6060", nil); err != nil {
			log.Fatalf("pprof failed: %v", err)
		}
	}()

	var connections []net.Conn
	defer func() {
		for _, conn := range connections {
			conn.Close()
		}
	}()

	for {
		conn, e := ln.Accept()
		if e != nil {
			if ne, ok := e.(net.Error); ok && ne.Temporary() {
				log.Printf("accept temp err: %v", ne)
				continue
			}

			log.Printf("accept err: %v", e)
			return
		}

		go handleConn(conn)
		connections = append(connections, conn)
		if len(connections)%1000 == 0 {
			log.Printf("total number of connections: %v", len(connections))
		}
	}
}

func Mark(count int) {
	mu.Lock()
	CurConnCount = CurConnCount + count
	mu.Unlock()
	if CurConnCount%1000 == 0 {
		fmt.Println(CurConnCount)
	}
}

func handleConn(conn net.Conn) {
	Mark(1)
	i, err := io.Copy(ioutil.Discard, conn)
	Mark(-1)
	if err != nil {
		fmt.Println(err)
		fmt.Println(i)
	}
}

func setLimit() {
	var rLimit syscall.Rlimit
	if err := syscall.Getrlimit(syscall.RLIMIT_NOFILE, &rLimit); err != nil {
		panic(err)
	}
	rLimit.Cur = rLimit.Max
	if err := syscall.Setrlimit(syscall.RLIMIT_NOFILE, &rLimit); err != nil {
		panic(err)
	}

	fmt.Printf("set cur limit: %d", rLimit.Cur)
	fmt.Println("setLimit ready")
}
