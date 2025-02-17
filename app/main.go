package main

import (
	"fmt"
	"net"
)

func main() {
	udpAddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:2053")
	if err != nil {
		fmt.Println("Failed to resolve UDP address:", err)
		return
	}
	
	udpConn, err := net.ListenUDP("udp", udpAddr)
	if err != nil {
		fmt.Println("Failed to bind to address:", err)
		return
	}
	defer udpConn.Close()
	
	buf := make([]byte, 512)
	
	for {
		size, source, err := udpConn.ReadFromUDP(buf)
		if err != nil {
			fmt.Println("Error receiving data:", err)
			break
		}
	
		receivedData := string(buf[:size])
		fmt.Printf("Received %d bytes from %s: %s\n", size, source, receivedData)
	
		// Create an empty response
		response := Message{
			Question: Question{
				Name: "codecrafters.io",
				Type: QTYPE_A,
				Class: QCLASS_IN,
			},
			Header: Header{
				ID: 1234,
				QR: true,
				OPCode: 0,
				Authoitative: false,
				Truncation: false,
				RecursionDesired: false,
				RecursionAvailable: false,
				Reserved: 0,
				RCODE: 0,
				QDCount: 1,
				ANCount: 0,
				NSCount: 0,
				ARCount: 0,
			},
		}.encode()
	
		_, err = udpConn.WriteToUDP(response, source)
		if err != nil {
			fmt.Println("Failed to send response:", err)
		}
	}
}
