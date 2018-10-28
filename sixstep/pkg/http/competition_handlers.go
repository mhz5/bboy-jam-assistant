// competition_handlers.go provides handlers for CRUD operations on competitions.
package http

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"google.golang.org/appengine/log"
	"net/http"
	"strconv"
)

const (
	compNameKey = "name"
	compIdParam = "compId"
)

func (rtr *Router) handleCreateCompetition(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	name := r.PostFormValue(compNameKey)

	c, err := rtr.competitionService.CreateCompetition(ctx, name)
	if err != nil {
		log.Errorf(ctx, "%v", err)
		http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
		return
	}

	compJson, err := json.Marshal(c)
	log.Infof(ctx, string(compJson))
	w.Write(compJson)
}

func (rtr *Router) handleGetCompetition(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	vars := mux.Vars(r)
	compId, err := strconv.ParseInt(vars[compIdParam], 10, 64)
	if err != nil {
		log.Errorf(ctx, "%v", err)
		http.Error(w, fmt.Sprintf("non-integer competitionId in URL: %v", err), http.StatusBadRequest)
		return
	}

	c, err := rtr.competitionService.Competition(ctx, compId)
	if err != nil {
		log.Errorf(ctx, "%v", err)
		http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
		return
	}

	compJson, err := json.Marshal(c)
	w.Write(compJson)
}
