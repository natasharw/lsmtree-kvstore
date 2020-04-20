package storage

import (
	"bytes"
	"encoding/gob"
	"log"
)

// StrToBytes : Takes a string and encodes to bytes
func StrToBytes(str string) []byte {
	log.Printf("Encoding string %s to bytes", str)
	buff := &bytes.Buffer{}
	gob.NewEncoder(buff).Encode(str)
	bs := buff.Bytes()
	log.Printf("Successfully encoded. Result: %v", bs)
	return bs
}
