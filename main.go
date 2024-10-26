package main

import (
	"log"
	"net/http"
	"root/cmd"
)

func main() {
	port := cmd.Load()

	fileServer := http.FileServer(http.Dir("./client"))
	http.Handle("/", fileServer)

	if err := http.ListenAndServe(port, nil); err != nil {
		log.Println(err)
	}
}
