package yandex_lavka

import (
	"context"
	"golang.org/x/time/rate"
	"net/http"
	"time"
)

type Server struct {
	httpServer *http.Server
	limiter    *rate.Limiter
}

func (s *Server) Run(port string, handler http.Handler) error {
	s.limiter = rate.NewLimiter(rate.Every(time.Second/10), 1) // Ограничение до 10 запросов в секунду
	s.httpServer = &http.Server{
		Addr:           ":" + port,
		Handler:        s.limitMiddleware(handler),
		MaxHeaderBytes: 1 << 20,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}

	return s.httpServer.ListenAndServe()
}

func (s *Server) limitMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !s.limiter.Allow() {
			http.Error(w, http.StatusText(http.StatusTooManyRequests), http.StatusTooManyRequests)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
