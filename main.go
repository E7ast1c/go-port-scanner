package main

import (
	"fmt"
	"net"
	"port-scanner/internal/flags"
	"sync"
	"time"
)

const (
	timeout time.Duration = time.Millisecond * 50
)

type scan struct {
	Addr     string
	Ports    []string
	ConnType string

	mx *sync.Mutex
}

func main() {
	addr, ports, connType := flags.ParseArgs()
	s := scan{
		Addr:     addr,
		Ports:    ports,
		ConnType: connType,
		mx:       &sync.Mutex{},
	}

	for i := range s.Ports {
		if err := scanPort(s.Addr, s.Ports[i]); err != nil {
			fmt.Printf("connect to %s:%s (%s) failed: %s\n", s.Addr, s.Ports[i], s.ConnType, err)
			return
		}
	}
}

func scanPort(addr, port string) error {
	conn, err := net.DialTimeout("tcp", addr+":"+port, timeout)
	if err != nil {
		return err
	}
	defer conn.Close()

	fmt.Printf("[+] connection to %s:%s [tcp/http] succeeded!\n", addr, port)
	return nil
}
