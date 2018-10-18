package main

import (
	"bboy-jam-assistant/sixstep/pkg/http"
	"google.golang.org/appengine"
)

func init() {
	s := http.NewServer()
	s.Serve()
	appengine.Main()
}
