package handler

import (
	"encoding/json"
	"go-basic-server/entity"
	"go-basic-server/usecase"
	"go-basic-server/util"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type ProductHandler interface {
	GetProduct(w http.ResponseWriter, r *http.Request)
	GetAllProduct(w http.ResponseWriter, r *http.Request)
	CreateProduct(w http.ResponseWriter, r *http.Request)
	UpdateProduct(w http.ResponseWriter, r *http.Request)
	DeleteProduct(w http.ResponseWriter, r *http.Request)
}

type ProductHandlerImpl struct {
	usecase usecase.ProductUsecase
}

func NewHandler(usecase usecase.ProductUsecase) ProductHandler {
	return &ProductHandlerImpl{
		usecase: usecase,
	}
}

func (u *ProductHandlerImpl) GetProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id := params["id"]
	intId, _ := strconv.Atoi(id)

	product, err := u.usecase.GetOneProductById(intId)
	if err != nil {
		json.NewEncoder(w).Encode(http.StatusExpectationFailed)
	}
	json.NewEncoder(w).Encode(&product)
}

func (u *ProductHandlerImpl) GetAllProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	products, _ := u.usecase.GetAllProduct()
	json.NewEncoder(w).Encode(&products)
}

func (u *ProductHandlerImpl) CreateProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var p entity.Product
	json.NewDecoder(r.Body).Decode(&p)
	err := u.usecase.CreateProduct(&p)
	if err != nil {
		panic(err)
	}
	util.ResponJson(w, http.StatusCreated, err)
}

func (u *ProductHandlerImpl) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id := params["id"]
	intId, _ := strconv.Atoi(id)
	var p entity.Product
	json.NewDecoder(r.Body).Decode(&p)

	err := u.usecase.UpdateProduct(intId, &p)
	if err != nil {
		panic(err)
	}
	util.ResponJson(w, http.StatusOK, err)
}

func (u *ProductHandlerImpl) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id := params["id"]
	intId, _ := strconv.Atoi(id)

	err := u.usecase.DeleteProduct(intId)
	if err != nil {
		panic(err)
	}
	util.ResponJson(w, http.StatusOK, err)
}
