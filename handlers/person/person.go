package person 

import (
    "net/http"
    "encoding/json"
)

type Person struct {
    Name string `json:"name"`
    Age int64 `json:"age"`
    Orientation string `json:"orientation"`
}

func Get(w http.ResponseWriter, r *http.Request) {

    traveler := Person{"Mila",1,"Female"}

    payload, err := json.Marshal(traveler)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusOK)
    w.Write(payload)
}

