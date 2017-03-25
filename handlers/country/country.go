package country 

import (
    "net/http"
    "encoding/json"
    "strconv"
    "github.com/gorilla/mux"
)

type Country struct {
    ID int `json:"id"`
    Name string `json:"name"`
    Capital string `json:"capital"`
    Language string `json:"language"`
    Population int `json:"population"`
    Continent string `json:"continent"`
}

func GetAll(w http.ResponseWriter, r *http.Request) {

    countries := getCountries() 

    payload, err := json.Marshal(countries)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusOK)
    w.Write(payload)
}

func Get(w http.ResponseWriter, r *http.Request) {
    
    vars := mux.Vars(r)
    ID, _ := strconv.Atoi(vars["id"])

    countries := getCountries() 

    country := Country{} 

    for i:=0;i<len(countries);i++ {
	if countries[i].ID == ID {
	    country = countries[i]
	    break
	}
    }

    if (Country{}) != country {

	payload, err := json.Marshal(country)
	if err != nil {
	    http.Error(w, err.Error(), http.StatusInternalServerError)
	    return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(payload)
	return
    }

    w.WriteHeader(http.StatusBadRequest)
    return
}


/*
 * Statically Defined for now - will add db logic once set up.
 */

func getCountries() (countries []Country) {

    usa := Country{1,"United States","Washington DC","English",309,"North America"}
    canada := Country{2,"Canada","Ottawa","English",35,"North America"}
    germany := Country{3,"Germany","Berlin","German",80,"Europe"}
    france := Country{4,"France","Paris","French",66,"Europe"}
    japan := Country{5,"Japan","Tokyo","Japanese",127,"Asia"}

    countries = append(countries, usa)
    countries = append(countries, canada)
    countries = append(countries, germany)
    countries = append(countries, france)
    countries = append(countries, japan)

    return
}
