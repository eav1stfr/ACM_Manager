package router

import (
	"acmmanager/internal/api/handlers"
	"github.com/go-chi/chi/v5"
)

func registerRegularMembersRoutes(r chi.Router) {
	r.Route("/members", func(r chi.Router) {
		r.Get("/", handlers.GetMembersHandler) // get all members
		r.Post("/", handlers.CreateMemberHandler)
		r.Delete("/", handlers.DeleteMembersHandler)
		r.Patch("/", handlers.PatchMembersHandler)
		r.Delete("/delete_all", handlers.DeleteAllMembers)
	})
}
