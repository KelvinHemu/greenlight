package main

import (
	"context"
	"net/http"

	"greenlight.chriss875.net/internal/data"
)

type contextKey string

// userContextKey is a context key used to store and retrieve user information from a context object.
const userContextKey = contextKey("user")

// contextSetUser returns a new request with the provided user added to the request's context.
func (app *application) contextSetUser(r *http.Request, user *data.User) *http.Request {
	ctx := context.WithValue(r.Context(), userContextKey, user)
	return r.WithContext(ctx)
}

func (app *application) contextGetUser(r *http.Request) *data.User {
	user, ok := r.Context().Value(userContextKey).(*data.User)
	if !ok {
		panic("missing user value in request context")
	}
	return user
}
