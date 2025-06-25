package router

import (
	"acmmanager/internal/api/handlers"
	"github.com/go-chi/chi/v5"
)

func registerTaskRoutes(r chi.Router) {
	r.Route("/tasks", func(r chi.Router) {
		r.Get("/", handlers.GetTasks)
		r.Post("/", handlers.CreateTasks)
		r.Patch("/", handlers.UpdateTask)
		r.Delete("/", handlers.DeleteTasks)
	})
}
