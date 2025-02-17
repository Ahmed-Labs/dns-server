package main

import (
	"encoding/binary"
	"strings"
)

type Message struct {
	Header Header
	Question Question
}

func (m Message) encode() []byte {
	data := []byte{}
	header := m.Header.encode()
	data = append(data, header[:]...)
	data = append(data, m.Question.encode()...)
	return data
}

type Header struct {
	ID uint16
	QR bool
	OPCode uint8
	Authoitative bool
	Truncation bool
	RecursionDesired bool
	RecursionAvailable bool
	Reserved uint8
	RCODE uint8
	QDCount uint16
	ANCount uint16
	NSCount uint16
	ARCount uint16
}

func (h Header) encode() [12]byte {
	data := [12]byte{}
	offset := 0

	binary.BigEndian.PutUint16(data[offset:], h.ID)
	offset += 2

	if h.QR {
		data[offset] |= 1 << 7
	}
	data[offset] |= (h.OPCode & 0b1111) << 3

	if h.Authoitative {
		data[offset] |= 1 << 2
	}
	if h.Truncation {
		data[offset] |= 1 << 1
	}
	if h.RecursionDesired {
		data[offset] |= 1
	}

	offset += 1
	if h.RecursionAvailable {
		data[offset] |= 1 << 7
	}
	data[offset] |= (h.Reserved & 0b111) << 4

	data[offset] |= h.RCODE & 0b1111
	offset += 1

	binary.BigEndian.PutUint16(data[offset:], h.QDCount)
	offset += 2

	binary.BigEndian.PutUint16(data[offset:], h.ANCount)
	offset += 2

	binary.BigEndian.PutUint16(data[offset:], h.NSCount)
	offset += 2

	binary.BigEndian.PutUint16(data[offset:], h.ARCount)

	return data
}

type Question struct {
	Name string
	Type QuestionType
	Class QuestionClass
}

func (q Question) encode() []byte {
	data := encodeDomainName(q.Name)
	data = binary.BigEndian.AppendUint16(data, uint16(q.Type))
	data = binary.BigEndian.AppendUint16(data, uint16(q.Class))
	return data
}

func encodeDomainName(name string) []byte {
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