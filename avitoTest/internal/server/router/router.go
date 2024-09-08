package router

import (
	h "avitoTest/avitoTest/internal/server/handlers"
	"avitoTest/avitoTest/internal/server/middlewares"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func GetRouter() {
	r := chi.NewRouter()

	r.Use(middlewares.GzipMiddleware)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Route("/api", func(r chi.Router) {
		r.Get("/ping", h.PingHandler)

		r.Route("/tenders", func(r chi.Router) {
			r.Get("/", nil)
			r.Post("/new", nil)
			r.Get("/my", nil)
			r.Route("/{tenderId}", func(r chi.Router) {
				r.Get("/status", nil)
				r.Put("/status", nil)
				r.Patch("/edit", nil)
				r.Put("/rollback/{version}", nil)
			})

		})

		r.Route("/bids", func(r chi.Router) {
			r.Post("/new", nil)
			r.Get("/my", nil)
			r.Get("/{tenderId}/list", nil)

			r.Route("/{bidId}", func(r chi.Router) {
				r.Get("/status", nil)
				r.Put("/status", nil)
				r.Patch("/edit", nil)
				r.Put("/submit_decision", nil)
				r.Put("/feedback", nil)
				r.Put("/rollback/{version}", nil)
				r.Get("/reviews", nil)
			})
		})

	})

}
