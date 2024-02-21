package main

import (
	"form-web/handlers"
	"net/http"
)

func main() {
	http.HandleFunc("/", handlers.SubscriptionHandler)

}
