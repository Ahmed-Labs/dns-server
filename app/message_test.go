package main

// import (
// 	"bytes"
// 	"encoding/json"
// 	"fmt"
// 	"reflect"
// 	"testing"
// )

// func TestSerialization(t *testing.T) {
// 	message := &Message{
// 		Header: &Header{
// 			ID:                 1234,
// 			QR:                 true,
// 			OPCode:             0,
// 			Authoritative:      false,
// 			Truncation:         false,
// 			RecursionDesired:   true,
// 			RecursionAvailable: false,
// 			Reserved:           0,
// 			RCODE:              0,
// 			QDCount:            1,
// 			ANCount:            1,
// 			NSCount:            0,
// 			ARCount:            0,
// 		},
// 		Questions: []Question{{
// 			Name:  "codecrafters.io",
// 			Type:  QTYPE_A,
// 			Class: QCLASS_IN,
// 		}},
// 		Answers: []Answer{{
// 			Name:   "codecrafters.io",
// 			Type:   QTYPE_A,
// 			Class:  QCLASS_IN,
// 			TTL:    60,
// 			Length: 4,
// 			Data:   []byte{8, 8, 8, 8},
// 		}},
// 	}

// 	tests := []struct {
// 		name string
// 		obj interface {
// 			serialize() []byte
// 			deserialize(*bytes.Reader) error
// 		}
// 	}{	
// 		{
// 			name: "serialize/deserialize header",
// 			obj: message.Header,
// 		},
// 	}

// 	// for _, test := range tests {
// 	// 	serialized := test.obj.serialize()
		
// 	// 	deserialized

// 	// 	if !reflect.DeepEqual(message, deserializedMsg) {
// 	// 		t.Error("serialization and deserialization are not symmetric.")
// 	// 	}
// 	// }

// }

// func PrintMessage(msg *Message) {
// 	data, err := json.MarshalIndent(msg, "", "  ")
// 	if err != nil {
// 		fmt.Println("Error printing message:", err)
// 		return
// 	}
// 	fmt.Println(string(data))
// }