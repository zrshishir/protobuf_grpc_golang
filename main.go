package main

import (
	"net/http"
	"protobuf_grpc/server"
)

func main() {
	mux := http.NewServeMux()
	server.New(mux)

	http.ListenAndServe(":8080", mux)
}
