package randomapi

import (
	"math/rand/v2"
	"net/http"
	"strconv"
)

func MyHandler(w http.ResponseWriter, r *http.Request) {
	num := rand.IntN(6) + 1
	w.Write([]byte(strconv.Itoa(num)))
}

func MyServer() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/random", MyHandler)
	return mux
}

func StartServer() {
	http.ListenAndServe(":8080", MyServer())
}
