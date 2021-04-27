package main

import (
	"log"
	"net/http"

	"github.com/ibrokethecloud/apidemo/pkg/server"
)

func main() {
	http.HandleFunc("/api", server.RequestServer)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
