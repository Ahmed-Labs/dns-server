package main

import (
	"encoding/json"
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
			continue
		}

		query := NewMessage()
		receivedData := buf[:size]

		err = query.deserialize(receivedData)
		if err != nil {
			fmt.Println("Failed to deserialize dns message:", err)
			continue
		}
		fmt.Println("QUERY")
		PrintMessage(query)

		response := NewResponse(query)

		fmt.Println("RESPONSE")
		PrintMessage(response)
		_, err = udpConn.WriteToUDP(response.serialize(), source)
		if err != nil {
			fmt.Println("Failed to send response:", err)
		}
	}
}

func PrintMessage(msg *Message) {
	data, err := json.MarshalIndent(msg, "", "  ")
	if err != nil {
		fmt.Println("Error printing message:", err)
		return
	}
	fmt.Println(string(data))
}