package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

func main() {

	myRouter := mux.NewRouter().StrictSlash(true)
	fmt.Println("REST API worked....")

	myRouter.HandleFunc("/parameters/{a}/{b}/{c}", GetParameters).Methods("POST")
	myRouter.HandleFunc("/parametersRes", GetLastResult).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", myRouter))
}

type Parameters struct {
	a int `json:"A"`
	b int `json:"B"`
	c int `json:"C"`
}

type ParamsAnswer struct {
	A      int `json:"A"`
	B      int `json:"B"`
	C      int `json:"C"`
	Nroots int `json:"N_Roots"`
}

var ParamsAnswers []ParamsAnswer
var Params Parameters

func GetParameters(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var err error
	Params.a, err = strconv.Atoi(vars["a"])
	Params.b, err = strconv.Atoi(vars["b"])
	Params.c, err = strconv.Atoi(vars["c"])
	if err != nil {
		json.NewEncoder(w).Encode("ошибка")
	}

	CalcResult()
}

func GetLastResult(w http.ResponseWriter, r *http.Request) {
	if len(ParamsAnswers) == 0 {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("Ничего нет")
	} else {
		json.NewEncoder(w).Encode(ParamsAnswers[len(ParamsAnswers)-1])
	}
}

func CalcResult() {
	a := Params.a
	b := Params.b
	c := Params.c

	var Nroots int

	if (a == 0 && b != 0) || (a != 0 && c == 0 && b == 0) || (a == b && c == 0) {
		Nroots = 1
	} else if a == 0 && b == 0 {
		Nroots = 0
	} else {
		D := b*b - 4*a*c
		if D < 0 {
			Nroots = 0
		} else if D > 0 {
			Nroots = 2
		} else {
			Nroots = 1
		}
	}

	p := ParamsAnswer{
		A:      a,
		B:      b,
		C:      c,
		Nroots: Nroots,
	}
	ParamsAnswers = append(ParamsAnswers, p)
}
