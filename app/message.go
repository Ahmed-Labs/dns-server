package main

import "encoding/binary"


type Message struct {
	Header Header

}

func (m Message) encode() []byte {
	header := m.Header.encode()
	return header[:]
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
		data[offset] = 1
	}
	offset += 1

	data[offset] = h.OPCode
	offset += 1

	if h.Authoitative {
		data[offset] = 1
	}
	offset += 1

	if h.Truncation {
		data[offset] = 1
	}
	offset += 1

	if h.RecursionDesired {
		data[offset] = 1
	}
	offset += 1

	if h.RecursionAvailable {
		data[offset] = 1
	}
	offset += 1

	data[offset] = h.Reserved
	offset += 1

	data[offset] = h.RCODE
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