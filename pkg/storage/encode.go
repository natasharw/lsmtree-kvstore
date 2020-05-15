package storage

import (
	"bytes"
	"encoding/gob"
	"log"
)

// Encode : Takes an undefined input and encodes to bytes
func Encode(input interface{}) []byte {
	log.Printf("Encoding input to bytes. Input: %s", input)
	buff := &bytes.Buffer{}
	gob.NewEncoder(buff).Encode(input)
	bs := buff.Bytes()
	log.Printf("Successfully encoded. Result: %v", bs)
	return bs
}
