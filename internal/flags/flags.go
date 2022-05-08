package flags

import (
	"flag"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

const AppName = "gnc"

var (
	maskNumber = regexp.MustCompile("\\d{1,5}")
	maskDash   = regexp.MustCompile("\\d{1,5}-\\d{1,5}")
	maskComma  = regexp.MustCompile("\\d{1,5}-\\d{1,5}")
	maskAddr   = regexp.MustCompile("^([01]?\\d\\d?|2[0-4]\\d|25[0-5])(?:\\.(?:[01]?\\d\\d?|2[0-4]\\d|25[0-5])){3}(?:/[0-2]\\d|/3[0-2])?$")
)

// These are command line flags we process
var (
	// addr = flag.String("address", "", "target server address")
	// ports = flag.String("ports", "", "server port(s), example 80, 100-200 or 400, 500, 600")
	udp = flag.Bool("u", false, "use UDP instead of TCP")
)

func ParseArgs() (string, []string, string) {
	flag.Parse()

	args := os.Args
	if len(args) < 3 {
		log.Fatalf("required at least two arguments, usage %s 127.0.0.1 80", AppName)
	}

	addr, ports := args[1], args[2]
	if !maskAddr.Match([]byte(addr)) {
		log.Fatalf("incorrect IP address, usage %s 127.0.0.1 80", AppName)
	} else if ports == "" {
		log.Fatalf("incorrect ports, usage %s 127.0.0.1 80", AppName)
	}

	connType := "tcp"
	if *udp {
		connType = "udp"
	}

	return addr, portCollector(ports), connType
}

// portCollector collect and validate ports
func portCollector(ports string) []string {
	// replace whitespace
	ports = strings.ReplaceAll(ports, " ", "")

	pv := func(port string) {
		if numPort, err := strconv.Atoi(port); err != nil {
			log.Fatalf("port parsing failed, invalid port: %s", port)
		} else if numPort < 0 || numPort > 65536 {
			log.Fatalf("port parsing failed, port %s out of range", port)
		}
	}

	switch {
	case maskNumber.Match([]byte(ports)):
		pv(ports)
		return []string{ports}
	case maskDash.Match([]byte(ports)):
		portsList := strings.Split(ports, "-")
		for p := range portsList {
			pv(portsList[p])
		}
		return portsList
	case maskComma.Match([]byte(ports)):
		portsList := strings.Split(ports, "-")
		for p := range portsList {
			pv(portsList[p])
		}
		return portsList
	default:
		log.Fatalf("incorrect ports, usage %s 127.0.0.1 80 or 80-85 or 90,95", AppName)
	}
	return nil
}
