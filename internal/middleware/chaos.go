package middleware

import (
	"net/http"

	"ChaosApi/internal/chaos"
)

func Chaos(engine *chaos.Engine) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Não atrapalha preflight OPTIONS
			if r.Method == http.MethodOptions {
				next.ServeHTTP(w, r)
				return
			}

			err := engine.Apply(r.Context(), r.URL.Path, r.Method)
			if err != nil {
				if err == chaos.ErrInjectedFailure {
					http.Error(w, "chaos injected failure", http.StatusInternalServerError)
					return
				}

				http.Error(w, err.Error(), http.StatusRequestTimeout)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
