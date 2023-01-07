package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	domain := "pilot-agent.agent-pilot.svc.cluster.local"

	_, addrs, err := net.LookupSRV("", "tcp", domain)
	if err != nil {
		log.Fatal(err)
	}

	for _, addr := range addrs {
		fmt.Printf("%s:%d\n", addr.Target, addr.Port)
	}

	ipaddrs, err := net.LookupIP(domain)
	if err != nil {
		log.Fatal(err)
	}

	for _, ipaddr := range ipaddrs {
		fmt.Printf("%s\n", ipaddr)
	}
}
