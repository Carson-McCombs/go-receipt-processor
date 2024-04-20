package main

import (
	api "GoReceiptProcessor/API"
	"net/http"
)

func main() {
	server := api.NewServer()
	http.ListenAndServe(":8080", server)
}
