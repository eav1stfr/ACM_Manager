package router

import (
	"acmmanager/internal/api/handlers"
	"github.com/go-chi/chi/v5"
)

func registerAttendanceRoutes(r chi.Router) {
	r.Route("/attendance", func(r chi.Router) {
		r.Patch("/attend", handlers.MarkAttendance)
		r.Patch("/miss", handlers.GetReasonOfAbsence)
	})
}
