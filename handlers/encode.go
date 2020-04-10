package handlers

import (
	"bytes"
	"encoding/gob"
)

//[TODO] - make DRY

func StrToBytes(str string) []byte {
	buff := &bytes.Buffer{}
	gob.NewEncoder(buff).Encode(str)
	bs := buff.Bytes()
	return bs
}

func StrsToBytes(strs []string) []byte {
	buff := &bytes.Buffer{}
	gob.NewEncoder(buff).Encode(strs)
	bs := buff.Bytes()
	return bs
}
