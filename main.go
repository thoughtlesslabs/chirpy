package main

import (
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(http.Dir(".")))
	serv := http.Server{Handler: mux, Addr: ":8080"}
	serv.ListenAndServe()
}
