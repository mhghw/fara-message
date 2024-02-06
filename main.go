package main

import (
	"flag"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/mhghw/fara-message/api"
)

var port = flag.Int("port", 8080, "Port to run the HTTP server")

func main() {
	flag.Parse()
	addr := fmt.Sprintf(":%d", *port)
	router := gin.Default()
	router.POST("/user/register", api.Register)
	router.Run(addr)

}
