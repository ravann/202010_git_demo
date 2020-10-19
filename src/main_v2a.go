package main

import (
    "net/http"
    "log"
	"github.com/gorilla/mux"
	"fmt"
)

func HelloFriend(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    w.WriteHeader(http.StatusOK)
    fmt.Fprintf(w, "Hello %v!!!\n", vars["name"])
}

func HelloWorld(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK)
    w.Write([]byte("Hello World!\n"))
}

func main() {
    r := mux.NewRouter()
    // Routes consist of a path and a handler function.
    r.HandleFunc("/{name}", HelloFriend)
    r.HandleFunc("/", HelloWorld)

    // Bind to a port and pass our router in
    log.Fatal(http.ListenAndServe(":8000", r))
}
