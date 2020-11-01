// adding JSON
package main

import (
  "net/http"
  "log"
	"github.com/gorilla/mux"
	"fmt"
  "time"
    "encoding/json"
)

func HelloFriend(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    w.WriteHeader(http.StatusOK)
    fmt.Fprintf(w, "Hello %v!!!\n", vars["name"])
}

type MessageStruct struct {
    Message string
}

func HelloWorld(w http.ResponseWriter, r *http.Request) {
    msg := MessageStruct {
        Message: "Hello World!!!",
    }
    b, err := json.Marshal(msg)
    if(err != nil) {
        w.Write([]byte("Hello World!!!"))
    } else {
        w.Write(b)
    }
}

func main() {
    r := mux.NewRouter()
    // Routes consist of a path and a handler function.
    r.HandleFunc("/{name}", HelloFriend)
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
