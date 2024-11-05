package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// Model
type Product struct {
	Id       int     `json:"id"`
	Name     string  `json:"name"`
	Cost     float64 `json:"cost"`
	Category string  `json:"-"`
}

// Inmemory data
var products []Product = []Product{
	Product{Id: 100, Name: "Pen", Cost: 10, Category: "stationary"},
	Product{Id: 101, Name: "Pencil", Cost: 5, Category: "stationary"},
	Product{Id: 102, Name: "Marker", Cost: 50, Category: "stationary"},
}

type AppServer struct {
}

func (appServer *AppServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s - %s\n", r.Method, r.URL.Path)
	switch r.URL.Path {
	case "/":
		fmt.Fprintln(w, "Welcome to the apis!")
	case "/products":
		switch r.Method {
		case http.MethodGet:
			if payload, err := json.Marshal(products); err != nil {
				http.Error(w, "interval server error", http.StatusInternalServerError)
				return
			} else {
				fmt.Fprint(w, string(payload))
			}
		case http.MethodPost:
			var newProduct Product
			if err := json.NewDecoder(r.Body).Decode(&newProduct); err != nil {
				http.Error(w, "request payload error", http.StatusBadRequest)
			}
			newProduct.Id = len(products) + 100
			products = append(products, newProduct)
			w.WriteHeader(http.StatusCreated)
			if err := json.NewEncoder(w).Encode(newProduct); err != nil {
				http.Error(w, "interval server error", http.StatusInternalServerError)
			}
		}

	case "/users":
		fmt.Fprintln(w, "Users list will be served")
	default:
		http.Error(w, "resource not found", http.StatusNotFound)
	}
}

func main() {
	appServer := &AppServer{}
	if err := http.ListenAndServe(":8080", appServer); err != nil {
		log.Fatalln("Error starting server :", err)
	}
}
