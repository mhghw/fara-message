package main

import (
	"flag"

	"github.com/mhghw/fara-message/api"
)

var port = flag.Int("port", 8080, "Port to run the HTTP server")

func main() {
	flag.Parse()
	api.WebServer(port)

}
