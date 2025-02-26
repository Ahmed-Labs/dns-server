package main

import (
	"strings"
)

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

func LabelsToDomain(labels []byte, compressionMap map[int]string, offset int) string {
	var sb strings.Builder
	type pieceType struct{
		name string
		netOffset int
	}
	idx := 0
	pieces := []pieceType{}

	for idx < len(labels) && labels[idx] != 0x00 {
		if labels[idx] >> 6 == 0b11 {
			domainPointer := (uint16(labels[idx] & 0x3F) << 8) | uint16(labels[idx+1])
			if domainName, ok := compressionMap[int(domainPointer)]; ok {
				sb.WriteString(domainName)
			}
			idx += 2
		} else {
			length := int(labels[idx])
			label := string(labels[idx+1 : idx+1+length])
			sb.WriteString(label)
			pieces = append(pieces, pieceType{
				name: label,
				netOffset: idx + offset,
			})
			idx += length + 1
		}
		if idx < len(labels) && labels[idx] != 0x00 {
			sb.WriteByte('.')
		}
	}

	suffix := ""
	for i := len(pieces)-1; i >= 0; i-- {
		piece := pieces[i]
		suffix = piece.name + suffix
		compressionMap[piece.netOffset] = suffix

		if i - 1 >= 0 {
			suffix = "." + suffix
		}
	}

	return sb.String()
}
