package main

import (
	"math/rand/v2"
	"net/http"
	"strconv"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/random", func(w http.ResponseWriter, r *http.Request) {
		num := rand.IntN(6) + 1
		w.Write([]byte(strconv.Itoa(num)))
	})

	http.ListenAndServe(":8080", mux)
}
