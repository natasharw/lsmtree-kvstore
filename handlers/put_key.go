package handlers

import (
	"log"
	"net/http"
	"net/url"

	"github.com/natasharw/lsmtree-kvstore/storage"
)

// PutKey : takes a key-value pair passed in as a parameter and passes it the store, printing message to user
func PutKey(buffer *storage.Memtable) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		log.Println("Handling set key request")
		log.Println("Params provided:", req.URL.Query())

		log.Println("Parsing query: ", req.URL.Query())
		raw := req.URL.RawQuery
		parsed, err := url.ParseQuery(raw)
		log.Println("Parsed: ", parsed)

		defer req.Body.Close()

		if len(parsed) == 0 {
			http.Error(w, "invalid SET request. hint: add a parameter of yourkey=yourvalue", http.StatusBadRequest)
		}

		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		// iterate over all passed key-value pairs
		i := 1
		ttl := len(parsed)
		for k, v := range parsed {
			log.Printf("Handling key-value pair %d of total pairs %d", i, ttl)
			log.Printf("Key: %s. Value: %s", k, v)

			value := StrToBytes(v)
			log.Printf("Passing key-value to store")
			storage.Put(buffer, k, value)
			i++
		}
	})
}
