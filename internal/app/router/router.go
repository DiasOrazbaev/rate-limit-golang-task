package router

import (
	"net/http"
	"rate-limit-golang-task/internal/app/handlers"
	"rate-limit-golang-task/pkg/middleware"
)

type Router struct {
	limiter *middleware.Limiter
}

func NewRouter(limiter *middleware.Limiter) *Router {
	return &Router{limiter: limiter}
}

func (s *Router) Run() error {
	Tmux := &http.ServeMux{}
	Tmux.Handle("/test", s.limiter.RateLimit(handlers.Home))

	srv := http.Server{
		Addr:    ":8082",
		Handler: Tmux,
	}

	return srv.ListenAndServe()
}
