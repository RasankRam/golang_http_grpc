package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"todo-list/internal/app/requests/auth"
	"todo-list/internal/app/services"
	"todo-list/internal/utils"
)

type AuthController struct {
	service *services.AuthService
}

func NewAuthController(service *services.AuthService) *AuthController {
	return &AuthController{
		service: service,
	}
}

func (c *AuthController) Login(w http.ResponseWriter, r *http.Request) {
	loginReq, err := auth_requests.CreateLoginReq(r.Body, &r.Header)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	ip := utils.IpFromRequest(r)
	userAgent := r.Header.Get("User-Agent")

	tokenResponse, err := c.service.Login(r.Context(), loginReq.Login, loginReq.Password, ip, userAgent)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tokenResponse)
}

func (c *AuthController) Logout(w http.ResponseWriter, r *http.Request) {
	logoutReq, err := auth_requests.CreateLogoutReq(r.Body, &r.Header)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if logoutReq.Token == "" {
		token, err := utils.BearerToken(r.Header.Get("Authorization"))
		if err != nil || token == "" {
			http.Error(w, "You are not logged in", http.StatusBadRequest)
			return
		}
	}

	token, err := c.service.Logout(r.Context(), logoutReq.Token)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(token))
}

func (c *AuthController) Register(w http.ResponseWriter, r *http.Request) {
	registerReq, err := auth_requests.CreateRegisterReq(r.Body, &r.Header)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	userId, err := c.service.Register(r.Context(), registerReq.Login, registerReq.Password, registerReq.Role)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(strconv.Itoa(userId)))
}
