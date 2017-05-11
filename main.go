package main

import (
    "log"
    "net/http"
)

/*
 * Run the API from within main on port 8080.
 */

func main() {

    log.Println("Starting Toddler Travel API V1.0.0")

    router := NewRouter()
    log.Fatal(http.ListenAndServe(":8080", router))
}
