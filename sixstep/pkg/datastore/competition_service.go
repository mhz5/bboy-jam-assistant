package datastore

import (
	"context"
	"fmt"

	"bboy-jam-assistant/sixstep/cmd/sixstep"

	"google.golang.org/appengine/datastore"
)

const (
	competitionKind = "Competition"
)

// TODO: Ask StackOverflow about wrapping datastore.
type CompetitionService struct {}

var _ sixstep.CompetitionService = &CompetitionService{}

func NewCompetitionService() *CompetitionService {
	return &CompetitionService{}
}

// Competition returns the Competition with the provided compId, or an error if the lookup fails.
func (s *CompetitionService) Competition(ctx context.Context, compId int64) (*sixstep.Competition, error) {
	query := datastore.NewQuery(competitionKind).Filter("Id = ", compId)
	return runCompetitionQuery(ctx, query)
}

// runCompetitionQuery returns the competition, or an error if the query fails.
func runCompetitionQuery(ctx context.Context, query *datastore.Query) (*sixstep.Competition, error) {
	results := query.Run(ctx)
	c := &sixstep.Competition{}
	_, err := results.Next(c)
	if err == datastore.Done {
		return nil, fmt.Errorf("competition not found")
	}
	if err != nil {
		return nil, err
	}

	return c, nil
}

// CreateCompetition creates a competition with the provided name, and saves it to datastore.
func (s *CompetitionService) CreateCompetition(ctx context.Context, name string) (*sixstep.Competition, error) {
	id, _, err := datastore.AllocateIDs(ctx, competitionKind, nil, 1)
	if err != nil {
		return nil, err
	}

	c := &sixstep.Competition{id, name}
	key := datastore.NewKey(ctx, competitionKind, "", id, nil)
	_, err = datastore.Put(ctx, key, c)
	if err != nil {
		return nil, err
	}

	return c, err
}
