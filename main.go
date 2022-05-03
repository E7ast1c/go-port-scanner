package main

import (
	"fmt"
	"github.com/alexflint/go-arg"
	"net"
	"port-scanner/pkg/net_validator"
	"sync"
	"time"
)

const (
	timeout    time.Duration = time.Millisecond * 50
	targetAddr               = "151.101.1.69"
	targetPort               = "80"

	netType = "tcp"
)

type scanArgs struct {
	Addr    string `arg:"positional,required" help:"server address"`
	Ports   string `arg:"positional,required" help:"server port(s), example 80, 100-200 or 400,500,600"`
	NetType string `arg:"positional" default:"tcp" help:"net protocol, by default using tcp"`
}

type scan struct {
	URI     string
	NetType string

	mx *sync.Mutex
}

func main() {
	scn := &scanArgs{}
	arg.MustParse(scn)
	fmt.Println("Addr:", scn.Addr)
	fmt.Println("Ports:", scn.Ports)
	fmt.Println("net type:", scn.NetType)

	net_validator.Address(scn.Addr)
	net_validator.PortCollector(scn.Ports)
	net_validator.NetType(scn.NetType)

	s := scan{
		URI:     scn.Addr + ":" + scn.Ports,
		NetType: scn.NetType,
		mx:      &sync.Mutex{},
	}

	if err := s.scanPort(); err != nil {
		s.prettyErr(err)
		return
	}
}

func (s *scan) scanPort() error {
	conn, err := net.DialTimeout(s.NetType, s.URI, timeout)
	defer conn.Close()
	if err != nil {
		return err
	}

	fmt.Printf("[+] connection to %s %s port [tcp/http] succeeded!\n", targetAddr, targetPort)
	return nil
}

func (s *scan) prettyErr(err error) {
	fmt.Printf("connect to %s (%s) failed: %s\n", s.URI, s.NetType, err)
}
