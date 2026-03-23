package main

import (
	randomapi "httpserver/2-random-api"
)

func main() {
	// router := http.NewServeMux()
	// NewHelloHandler(router)

	// server := http.Server{
	// 	Addr:    ":8081",
	// 	Handler: router,
	// }

	// fmt.Println("Server is listening on port 8081")
	// server.ListenAndServe()
	randomapi.StartServer()
}
