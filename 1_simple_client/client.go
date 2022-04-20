package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"syscall"
	"time"
)

var (
	ip          = flag.String("ip", "192.168.1.104", "server IP")
	connections = flag.Int("conn", 30000, "number of tcp connections")
	startMetric = flag.String("sm", time.Now().Format("2006-01-02T15:04:05 -0700"), "start time point of all clients")
)

func main() {
	fmt.Println("lsfjsldf")
	fmt.Println("lsfjsldf")
	fmt.Println("lsfjsldf")
	fmt.Println("lsfjsldf")
	fmt.Println("lsfjsldf")
	flag.Parse()

	setLimit()

	addr := *ip + ":8972"
	log.Printf("连接到 %s", addr)
	var conns []net.Conn
	for i := 0; i < *connections; i++ {
		c, err := net.DialTimeout("tcp", addr, 10*time.Second)
		if err != nil {
			fmt.Println("failed to connect", i, err)
			i--
			continue
		}
		conns = append(conns, c)
		time.Sleep(time.Millisecond)
	}

	defer func() {
		log.Printf("defer %d 连接", len(conns))
		for _, c := range conns {
			c.Close()
		}
	}()

	log.Printf("完成初始化 %d 连接", len(conns))

	tts := 20 * time.Second
	// if *connections > 100 {
	// 	tts = time.Millisecond * 5
	// }

	for {
		for i := 0; i < len(conns); i++ {
			conn := conns[i]
			// log.Printf("连接 %d 发送数据", i)
			conn.Write([]byte("hello world\r\n"))
		}
		log.Printf("sleep %d 连接", len(conns))
		time.Sleep(tts)
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
	fmt.Println(rLimit)

}
