package handlers

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/natasharw/lsmtree-kvstore/storage"
)

func PutKey(buffer *storage.Memtable) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		fmt.Println("SET params were:", req.URL.Query())
		raw := req.URL.RawQuery
		parsed, err := url.ParseQuery(raw)

		defer req.Body.Close()

		if len(parsed) == 0 {
			http.Error(w, "invalid SET request. hint: add a parameter of yourkey=yourvalue", http.StatusBadRequest)
		}

		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		for k, v := range parsed {
			fmt.Fprintln(w, k) //[TODO] placeholder only - remove
			fmt.Fprintln(w, v) // [TODO] placeholder only - remove

			key := k
			// key := StrToBytes(k)
			val := StrToBytes(v)
			storage.Put(key, val)
		}
	})
}
