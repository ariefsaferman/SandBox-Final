package main

import (
	"go-basic-server/database"
	"go-basic-server/handler"
	"go-basic-server/repository"
	"go-basic-server/usecase"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// type Digitalent struct {
// 	ListName []*Name `json:"digitalent"`
// }

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

type Response struct {
	StatusCode int    `json:"statusCode"`
	Status     string `json:"status"`
	Message    string `json:"message"`
}

// var products []*Product

// var database *sql.DB

func main() {
	r := mux.NewRouter()
	DB := database.InitDb()
	defer DB.Close()

	repo := repository.NewProductRepository(DB)
	use := usecase.NewProductUsecase(repo)
	hand := handler.NewHandler(use)

	// r.HandleFunc("/digitalent", getAllDigitalent).Methods(http.MethodGet)
	r.HandleFunc("/product", hand.GetAllProduct).Methods(http.MethodGet)
	r.HandleFunc("/product/{id}", hand.GetProduct).Methods(http.MethodGet)
	// r.HandleFunc("/product/{id}/category", hand.CreateProduct).Methods(http.MethodPost)
	r.HandleFunc("/product", hand.CreateProduct).Methods(http.MethodPost)
	r.HandleFunc("/product/{id}", hand.UpdateProduct).Methods(http.MethodPut)
	r.HandleFunc("/product/{id}", hand.DeleteProduct).Methods(http.MethodDelete)
	log.Print("server is running on 8080")
	http.ListenAndServe(":8080", r)
}

// func responJson(w http.ResponseWriter, status int, err error) {
// 	msg := "success"
// 	if err != nil {
// 		msg = err.Error()
// 	}

// 	resp := Response{
// 		StatusCode: status,
// 		Status:     http.StatusText(status),
// 		Message:    msg,
// 	}

// 	jsonResp, err := json.Marshal(resp)

// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 	}

// 	w.Write(jsonResp)
// }

// func getAllDigitalent(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(201)
// 	listName := Digitalent{
// 		ListName: []*Name{
// 			{"arief"},
// 			{"yan"},
// 			{"rei"},
// 			{"andra"},
// 			{"merissa"},
// 		},
// 	}
// 	json.NewEncoder(w).Encode(listName)
// }

// func findOrCreateProduct(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")

// 	switch r.Method {
// 	case "POST":
// 		var newProduct Product
// 		json.NewDecoder(r.Body).Decode(&newProduct)
// 		products = append(products, &newProduct)
// 		json.NewEncoder(w).Encode(&products)
// 		responJson(w, http.StatusCreated, nil)
// 		return
// 	case "GET":
// 		w.Header().Add("Content-Type", "application/json")
// 		name := r.URL.Query().Get("name")

// 		switch name {
// 		case "":
// 			var listProduct []*Product
// 			listProduct = GetAllProductQuery(database)
// 			json.NewEncoder(w).Encode(listProduct)
// 			responJson(w, http.StatusAccepted, nil)
// 			return
// 		default:
// 			//not implemented
// 			// product := GetProductByNameQuery(database)
// 			// json.NewEncoder(w).Encode(product)
// 		}
// 	}
// }

// func getProductById(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	params := mux.Vars(r)
// 	id := params["id"]
// 	intId, _ := strconv.Atoi(id)

// 	product := GetProductByIdQuery(database, intId)
// 	json.NewEncoder(w).Encode(&product)
// }

// func addCategory(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	params := mux.Vars(r)
// 	id := params["id"]
// 	intId, _ := strconv.Atoi(id)

// 	var product *Product

// 	for _, prod := range products {
// 		if prod.Id == intId {
// 			product = prod
// 			break
// 		}
// 	}

// 	var category *Category
// 	json.NewDecoder(r.Body).Decode(&category)
// 	product.Category = category
// 	responJson(w, http.StatusOK, nil)

// }

// func updateProduct(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	params := mux.Vars(r)
// 	id := params["id"]
// 	intId, _ := strconv.Atoi(id)

// 	for _, product := range products {
// 		var updateProduct Product
// 		json.NewDecoder(r.Body).Decode(&updateProduct)
// 		if product.Id == intId {
// 			if updateProduct.Name != "" {
// 				product.Name = updateProduct.Name
// 			}

// 			if updateProduct.Description != "" {
// 				product.Description = updateProduct.Description
// 			}

// 			if updateProduct.Category != nil {
// 				product.Category = updateProduct.Category
// 			}
// 			responJson(w, http.StatusAccepted, nil)
// 			return
// 		}
// 	}
// }

// func deleteProduct(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	params := mux.Vars(r)
// 	id := params["id"]
// 	intId, _ := strconv.Atoi(id)

// 	for idx, product := range products {
// 		if product.Id == intId {
// 			products = append(products[:idx], products[idx+1:]...)
// 			responJson(w, http.StatusAccepted, nil)
// 			return
// 		}
// 	}
// }
