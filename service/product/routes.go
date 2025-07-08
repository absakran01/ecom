package product

import (
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"github.com/absakran01/ecom/utils"
	"github.com/absakran01/ecom/types"
	"fmt"
)

type handler struct {
	store *Store
}

func NewHandler(store *Store) *handler {
	return &handler{store: store}
}

func (h *handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/products", h.handleGetProducts).Methods("GET")
	router.HandleFunc("/products/{id}", h.handleGetProductByID).Methods("GET")
	router.HandleFunc("/products", h.handleCreateProduct).Methods("POST")
}

func (h *handler) handleGetProducts(w http.ResponseWriter, r *http.Request) {
	//debug
	fmt.Println("Get Products endpoint hit")

	products, err := h.store.GetProducts()
	if err != nil {
		http.Error(w, "Error fetching products", http.StatusInternalServerError)
		return
	}
	utils.WriteJSON(w, http.StatusOK, products)
}
func (h *handler) handleGetProductByID(w http.ResponseWriter, r *http.Request) {
	//debug
	fmt.Println("Get Product by ID endpoint hit")

	idStr := mux.Vars(r)["id"]
	if idStr == "" {
		http.Error(w, "Product ID is required", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	product, err := h.store.GetProductByID(id)
	if err != nil {
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}
	utils.WriteJSON(w, http.StatusOK, product)
}


func (h *handler) handleCreateProduct(w http.ResponseWriter, r *http.Request) {
	//debug
	fmt.Println("Create Product endpoint hit")
	
	var product *types.CreateProductPayLoad
	if err := utils.ParseJSON(r, &product); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if err := h.store.CreateProduct(product); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusCreated, product)
}

