package net_validator

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
)

var (
	maskNumber = regexp.MustCompile("\\d{1,5}")
	maskDash   = regexp.MustCompile("\\d{1,5}-\\d{1,5}")
	maskComma  = regexp.MustCompile("\\d{1,5}-\\d{1,5}")
	maskIPAddr = regexp.MustCompile("^([01]?\\d\\d?|2[0-4]\\d|25[0-5])(?:\\.(?:[01]?\\d\\d?|2[0-4]\\d|25[0-5])){3}(?:/[0-2]\\d|/3[0-2])?$")
)

// PortCollector collect and validate ports
func PortCollector(ports string) []string {
	// replace whitespace
	ports = strings.ReplaceAll(ports, " ", "")

	switch {
	case maskNumber.Match([]byte(ports)):
		if err := portValidator(ports); err != nil {
			log.Fatal(err)
		}
		return []string{ports}
	case maskDash.Match([]byte(ports)):
		portsList := strings.Split(ports, "-")
		for p := range portsList {
			if err := portValidator(portsList[p]); err != nil {
				log.Fatal(err)
			}
		}
		return portsList
	case maskComma.Match([]byte(ports)):
		portsList := strings.Split(ports, "-")
		for p := range portsList {
			if err := portValidator(portsList[p]); err != nil {
				log.Fatal(err)
			}
		}
		return portsList
	default:
		log.Fatalf("the entered port combination does not fit the existing masks: %s", ports)
	}
	return nil
}

func portValidator(port string) error {
	if numPort, err := strconv.Atoi(port); err != nil {
		return fmt.Errorf("port parsing failed, invalid port: %s", port)
	} else if numPort < 0 || numPort > 65536 {
		return fmt.Errorf("port parsing failed, port %s out of range", port)
	}

	return nil
}

func NetType(netType string) {
	if netType == "tcp" || netType == "udp" {
		return
	}
	log.Fatalf("unknown net type %s", netType)
}

func Address(addr string) {
	if maskIPAddr.Match([]byte(addr)) {
		return
	}
	log.Fatalf("incorrect IP address %s", addr)
}
