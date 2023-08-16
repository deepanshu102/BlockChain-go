package utils

import (
	"fmt"
	"net"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func IsFoundHost(host string, port uint16) bool {
	target := fmt.Sprintf("%s:%d", host, port)
	_, err := net.DialTimeout("tcp", target, 1*time.Second)
	if err != nil {
		fmt.Printf("%s %v \n", target, err)
		return false
	}
	return true
}

// looking for more servers
// but in our case is we are chaning port instead of ip addresses
var PATTERN = regexp.MustCompile(`((25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?\.){3})(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)`)

func FindNeighbors(myHost string, myPort uint16, startIp uint8, endIp uint8, startPort uint16, endPort uint16) []string {
	address := fmt.Sprintf("%s:%d", myHost, myPort)
	m := PATTERN.FindStringSubmatch(myHost)
	if m == nil {
		return nil
	}
	prefixHost := m[1]
	lastIP, _ := strconv.Atoi(m[len(m)-1])
	neighbors := make([]string, 0)
	for port := startPort; port <= endPort; port += 1 {
		for ip := startIp; ip <= endIp; ip += 1 {
			guessHost := fmt.Sprintf("%s%d", prefixHost, lastIP+int(ip))
			guessTarget := fmt.Sprintf("%s:%d", guessHost, port)
			if guessTarget != address && IsFoundHost(guessHost, port) && guessTarget[strings.Index(guessTarget, ":")+1:] != address[strings.Index(address, ":")+1:] {
				neighbors = append(neighbors, guessTarget)
			}
		}
	}
	return neighbors
}

// find own ip address

func GetHost() string {
	hostname, err := os.Hostname()
	if err != nil {
		return "127.0.0.1"
	}
	fmt.Println("hostName:", hostname)
	address, err := net.LookupHost(hostname)
	if err != nil {
		return "127.0.0.1"
	}
	fmt.Println("Address:", address)
	return address[0]
}
