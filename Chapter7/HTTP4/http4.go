// includes Exercises 7.11 and 7.12
package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

type dollars float32

func (d dollars) String() string { return fmt.Sprintf("$%.2f", d) }

type database map[string]dollars

func (db database) list(w http.ResponseWriter, req *http.Request) {
	//	for item, price := range db {
	//		fmt.Fprintf(w, "%s: %s\n", item, price)
	//	}
	t, _ := template.ParseFiles("list.html")
	fmt.Println(t.Execute(w, db))

}

func (db database) price(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	price, ok := db[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound) //404
		fmt.Fprintf(w, "no such item: %q\n", item)
		return
	}
	fmt.Fprintf(w, "%s\n", price)
}

func (db database) update(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	priceS := req.URL.Query().Get("price")
	price, _ := strconv.Atoi(priceS)
	if db[item] != 0 {
		db[item] = dollars(price)
	} else {
		w.WriteHeader(http.StatusNotFound) //404
		fmt.Fprintf(w, "no such item: %q\n", item)
		return
	}
}

func (db database) create(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	priceS := req.URL.Query().Get("price")
	price, _ := strconv.Atoi(priceS)
	// if item does not exists then create it with given price
	if db[item] == 0 {
		db[item] = dollars(price)
	} else {
		w.WriteHeader(http.StatusNotFound) //404
		fmt.Fprintf(w, "item: %q already exists\n", item)
		return
	}
}

func main() {
	db := database{"shoes": 50, "socks": 5}
	http.HandleFunc("/list", db.list)
	http.HandleFunc("/price", db.price)
	http.HandleFunc("/update", db.update)
	http.HandleFunc("/create", db.create)

	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
