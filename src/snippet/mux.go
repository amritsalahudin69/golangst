package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", ParamHandlerWithoutInput)
	r.HandleFunc("/{key}", ParamHandler)

	http.ListenAndServe(":8080", r)
}

func ParamHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)

	fmt.Fprint(w, "Input Param : %v\n", vars["key"])
}

func ParamHandlerWithoutInput(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Success Ok")
}
