package handlers

import (
	"fmt"
	"net/http"

	"github.com/natasharw/lsmtree-kvstore/storage"
)

func DeleteKey(buffer *storage.Memtable) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "delete key functionality is not yet supported. supported endpoints: /set?yourkey=yourvalue and /get?key=yourkey")
		fmt.Println("delete key endpoint hit")
		defer req.Body.Close()
		w.WriteHeader(http.StatusNotFound)
	})
}
