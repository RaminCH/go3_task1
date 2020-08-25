package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"net/http"

	"github.com/gorilla/mux"
)

type Item struct {
	ID     string `json:"id"`
	a      int    `json:"a"`
	b      int    `json:"b"`
	c      int    `json:"c"`
	result int    `json:"result"`
}

//Items ...
var Items []Item = []Item{
	Item{"1", 1, 2, 3, 6},
}

//if x = 1
//(1*1)**2 + 2*1 + 3 = 6    NOT SURE !

//GetItems ...
func GetItems(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(Items)
}

//Sqr func...
func Sqr(a, b, c int) int {
	discr := (b * *2) - (4 * a * c)

	if discr > 0 {
		x1 := (float64(-b) + math.Sqrt(discr)) / (2 * float64(a))
		x2 := (float64(-b) - math.Sqrt(discr)) / (2 * float64(a))
	} else if discr == 0 {
		x := -b / (2 * a)
	} else {
		fmt.Println("Корней нет")
	}
	return discr
}

//PostItem ...
func PostItem(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var item Item
	json.Unmarshal(reqBody, &item)
	w.WriteHeader(http.StatusCreated)
	Items = append(Items, item)
}

func main() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/items/{a}/{b}/{c}", GetItems).Methods("GET")

	router.HandleFunc("/item", PostItem).Methods("POST")

	log.Fatal(http.ListenAndServe(":8000", router))
}
