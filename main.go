package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Item struct {
	ID     string `json:"id"`
	a      int    `json:"a"`
	b      int    `json:"b"`
	c      int    `json:"c"`
	nRoots int    `json:"nRoots"` // количество корней функции
}

//Items ...
var Items []Item = []Item{
	Item{"1", 1, 2, 3, 6},
}

//if x = 1
//(1*1)**2 + 2*1 + 3 = 6    NOT SURE !

//GetItems ...
func GetItems(w http.ResponseWriter, r *http.Request) { //возвращаем все аргументы квадратной функции с кол-вом корней
	json.NewEncoder(w).Encode(Items)
}

//PostItem ...
func PostItem(w http.ResponseWriter, r *http.Request) {
	a := strconv.Atoi(mux.Vars()["a"]) //[???]
	b := strconv.Atoi(mux.Vars()["b"])
	c := strconv.Atoi(mux.Vars()["c"])

	discr := (b * *2) - (4 * a * c) //находим дискриминант			[???]

	var count []int

	// если дискр больше 0, то 2 корня
	if discr > 0 {
		x1 := (float64(-b) + math.Sqrt(discr)) / (2 * float64(a))
		x2 := (float64(-b) - math.Sqrt(discr)) / (2 * float64(a))
		count = append(count, x1, x2)
		fmt.Printf("%d", len(count))
	} else if discr == 0 { //если дискр равен 0 то 1 корень
		x := -b / (2 * a)
		count = append(count, x)
		fmt.Printf("%d", len(count))
	} else { //если дискр меньше 0 то корней нет
		fmt.Printf("%d", len(count))
	}
	reqBody, _ := ioutil.ReadAll(r.Body)
	var item Item
	json.Unmarshal(reqBody, &item)
	w.WriteHeader(http.StatusCreated)
	Items = append(Items, item)
}

func main() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/solve/{a}/{b}/{c}", GetItems).Methods("GET")

	router.HandleFunc("/solution", PostItem).Methods("POST")

	log.Fatal(http.ListenAndServe(":8000", router))
}
