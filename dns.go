package main

import (
	"fmt"
	"net"
)

func DNS(host string) string {
	ips, err := net.LookupIP(host)
	if err != nil {
		fmt.Println("Error: ", err)
		return ">Error: Invalid Host"
	}
	out := ""
	for _, ip := range ips {
		out += fmt.Sprintf(">"+host+" IN A %s\n", ip.String())
	}
	return out
}
