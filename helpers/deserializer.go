package helpers

import (
	"fmt"
	"log"

	ent "github.com/guard-trader/entities"
	"github.com/ugorji/go/codec"
)

// DeserializerAccount deselialize acc
func DeserializerAccount(b []byte) ent.Account {
	var acc ent.Account
	dec := codec.NewDecoderBytes(b, &ch)
	if err := dec.Decode(&acc); err != nil {
		log.Printf("error on deserialize acc : %s\n", err)
	}
	return acc
}

// DeserializerString deselialize a string
func DeserializerString(d []byte) string {
	var str string
	dec := codec.NewDecoderBytes(d, &ch)
	if err := dec.Decode(&str); err != nil {
		fmt.Printf("error on deserializer string: %s\n", err)
	}
	return str
}
