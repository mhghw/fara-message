package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/mhghw/fara-message/api"
	"github.com/mhghw/fara-message/db"
	"github.com/rs/xid"
)

// implement this with os args
var port = flag.Int("port", 8080, "Port to run the HTTP server")

func main() {
	guid := xid.New()
	fmt.Println(guid.String())
	db.NewDatabase()
	flag.Parse()
	err := api.RunWebServer(*port)
	if err != nil {
		log.Print("failed to start HTTP server:", err)
	}

}
