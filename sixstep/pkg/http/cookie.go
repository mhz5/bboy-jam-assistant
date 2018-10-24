package http

import (
	"fmt"
	"net/http"
)

func setCookie(w http.ResponseWriter, name, value string) {
	w.Header().Add("Set-Cookie", fmt.Sprintf("%s=%s", name, value))
}
