package api

import (
	"encoding/json"
	"net/http"

	"github.com/princetheprogrammerbtw/flowforge/internal/auth"
	"github.com/princetheprogrammerbtw/flowforge/internal/config"
	"github.com/princetheprogrammerbtw/flowforge/internal/repository"
)

type AuthHandler struct {
	userRepo *repository.UserRepository
	cfg      *config.Config
}

func NewAuthHandler(userRepo *repository.UserRepository, cfg *config.Config) *AuthHandler {
	return &AuthHandler{userRepo: userRepo, cfg: cfg}
}

type registerRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	var req registerRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		ErrorResponse(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	hashedPassword, err := auth.HashPassword(req.Password)
	if err != nil {
		ErrorResponse(w, http.StatusInternalServerError, "Could not hash password")
		return
	}

	user, err := h.userRepo.Create(r.Context(), req.Email, hashedPassword)
	if err != nil {
		ErrorResponse(w, http.StatusInternalServerError, "Could not create user")
		return
	}

	token, _ := auth.GenerateToken(user.ID, h.cfg.JWTSecret)
	JSONResponse(w, http.StatusCreated, map[string]interface{}{
		"user":  user,
		"token": token,
	})
}
