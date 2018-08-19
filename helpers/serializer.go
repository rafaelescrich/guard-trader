package helpers

import (
	"bytes"
	"fmt"
	"log"

	ent "github.com/guard-trader/entities"
	"github.com/ugorji/go/codec"
)

var (
	ch codec.CborHandle
)

func SerializerString(b string) []byte {
	var result bytes.Buffer
	enc := codec.NewEncoder(&result, &ch)
	if err := enc.Encode(b); err != nil {
		fmt.Printf("error on serializer: %s\n", err)
	}
	return result.Bytes()
}

// SerializerAccount serialize acc
func SerializerAccount(account ent.Account) []byte {
	var result bytes.Buffer
	enc := codec.NewEncoder(&result, &ch)
	if err := enc.Encode(account); err != nil {
		log.Printf("error on serialize acc: %s\n", err)
	}
	return result.Bytes()
}
