package main

import (
	"github.com/calee0219/onvmNet"

	"net"
)

func main() {
	//onvmNet.Hi()
	conn, err := onvmNet.ListenUDP("udp", &net.UDPAddr{IP: net.ParseIP("127.0.0.1"), Port: 9981})
	if err != nil {
		panic(err)
	}
	conn.Close()
}
