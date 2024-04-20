package main

import (
	api "go-receipt-processor/API"
	"net/http"
)

func main() {
	server := api.NewServer()
	http.ListenAndServe(":8080", server)
}
