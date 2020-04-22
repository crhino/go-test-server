package main

import (
	"flag"
	"fmt"
	"net/http"
)

var port = flag.Int("port", 1234, "specify the port to listen on")

func main() {
	http.HandleFunc("/", HelloServer)
	err := http.ListenAndServe(fmt.Sprintf(":%d", *port), nil)
	if err != nil {
		fmt.Printf("Error running listener: %s", err)
	}
}

func HelloServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
}
