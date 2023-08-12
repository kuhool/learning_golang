package crossingLevels

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"log"
)

func ProtocolBufferTest() {
	person := &Person{
		Name: "John",
		Age:  25,
	}

	// 编码消息
	data, err := proto.Marshal(person)
	if err != nil {
		log.Fatalf("Failed to encode person: %v", err)
	}

	fmt.Println("Encoded data:", data)

	// 解码消息
	decodedPerson := &Person{}
	err = proto.Unmarshal(data, decodedPerson)
	if err != nil {
		log.Fatalf("Failed to decode person: %v", err)
	}

	fmt.Println("Decoded person:", decodedPerson)
	fmt.Println("Name:", decodedPerson.Name)
	fmt.Println("Age:", decodedPerson.Age)
}
