package api 

import (
    "net/http"
    "encoding/json"
)

/*
 * Informational API functions to be used for boilerplate purposes
 */

func Index(w http.ResponseWriter, r *http.Request) {

    payload, err := json.Marshal("Toddler Travel API - V1")
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusOK)
    w.Write(payload)

}

