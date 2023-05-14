package main

import (
	"fmt"
	"net/http"
)

func main() {
	originalHandler := http.HandlerFunc(requestHandler)
	http.Handle("/", middleware(originalHandler))
	http.ListenAndServe(":8000", nil)
}

func requestHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Executing main request handler")
	w.Write([]byte("OK"))
}

func middleware(originalHandler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Executing before the request phase")
		w.Write([]byte("REQUEST HIJACKED "))
		originalHandler.ServeHTTP(w, r)
		fmt.Println("Executing after the request phase")
	})
}
