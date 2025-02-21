package main

import "strings"

// creates a label sequence from a domain name
func DomainToLabels(name string) []byte {
	if name == "" {
		return []byte{0x00}
	}
	labels := strings.Split(name, ".")
	totalSize := len(name) + len(labels)
	data := make([]byte, totalSize)

	idx := 0
	for _, label := range labels {
		data[idx] = byte(len(label))
		idx++
		copy(data[idx:], label)
		idx += len(label)
	}

	data[idx] = 0x00
	return data
}