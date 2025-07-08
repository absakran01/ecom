package user

import (
	"fmt"
	"net/http"


	"github.com/absakran01/ecom/service/auth"
	"github.com/absakran01/ecom/types"
	"github.com/absakran01/ecom/utils"
	"github.com/gorilla/mux"

)

type Handler struct {
	store types.UserStore
}

func NewHandler(store types.UserStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/login", h.handleLogin).Methods("POST")

	router.HandleFunc("/register", h.handleRegister).Methods("POST")
}

func (h *Handler) handleLogin(w http.ResponseWriter, r *http.Request) {
	//debug
	fmt.Println("Login endpoint hit")


	var payload types.LoginUserPayLoad
	err := utils.ParseJSON(r, &payload)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload: %v", err))
		return
	}

	// Validate required fields
	if  payload.Email == "" || payload.Password == "" {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("all fields are required"))
		return
	}
	user, err := h.store.GetUserByEmail(payload.Email)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("user with email %s not found", payload.Email))
		return
	}
	if !auth.CheckPasswordHash(payload.Password, user.Password) {
		utils.WriteError(w, http.StatusUnauthorized, fmt.Errorf("invalid credentials"))
		return
	}

	// Generate JWT token
	token, err := auth.GenJWT(user.ID)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("error generating JWT token: %v", err))
		return
	}	

	utils.WriteJSON(w, http.StatusOK, map[string]string{"message": "login successful", "user_name": fmt.Sprintf("%s", user.FirstName), "token": token})
}

func (h *Handler) handleRegister(w http.ResponseWriter, r *http.Request) {
	//debug
	fmt.Println("Register endpoint hit")
	var payload types.RegisterUserPayLoad
	err := utils.ParseJSON(r, &payload)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload: %v", err))
		return
	}

	// Validate required fields
	if payload.FirstName == "" || payload.LastName == "" || payload.Email == "" || payload.Password == "" {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("all fields are required"))
		return
	}
	_, err = h.store.GetUserByEmail(payload.Email)
	if err == nil {
		utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("user with email %s already exists", payload.Email))
		return
	}

	hashedPass, err := auth.HashPassword(payload.Password)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("error hashing password: %v", err))
		return
	}

	err = h.store.CreateUser(&types.User{
		FirstName: payload.FirstName,
		LastName:  payload.LastName,
		Email:     payload.Email,
		Password:  hashedPass,
	})
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("error creating user: %v", err))
		return
	}

	utils.WriteJSON(w, http.StatusCreated, map[string]string{"message": "user created successfully"})
}
