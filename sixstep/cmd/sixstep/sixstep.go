package sixstep

import (
	"context"
)

// Domain types

type User struct {
	Id           int64  `json:"id"`
	Username     string `json:"username"`
	PasswordHash string `json:"-"`
}

type UserService interface {
	User(ctx context.Context, userId int64) (*User, error)
	UserByName(ctx context.Context, username string) (*User, error)
	CreateUser(ctx context.Context, username, passwordHash string) (*User, error)
	DeleteUser(ctx context.Context, userId int64) error
}

type Competition struct {
	Id   int64
	Name string
}

type CompetitionService interface {
	Competition(ctx context.Context, compId int64) (*Competition, error)
	CreateCompetition(ctx context.Context, name string) (*Competition, error)
}

type AuthService interface {
	Authenticate(ctx context.Context, username string, password string) (*User, error)
	Authorize()
}

type Router interface {
	Handle()
}

type Server interface {
	Serve()
}
