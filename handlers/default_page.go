package handlers

import (
	"fmt"
	"net/http"
)

func DefaultPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "welcome to the key-value store homepage. supported endpoints: /set?yourkey=yourvalue and /get?key=yourkey")
	w.WriteHeader(http.StatusOK)
}

// func DefaultPage(h http.Handler) http.HandlerFunc {
// 	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
// 		fmt.Fprintf(w, "welcome to the key-value store homepage. supported endpoints: /set?yourkey=yourvalue and /get?key=yourkey")
// 		defer req.Body.Close()
// 		w.WriteHeader(http.StatusOK)

// 		h.ServeHTTP(w, req)
// 	})
// }
