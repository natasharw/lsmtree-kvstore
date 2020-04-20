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
	store := storage.NewKVStore()
	fmt.Println("Ready")
	run()
	defer fmt.Println("Exiting key-value store")
}

func run() {
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter command: ")
		cmd, err := reader.ReadString('\n')
		if err != nil {
			panic(err)
		}
		log.Printf("Input received: %s", cmd)
		process(cmd)
	}
}

func process() {
	c := strings.Fields(cmd)
	switch c[0] {
	case "get":
		if len(c) != 1 {
			fmt.Print("Incorrect command supplied. please use \"get <yourkey>\"")
		}
		log.Printf("Processing get request")
		key := c[1]
		store.Get(key)
	case "set":
		if len(c) != 2 {
			fmt.Print("Incorrect command supplied. please use \"set <yourkey> <yourvalue>\"")
		}
		log.Printf("Processing set request")
		key := c[1]
		val := c[2]
		store.Set(key, val)
	case "exit":
		log.Printf("Thanks for visiting the key-value store. Exiting.")
		os.Exit(0)
	default:
		fmt.Print("Incorrect command supplied. hint: see README and try again")
	}
}
