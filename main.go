package main

import (
    "log"
    "net/http"
)

/*
 * Run the API from within main on port 8080.
 */

func main() {

    router := NewRouter()
    log.Fatal(http.ListenAndServe(":8080", router))
}
