package main

import (
	"broker/cmd/api"
	"broker/cmd/util"
	"fmt"
	"net/http"
)

var port string = util.AssertEnv("SERVICE_PORT")

func main() {
	fmt.Printf("Starting Broker service on port %s\n", port)

	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: api.Routes(),
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
