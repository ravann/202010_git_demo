package main

import (
    "net/http"
    "log"
	"github.com/gorilla/mux"
    "time"
)

func HelloWorld(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Hello World!\n"))
}

func main() {
    r := mux.NewRouter()
    // Routes consist of a path and a handler function.
    r.HandleFunc("/", HelloWorld)

    server := &http.Server{
        Addr:           ":8000",
        Handler:        r,
        ReadTimeout:    5 * time.Second,
        WriteTimeout:   5 * time.Second,
    }

    // Bind to a port and pass our router in
    log.Fatal(server.ListenAndServe())
}
