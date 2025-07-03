package router

import (
	"acmmanager/internal/api/handlers"
	"github.com/go-chi/chi/v5"
)

func registerTaskRoutes(r chi.Router) {
	r.Route("/tasks", func(r chi.Router) {
		r.Get("/", handlers.GetTasks)
		r.Post("/", handlers.CreateTask)
		r.Patch("/finish", handlers.MarkTaskAsDone)
		r.Patch("/assign", handlers.AssignTask)
		r.Delete("/", handlers.DeleteTasks)
	})
}
