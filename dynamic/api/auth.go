package api

import (
	"context"
	"net/http"
	"strings"
)

func WithAuthHandler(base http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		auth := r.Header.Get("authorization")
		if auth == "" {
			base.ServeHTTP(w, r)
			return
		}
		authFields := strings.Fields(auth)
		if len(authFields) != 2 {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
		if authFields[0] != "Bearer" {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}

		r.WithContext(WithAuth(r.Context(), authFields[1]))
		base.ServeHTTP(w, r)
	})
}

type auth_key struct{}

func WithAuth(ctx context.Context, auth string) context.Context {
	return context.WithValue(ctx, auth_key{}, auth)
}

func GetAuth(ctx context.Context) (string, bool) {
	auth, ok := ctx.Value(auth_key{}).(string)
	return auth, ok
}
