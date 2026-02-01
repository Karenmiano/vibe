package middleware

import (
	"context"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/sessions"
)

type contextKey string
const UserIDKey contextKey = "userId"

type AuthMiddleware struct {
	sessionStore sessions.Store
}

func NewAuthMiddleware(sessionStore sessions.Store) *AuthMiddleware {
	return &AuthMiddleware{
		sessionStore: sessionStore,
	}
}


func (m *AuthMiddleware) Authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		session, _ := m.sessionStore.Get(r, "vibe")
		userId, ok := session.Values["userId"].(uuid.UUID)
		if !ok {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}

		ctx := context.WithValue(r.Context(), UserIDKey, userId)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}