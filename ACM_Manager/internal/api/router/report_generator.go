package router

import (
	"acmmanager/internal/api/handlers"
	"github.com/go-chi/chi/v5"
)

func registerReportRoute(r chi.Router) {
	r.Route("/report", func(r chi.Router) {
		r.Get("/member", handlers.GenerateReportForMember)
		r.Get("/department", handlers.GenerateReportForDepartment)
		r.Get("/", handlers.PDFHandler)
	})
}
