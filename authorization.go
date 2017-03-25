package main

import (
    "net/http"
    "encoding/base64"
    "strings"
)

func authorize(h http.HandlerFunc, middleware ...func(http.HandlerFunc) http.HandlerFunc) http.HandlerFunc {
	for _, m := range middleware {
		h = m(h)
	}

	return h
}

func basicAuth(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

        w.Header().Set("Access-Control-Allow-Methods", "GET")
        w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding")
        w.Header().Set("Access-Control-Allow-Credentials", "true")
        w.Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)
        w.Header().Set("Content-Type", "application/json")


		s := strings.SplitN(r.Header.Get("Authorization"), " ", 2)
		if len(s) != 2 {
			http.Error(w, "Not Authorized: Username and password required.", 401)
            return
		}

		b, err := base64.StdEncoding.DecodeString(s[1])
		if err != nil {
			http.Error(w, err.Error(), 401)
			return
		}

		pair := strings.SplitN(string(b), ":", 2)
		if len(pair) != 2 {
			http.Error(w, "Not Authorized: Username and password required.", 401)
			return
		}

		if pair[0] != "admin" || pair[1] != "admin" {
			http.Error(w, "Not Authorized: Incorrect username or password.", 401)
			return
		}

		h.ServeHTTP(w, r)
	}
}
