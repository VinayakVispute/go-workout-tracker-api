package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/vinayakvispute/project/internal/store"
	"github.com/vinayakvispute/project/internal/tokens"
	"github.com/vinayakvispute/project/internal/utils"
)

type UserMiddleware struct {
	UserStore store.UserStore
}

type contextKey string

// doubt
const UserContextKey = contextKey("user")

func SetUser(r *http.Request, user *store.User) *http.Request {

	ctx := context.WithValue(r.Context(), UserContextKey, user)

	return r.WithContext(ctx)
}

func GetUser(r *http.Request) *store.User {
	user, ok := r.Context().Value(UserContextKey).(*store.User)

	if !ok {
		panic("missing user in the request")
	}

	return user
}

func (um *UserMiddleware) Authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// within this anonymouse function
		// we can interject any incoming requests to our server

		w.Header().Add("Vary", "Authorization")
		authHeader := r.Header.Get("Authorization")

		if authHeader == "" {
			r = SetUser(r, store.AnonymousUser)
			next.ServeHTTP(w, r)
			return
		}

		headerParts := strings.Split(authHeader, " ") // Bearer <TOKEN>
		if len(headerParts) != 2 || headerParts[0] != "Bearer" {
			utils.WriteJson(w, http.StatusUnauthorized, utils.Envelop{"error": "invalid authorization header"})
			return
		}

		token := headerParts[1]
		user, err := um.UserStore.GetUserToken(tokens.ScopeAuth, token)
		if err != nil {
			utils.WriteJson(w, http.StatusUnauthorized, utils.Envelop{"error": "invalid token"})
			return
		}

		if user == nil {
			utils.WriteJson(w, http.StatusUnauthorized, utils.Envelop{"error": "token expired or invalid"})
			return
		}

		r = SetUser(r, user)
		next.ServeHTTP(w, r)
		return
	})
}

func (um *UserMiddleware) RequireUser(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user := GetUser(r)

		if user.IsAnonymous() {
			utils.WriteJson(w, http.StatusUnauthorized, utils.Envelop{"error": "you must be logged in to access this route"})
			return
		}

		next.ServeHTTP(w, r)
	})
}
