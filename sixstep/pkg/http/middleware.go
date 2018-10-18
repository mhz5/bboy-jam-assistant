package http

import (
	"net/http"

	"github.com/rs/cors"
)

func (r *Router) corsRouter(h http.Handler) http.Handler {
	c := cors.New(cors.Options{
		AllowedOrigins: []string{allowedOrigin},
		AllowCredentials: true,
		AllowedMethods: []string{"GET", "POST"},
		// Enable Debugging for testing, consider disabling in production
		Debug: true,
	})

	return c.Handler(h)
}

func (r *Router) authorizeRequest(h http.Handler) http.Handler {
	return nil
}
