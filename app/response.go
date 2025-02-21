package main

func NewResponse(query *Message) *Message {
	var rcode uint8
	if query.Header.OPCode != 0 {
		rcode = 4 // not implemented
	}
	return &Message{
		Header: &Header{
			ID:                 query.Header.ID,
			QR:                 true,
			OPCode:             query.Header.OPCode,
			Authoritative:      false,
			Truncation:         false,
			RecursionDesired:   query.Header.RecursionDesired,
			RecursionAvailable: false,
			Reserved:           0,
			RCODE:              rcode,
			QDCount:            1,
			ANCount:            1,
			NSCount:            0,
			ARCount:            0,
		},
		Questions: []Question{{
			Name:  "codecrafters.io",
			Type:  QTYPE_A,
			Class: QCLASS_IN,
		}},
		Answers: []Answer{{
			Name:   "codecrafters.io",
			Type:   QTYPE_A,
			Class:  QCLASS_IN,
			TTL:    60,
			Length: 4,
			Data:   []byte{8, 8, 8, 8},
		}},
	}
}
