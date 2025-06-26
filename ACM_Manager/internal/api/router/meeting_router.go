package router

import (
	"acmmanager/internal/api/handlers"
	"github.com/go-chi/chi/v5"
)

func registerMeetingsRouter(r chi.Router) {
	r.Route("/meetings", func(r chi.Router) {
		r.Get("/", handlers.GetUpcomingMeeting)
	})
}
