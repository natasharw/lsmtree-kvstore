package handlers

import (
	"bytes"
	"encoding/gob"
	"log"
)

//[TODO] - make DRY

// StrToBytes : Takes a string and encodes to bytes
func StrToBytes(str interface{}) []byte {
	log.Printf("Encoding string %s to bytes", str)
	buff := &bytes.Buffer{}
	gob.NewEncoder(buff).Encode(str)
	bs := buff.Bytes()
	log.Printf("Successfully encoded. Result: %v", bs)
	return bs
}

// StrsToBytes : Takes a slice of strings and encodes to bytes
func StrsToBytes(strs []string) []byte {
	log.Printf("Encoding slice of strings %s to bytes", strs)
	buff := &bytes.Buffer{}
	gob.NewEncoder(buff).Encode(strs)
	bs := buff.Bytes()
	log.Printf("Successfully encoded. Result: %v", bs)
	return bs
}
