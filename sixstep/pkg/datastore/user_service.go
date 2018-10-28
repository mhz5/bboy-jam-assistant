package datastore

import (
	"context"
	"fmt"

	"bboy-jam-assistant/sixstep/cmd/sixstep"

	"google.golang.org/appengine/datastore"
)

const (
	userKind = "User"
)

// TODO: Ask StackOverflow about wrapping datastore.
type UserService struct {}

var _ sixstep.UserService = &UserService{}

func NewUserService() *UserService {
	return &UserService{}
}

// User returns the User with the provided userId, or an error if the lookup fails.
func (s *UserService) User(ctx context.Context, userId int64) (*sixstep.User, error) {
	query := datastore.NewQuery(userKind).Filter("Id = ", userId)
	return runUserQuery(ctx, query)
}

// UserByName returns the User with the provided username, or an error if the lookup fails.
// This function should be invoked only when userId is not known (eg. when logging in).
func (s *UserService) UserByName(ctx context.Context, username string) (*sixstep.User, error) {
	query := datastore.NewQuery(userKind).Filter("Username = ", username)
	return runUserQuery(ctx, query)
}

// runUserQuery returns the user, or an error if the query fails.
func runUserQuery(ctx context.Context, query *datastore.Query) (*sixstep.User, error) {
	results := query.Run(ctx)
	user := &sixstep.User{}
	_, err := results.Next(user)
	if err == datastore.Done {
		return nil, fmt.Errorf("user not found")
	}
	if err != nil {
		return nil, err
	}

	return user, nil
}

// CreateUser creates a user with the provided username and passwordHash, and saves it to datastore.e
func (s *UserService) CreateUser(ctx context.Context, username, passwordHash string) (*sixstep.User, error) {
	id, _, err := datastore.AllocateIDs(ctx, userKind, nil, 1)
	if err != nil {
		return nil, err
	}

	u := &sixstep.User{id, username, passwordHash}
	key := datastore.NewKey(ctx, userKind, "", id, nil)
	_, err = datastore.Put(ctx, key, u)
	if err != nil {
		return nil, err
	}

	return u, err
}

func (s *UserService) DeleteUser(ctx context.Context, id int64) error {
	return fmt.Errorf("not implemented")
}
