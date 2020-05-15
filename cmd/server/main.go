package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/natasharw/lsmtree-kvstore/pkg/storage"
)

func main() {
	fmt.Println("Welcome to the key-value store. Initialising...")
	var store storage.Store
	store = storage.NewKVStore()
	fmt.Println("Ready")
	run(store)
	defer fmt.Println("Exiting key-value store")
}

func run(store storage.Store) {
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter command: ")
		cmd, err := reader.ReadString('\n')
		if err != nil {
			panic(err)
		}
		log.Printf("Input received: %s", cmd)
		process(cmd, store)
	}
}

func process(cmd string, store storage.Store) {
	c := strings.Fields(cmd)
	switch c[0] {
	case "get":
		if len(c[1:]) != 1 {
			fmt.Println("Incorrect command supplied. please use \"get <yourkey>\"")
			break
		}
		log.Printf("Processing get request")
		key := c[1]
		store.Get(key)
	case "set":
		if len(c[1:]) != 2 {
			fmt.Println("Incorrect command supplied. please use \"set <yourkey> <yourvalue>\"")
			break
		}
		log.Printf("Processing set request")
		key, val := c[1], c[2]
		store.Set(key, val)
	case "exit":
		log.Printf("Thanks for visiting the key-value store.")
		os.Exit(0)
	default:
		fmt.Println("Incorrect command supplied. hint: see README and try again")
	}
}
