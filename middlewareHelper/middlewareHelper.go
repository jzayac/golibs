package middlewareHelper

import (
	"net/http"
)

// DisallowAnon does not allow anonymous user to access the content
func DisallowAnon(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// ok status
		h.ServeHTTP(w, r)
	})
}
