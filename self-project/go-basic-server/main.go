package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

type Digitalent struct {
	ListName []*Name `json:"digitalent"`
}

type Name struct {
	Name string `json:"name"`
}

type Product struct {
	Id          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Category    *Category `json:"category,omitempty"`
}

type Category struct {
	Name string `json:"name,omitempty"`
}

var products []*Product

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/digitalent", getAllDigitalent).Methods(http.MethodGet)
	// r.HandleFunc("/digitalent/{name}", getDigitalentByName).Methods(http.MethodGet)
	r.HandleFunc("/product", findOrCreateProduct)
	r.HandleFunc("/product/{id}", getProductById).Methods(http.MethodGet)
	r.HandleFunc("/product/{id}/category", addCategory).Methods(http.MethodPost)
	r.HandleFunc("/product/{id}", updateProduct).Methods(http.MethodPut)
	r.HandleFunc("/product/{id}", deleteProduct).Methods(http.MethodDelete)
	log.Print("server is running on 8080")
	http.ListenAndServe(":8080", r)
}

func getAllDigitalent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)
	listName := Digitalent{
		ListName: []*Name{
			{"arief"},
			{"yan"},
			{"rei"},
			{"andra"},
			{"merissa"},
		},
	}
	json.NewEncoder(w).Encode(listName)
}

func findOrCreateProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case "POST":
		var newProduct Product
		json.NewDecoder(r.Body).Decode(&newProduct)
		products = append(products, &newProduct)
		json.NewEncoder(w).Encode(&products)
		w.Write([]byte(`{ "message": "new product is accepted" }`))
		return
	case "GET":
		w.Header().Add("Content-Type", "application/json")
		name := r.URL.Query().Get("name")
		for _, value := range products {
			if strings.Contains(value.Name, name) {
				w.WriteHeader(201)
				json.NewEncoder(w).Encode(value)
				return
			}

		}
		w.Write([]byte(`{ "message": "product is not here" }`))
		return
	}
}

func getProductById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	id := params["id"]
	intId, _ := strconv.Atoi(id)

	for _, product := range products {
		if product.Id == intId {
			json.NewEncoder(w).Encode(&product)
			return
		}
	}
	w.Write([]byte(`{ "message": "product is not found" }`))
}

func addCategory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id := params["id"]
	intId, _ := strconv.Atoi(id)

	var product *Product

	for _, prod := range products {
		if prod.Id == intId {
			product = prod
			break
		}
	}

	var category *Category
	json.NewDecoder(r.Body).Decode(&category)
	product.Category = category
	w.Write([]byte(`{ "message": "category is added" }`))

}

func updateProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id := params["id"]
	intId, _ := strconv.Atoi(id)

	for _, product := range products {
		var updateProduct Product
		json.NewDecoder(r.Body).Decode(&updateProduct)
		if product.Id == intId {
			if updateProduct.Name != "" {
				product.Name = updateProduct.Name
			}

			if updateProduct.Description != "" {
				product.Description = updateProduct.Description
			}

			if updateProduct.Category != nil {
				product.Category = updateProduct.Category
			}
			return
		}
	}
}

func deleteProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id := params["id"]
	intId, _ := strconv.Atoi(id)

	for idx, product := range products {
		if product.Id == intId {
			products = append(products[:idx], products[idx+1:]...)
		}
	}
}
