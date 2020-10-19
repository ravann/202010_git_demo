package main

import (
    "net/http"
    "log"
    "github.com/gorilla/mux"
)

func HelloWorld(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Hello World!\n"))
}

func main() {
    r := mux.NewRouter()
    // Routes consist of a path and a handler function.
    r.HandleFunc("/", HelloWorld)

    // Bind to a port and pass our router in
    log.Fatal(http.ListenAndServe(":8000", r))
}
