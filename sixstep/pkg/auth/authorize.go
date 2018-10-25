// authorize.go provides functions to guard APIs from unauthorized access.
package auth

import (
	"context"
)

const (
	authorizedUserIdKey = "authorizedUserId"
)

// NewAuthContext returns a new context in which the provided username is authorized.
func NewAuthContext(ctx context.Context, userId int64) context.Context {
	return context.WithValue(ctx, authorizedUserIdKey, userId)
}

// UserIsAuthorized returns true if the provided userId is authorized, else false.
func UserIsAuthorized(ctx context.Context, userId int64) bool {
	authorizedUsername := ctx.Value(authorizedUserIdKey)
	if authorizedUsername == nil {
		return false
	}
	if authorizedUsername.(int64) != userId {
		return false
	}
	return true
}
