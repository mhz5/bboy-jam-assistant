package http

import (
	"context"
	"fmt"
	"net/http"

	"bboy-jam-assistant/sixstep/pkg/sessions"
)

// authorize asserts that the username cookie matches the username encrypted in the session.
// Else send 401 - Unauthorized http error.
func authorize(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		s, err := sessions.UserSession(r)
		c, err := r.Cookie(usernameKey)
		if err != nil {
			http.Error(w, fmt.Sprintf("%v: username cookie not sent", err), http.StatusBadRequest)
			return
		}
		username := c.Value
		expectedUsername, err := s.Username()
		if err != nil {
			http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
			return
		}
		if expectedUsername != username {
			http.Error(w, "user not authorized to view resource", http.StatusUnauthorized)
		}
		if err != nil {
			http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
			return
		}

		userId, err := s.UserId()
		if err != nil {
			http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
			return
		}

		// TODO: extract this into separate function.
		// TODO: Need to use authorized user in handlers that need authorization.
		ctx := context.WithValue(r.Context(), "authorizedUserId", userId)
		next.ServeHTTP(w, r.WithContext(ctx))
	}
}
