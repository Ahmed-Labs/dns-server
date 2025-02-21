package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

type Message struct {
	Header   *Header
	Questions []Question
	Answers   []Answer
}

func NewMessage() *Message {
	return &Message{
		Header: &Header{},
	}
}

func (m *Message) WithHeader(h Header) *Message {
	m.Header = &h
	return m
}

func (m Message) serialize() []byte {
	var header [12]byte
	var question, answer []byte
	size := 0

	if m.Header != nil {
		header = m.Header.serialize()
		size += 12
	}
	if m.Questions != nil {
		for _, q := range m.Questions {
			question = append(question, q.serialize()...)
		}
		size += len(question)
	}
	if m.Answers != nil {
		for _, a := range m.Answers {
			answer = append(answer, a.serialize()...)
		}
		size += len(answer)
	}

	data := make([]byte, 0, size)
	data = append(data, header[:]...)
	data = append(data, question...)
	data = append(data, answer...)
	
	return data
}

func (m *Message) deserialize(b []byte) error {
	reader := bytes.NewReader(b)

	err := m.Header.deserialize(reader)
	if err != nil {
		return err
	}
	return nil
}

type Header struct {
	ID                 uint16
	QR                 bool
	OPCode             uint8
	Authoritative      bool
	Truncation         bool
	RecursionDesired   bool
	RecursionAvailable bool
	Reserved           uint8
	RCODE              uint8
	QDCount            uint16
	ANCount            uint16
	NSCount            uint16
	ARCount            uint16
}

func (h Header) serialize() [12]byte {
	data := [12]byte{}
	offset := 0

	binary.BigEndian.PutUint16(data[offset:], h.ID)
	offset += 2

	if h.QR {
		data[offset] |= 1 << 7
	}
	data[offset] |= (h.OPCode & 0b1111) << 3

	if h.Authoritative {
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

func (h *Header) deserialize(reader *bytes.Reader) error {
	err := binary.Read(reader, binary.BigEndian, &h.ID)
	if err != nil {
		return err
	}

	fmt.Println(reader)
	fmt.Println("ID", h.ID)

	var b byte
	err = binary.Read(reader, binary.BigEndian, &b)
	if err != nil {
		return err
	}
	h.QR = (b & (1 << 7)) != 0
	h.OPCode = (b >> 3) & 0b1111
	h.Authoritative = (b & (1 << 2)) != 0
	h.Truncation = (b & (1 << 1)) != 0
	h.RecursionDesired = (b & 1) != 0

	err = binary.Read(reader, binary.BigEndian, &b)
	if err != nil {
		return err
	}
	h.RecursionAvailable = (b & (1 << 7)) != 0
	h.Reserved = (b >> 4) & 0b111
	h.RCODE = b & 0b1111

	err = binary.Read(reader, binary.BigEndian, &h.QDCount)
	if err != nil {
		return err
	}
	err = binary.Read(reader, binary.BigEndian, &h.ANCount)
	if err != nil {
		return err
	}
	err = binary.Read(reader, binary.BigEndian, &h.NSCount)
	if err != nil {
		return err
	}
	err = binary.Read(reader, binary.BigEndian, &h.ARCount)
	if err != nil {
		return err
	}

	return nil
}

type Question struct {
	Name  string
	Type  QuestionType
	Class QuestionClass
}

func (q Question) serialize() []byte {
	data := DomainToLabels(q.Name)
	data = binary.BigEndian.AppendUint16(data, uint16(q.Type))
	data = binary.BigEndian.AppendUint16(data, uint16(q.Class))
	return data
}
type Answer struct {
	Name   string
	Type   QuestionType
	Class  QuestionClass
	TTL    uint32
	Length uint16
	Data   []byte
}

func (a Answer) serialize() []byte {
	data := DomainToLabels(a.Name)
	data = binary.BigEndian.AppendUint16(data, uint16(a.Type))
	data = binary.BigEndian.AppendUint16(data, uint16(a.Class))
	data = binary.BigEndian.AppendUint32(data, a.TTL)
	data = binary.BigEndian.AppendUint16(data, a.Length)

	if a.Data != nil {
		data = append(data, a.Data...)
	}
	return data
}
