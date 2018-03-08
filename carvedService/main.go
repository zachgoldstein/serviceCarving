package main

import _ "net/http/pprof"

import (
	"fmt"
	"log"
	"net/http"

	"github.com/zachgoldstein/serviceCarving/carvedService/rpc"
	"github.com/zachgoldstein/serviceCarving/carvedService/server"
)

func main() {
	fmt.Printf("Superfast Counting Service running on :6666")

	go func() {
		log.Println(http.ListenAndServe("localhost:8080", nil))
	}()

	server := server.NewServer()
	twirpHandler := rpc.NewCarvedServer(server, nil)

	log.Fatal(http.ListenAndServe(":6666", twirpHandler))
}
