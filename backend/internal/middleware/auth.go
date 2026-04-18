package middleware

import (
	"context"
	"net/http"

	"github.com/alexedwards/scs/v2"
	"github.com/google/uuid"

	"github.com/Karenmiano/vibe/pkg/utilities"
)

type contextKey string
const UserIDKey contextKey = "userId"

type AuthMiddleware struct {
	sessionManager *scs.SessionManager
}

func NewAuthMiddleware(sessionManager *scs.SessionManager) *AuthMiddleware {
	return &AuthMiddleware{
		sessionManager: sessionManager,
	}
}


func (m *AuthMiddleware) Authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userId, ok := m.sessionManager.Get(r.Context(), "userId").(uuid.UUID)
		if !ok {
			utilities.WriteJSON(w, http.StatusUnauthorized, map[string]string{"message": "authentication required"})
			return
		}

		ctx := context.WithValue(r.Context(), UserIDKey, userId)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}