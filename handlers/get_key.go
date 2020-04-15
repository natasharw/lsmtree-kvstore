package handlers

import (
	"fmt"
	"net/http"

	"github.com/natasharw/lsmtree-kvstore/storage"
)

// GetKey : takes a key passed in as a parameter and prints its value to browser if it exists in the store
func GetKey(buffer *storage.Memtable) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		passed := req.URL.Query().Get("key")
		defer req.Body.Close()

		if passed == "" {
			http.Error(w, "invalid GET request. hint: param of key=yourkey must be provided", http.StatusBadRequest)
			return
		}

		key := passed
		// key := StrToBytes(passed)
		value, err := storage.Get(buffer, key)

		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}

		fmt.Println(w, "key-value pair found. %v : $v", key, value) // [TODO] - convert to strings
		w.WriteHeader(http.StatusOK)
	})
}
