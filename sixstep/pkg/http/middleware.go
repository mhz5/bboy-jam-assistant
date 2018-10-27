package http

import (
	"fmt"
	"net/http"
	"os"
	"sync"

	"bboy-jam-assistant/sixstep/pkg/auth"
	"bboy-jam-assistant/sixstep/pkg/sessions"
	"github.com/rs/cors"
	"google.golang.org/appengine"
)

var (
	once = sync.Once{}
	allowedOrigin = os.Getenv("ALLOWED_ORIGIN")
)

// TODO: Is there a better way than making middleware like this?
// appengineCtxRouter returns a handler that sets an appengine context as the request context.
func appengineCtxRouter(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := appengine.NewContext(r)
		r = r.WithContext(ctx)
		h.ServeHTTP(w, r)
	})
}

// corsRouter returns a handler that sets CORS headers on all incoming requests.
func corsRouter(h http.Handler) http.Handler {
	c := cors.New(cors.Options{
		AllowedOrigins: []string{allowedOrigin},
		AllowCredentials: true,
		AllowedMethods: []string{"GET", "POST"},
		// Enable Debugging for testing, consider disabling in production
		Debug: true,
	})

	return c.Handler(h)
}

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

		ctx := auth.NewAuthContext(r.Context(), userId)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	}
}

// warmup performs initializations once.
func warmup(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		once.Do(func() {
			sessions.InitStore(r)
		})
		h.ServeHTTP(w, r)
	})
}