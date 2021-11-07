package concurrency

import (
	"net/http"
)

type Semaphore interface {
	TryAcquire() error
	Release()
}

type Middleware struct {
	s Semaphore
}

func NewMiddlware(s Semaphore) *Middleware {
	return &Middleware{s: s}
}

func (m *Middleware) Handle(next http.Handler) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if err := m.s.TryAcquire(); err != nil {
			http.Error(w, http.StatusText(http.StatusTooManyRequests), http.StatusTooManyRequests)
			return
		}

		next.ServeHTTP(w, r)

		m.s.Release()
	})
}
