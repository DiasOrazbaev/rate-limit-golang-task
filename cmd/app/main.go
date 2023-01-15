package main

import (
	"fmt"
	"log"
	"rate-limit-golang-task/internal/app/config"
	"rate-limit-golang-task/internal/app/router"
	"rate-limit-golang-task/pkg/middleware"
	"time"
)

func main() {
	cfg := config.GetConfig()

	d, err := time.ParseDuration(fmt.Sprintf("%ds", cfg.BlockDuration))
	if err != nil {
		log.Fatalln(err.Error())
	}

	limiter := middleware.NewLimiter(
		cfg.Limit,
		cfg.Limit,
		d,
	)
	r := router.NewRouter(limiter)

	if err := r.Run(); err != nil {
		log.Fatalln(err.Error())
	}
}
