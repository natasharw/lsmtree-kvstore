package main

import (
	"log"
	"net/http"

	"github.com/natasharw/lsmtree-kvstore/handlers"
	"github.com/natasharw/lsmtree-kvstore/storage"
)

func main() {
	buffer := storage.MemtableInit()

	mux := http.NewServeMux()

	mux.Handle("/", handlers.DefaultPage(buffer)) // [TODO] - change so not called superflously from other endpoints
	mux.Handle("/get", handlers.GetKey(buffer))
	mux.Handle("/set", handlers.PutKey(buffer))

	log.Println("Listening on port 8010....")
	log.Fatal(http.ListenAndServe(":8010", mux))
}
