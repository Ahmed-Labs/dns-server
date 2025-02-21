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

func LabelsToDomain(labels []byte) string {
	var sb strings.Builder
	idx := 0

	for idx < len(labels) && labels[idx] != 0x00 {
		length := int(labels[idx])
		label := string(labels[idx+1 : idx+1+length])

		sb.WriteString(label)

		idx += length + 1
		if idx < len(labels) && labels[idx] != 0x00 {
			sb.WriteByte('.')
		}
	}

	return sb.String()
}
