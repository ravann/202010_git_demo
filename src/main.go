// adding JSON
package main

import (
    "net/http"
    "log"
    "github.com/gorilla/mux"
    "encoding/json"
)

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
    r.HandleFunc("/", HelloWorld)

    // Bind to a port and pass our router in
    log.Fatal(http.ListenAndServe(":8000", r))
}
