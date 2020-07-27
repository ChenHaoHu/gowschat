package main

import (
	"gows/server"
	"log"
)

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	server.StartServer(8081)
}
