// Package recover provides a middleware handler that recovers from panics, outputs the error to the log, and writes a
// StatusCode of 500.
package recover

import (
	"log"
	"net/http"
)

// Handler returns a Recover middleware.
func Handler(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("panic: %+v", err)
				http.Error(w, "500 - Internal Server Error", http.StatusInternalServerError)
			}
		}()

		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}
