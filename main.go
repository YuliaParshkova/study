package main

import (
	"encoding/json"
	"fmt"
	"github.com/fatih/structs"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
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
	a       int `json:"A"`
	b       int `json:"B"`
	c       int `json:"C"`
	n_roots int `json:"Roots"`
}

var ParamsAnswers []ParamsAnswer
var m map[string]interface{}

func GetParameters(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var par Parameters
	json.Unmarshal(reqBody, &par)

	m = structs.Map(par)
	CalcResult()
}

func GetLastResult(w http.ResponseWriter, r *http.Request) {
	if len(ParamsAnswers) == 0 {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("Ничего не нашли")
	} else {
		json.NewEncoder(w).Encode(ParamsAnswers[len(ParamsAnswers)-1])
	}
}

func CalcResult() {
	a := m["a"].(int)
	b := m["b"].(int)
	c := m["c"].(int)

	var n_roots int

	if (a == 0 && b != 0) || (a != 0 && c == 0 && b == 0) || (a == b && c == 0) {
		n_roots = 1
		return
	} else if a == 0 && b == 0 {
		n_roots = 0
		return
	} else {
		D := b*b - 4*a*c
		if D < 0 {
			n_roots = 0
		} else if D > 0 {
			n_roots = 2
		} else {
			n_roots = 1
		}
	}

	p := ParamsAnswer{
		a:       a,
		b:       b,
		c:       c,
		n_roots: n_roots,
	}
	ParamsAnswers = append(ParamsAnswers, p)
}
