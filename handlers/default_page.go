package handlers

import (
	"fmt"
	"net/http"

	"github.com/natasharw/lsmtree-kvstore/storage"
)

// DefaultPage : prints a message to browser directing users to supported endpoints
func DefaultPage(buffer *storage.Memtable) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "welcome to the key-value store homepage. supported endpoints: /set?yourkey=yourvalue and /get?key=yourkey")
		defer req.Body.Close()
		w.WriteHeader(http.StatusOK)
	})
}
