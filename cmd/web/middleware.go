package main

import (
	"fmt"
	"net/http"
)

func PrintToConsole(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Serve the next page")
		next.ServeHTTP(w, r)
	})
}