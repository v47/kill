package main

import (
	"net/http"
	"fmt"
)

func Add(x, y int) int {
	return x+y
}

func EditHandlerOk(w http.ResponseWriter, r *http.Request) {
	fmt.Println("done ok")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("text"))
}