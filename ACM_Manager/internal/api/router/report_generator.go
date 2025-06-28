package router

import "github.com/go-chi/chi/v5"

func registerReportRoute(r chi.Router) {
	r.Route("/report", func(r chi.Router) {
		//r.Get("/", handlers)
	})
}
