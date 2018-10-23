package http

import (
	"context"
	"fmt"
	"net/http"

	"bboy-jam-assistant/sixstep/pkg/sessions"
	"google.golang.org/appengine/log"
)

func (rtr *Router) authorize(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		s, err := sessions.UserSession(r)
		if err != nil {
			log.Errorf(r.Context(), "%v", err)
			http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
			return
		}
		ctx := context.WithValue(r.Context(), "userId", s.UserId())
		next.ServeHTTP(w, r.WithContext(ctx))
	}
}
