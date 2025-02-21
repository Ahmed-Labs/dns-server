package main

func NewResponse(query *Message) *Message {
	var rcode uint8
	if query.Header.OPCode != 0 {
		rcode = 4 // not implemented
	}
	msg :=  &Message{
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
			QDCount:            uint16(len(query.Questions)),
			ANCount:            0,
			NSCount:            0,
			ARCount:            0,
		},
	}

	for _, question := range query.Questions {
		msg.Questions = append(msg.Questions, question)
		msg.Answers = append(msg.Answers, Answer{
			Name:   question.Name,
			Type:   QTYPE_A,
			Class:  QCLASS_IN,
			TTL:    60,
			Length: 4,
			Data:   []byte{8, 8, 8, 8},
		})
		msg.Header.ANCount++
	}

	return msg
}
