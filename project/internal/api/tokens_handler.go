package api

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/vinayakvispute/project/internal/store"
	"github.com/vinayakvispute/project/internal/tokens"
	"github.com/vinayakvispute/project/internal/utils"
)

type TokenHandler struct {
	tokenStore store.TokenStore
	userStore  store.UserStore
	logger     *log.Logger
}

type createTokenRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func NewTokenHandler(tokenStore store.TokenStore, userStore store.UserStore, logger *log.Logger) *TokenHandler {

	return &TokenHandler{
		tokenStore: tokenStore,
		userStore:  userStore,
		logger:     logger,
	}
}

func (h *TokenHandler) HandleCreateToken(w http.ResponseWriter, r *http.Request) {

	var req createTokenRequest
	err := json.NewDecoder(r.Body).Decode(&req)

	if err != nil {
		h.logger.Printf("ERROR: createTokenRequest: %v", err)
		utils.WriteJson(w, http.StatusBadGateway, utils.Envelop{"error": "invalid request payload"})
		return
	}

	// lets get the users

	user, err := h.userStore.GetUserByUsername(req.Username)
	if err != nil || user == nil {
		h.logger.Printf("ERROR: GetUserByUsername: %v", err)
		utils.WriteJson(w, http.StatusInternalServerError, utils.Envelop{"error": "internal server error "})
		return
	}

	passwordDoMatch, err := user.PasswordHash.Matches(req.Password)
	if err != nil {
		h.logger.Printf("ERROR: PasswordHash.Matches: %v", err)
		utils.WriteJson(w, http.StatusInternalServerError, utils.Envelop{"error": "internal server error"})
		return
	}

	if !passwordDoMatch {
		utils.WriteJson(w, http.StatusUnauthorized, utils.Envelop{"error": "invalid credentials"})
		return
	}

	token, err := h.tokenStore.CreateNewToken(user.ID, 24*time.Hour, tokens.ScopeAuth)
	if err != nil {
		h.logger.Printf("ERROR: CreateNewToke: %v", err)
		utils.WriteJson(w, http.StatusInternalServerError, utils.Envelop{"error": "internal server error "})
		return
	}

	utils.WriteJson(w, http.StatusCreated, utils.Envelop{"auth_token": token})
}
