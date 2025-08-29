package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/google/uuid"
)

type User struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type UserRepository interface {
	GetAll() ([]User, error)
	Create(user User) (User, error)
}

type UserService struct {
	repo UserRepository
}

func NewUserService(repo UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) GetAllUsers() ([]User, error) {
	return s.repo.GetAll()
}

func generateUUID() string {
	return uuid.New().String()
}

func (s *UserService) CreateUser(user User) (User, error) {
	if user.Name == "" || user.Email == "" {
		return User{}, ErrValidation
	}
	user.ID = generateUUID()
	return s.repo.Create(user)
}

var ErrValidation = &APIError{Message: "Name and email are required", Status: http.StatusBadRequest}

type APIError struct {
	Message string
	Status  int
}

func (e *APIError) Error() string { return e.Message }

type UserHandler struct {
	service *UserService
}

func NewUserHandler(service *UserService) *UserHandler {
	return &UserHandler{service: service}
}

func (h *UserHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.handleGetUsers(w, r)
	case http.MethodPost:
		h.handleCreateUser(w, r)
	case http.MethodOptions:
		h.handleUsersOptions(w, r)
	default:
		w.Header().Set("Allow", "GET, POST, OPTIONS")
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}

func (h *UserHandler) handleGetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := h.service.GetAllUsers()
	if err != nil {
		h.respondError(w, err)
		return
	}
	h.respondJSON(w, users, http.StatusOK)
}

func (h *UserHandler) handleCreateUser(w http.ResponseWriter, r *http.Request) {
	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	created, err := h.service.CreateUser(user)
	if err != nil {
		h.respondError(w, err)
		return
	}

	h.respondJSON(w, created, http.StatusCreated)
}

func (h *UserHandler) handleUsersOptions(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Allow", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.WriteHeader(http.StatusNoContent)
}

func (h *UserHandler) respondJSON(w http.ResponseWriter, data any, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Printf("Error encoding response: %v", err)
	}
}

func (h *UserHandler) respondError(w http.ResponseWriter, err error) {
	if apiErr, ok := err.(*APIError); ok {
		http.Error(w, apiErr.Message, apiErr.Status)
	} else {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
