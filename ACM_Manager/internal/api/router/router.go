package router

import (
	"acmmanager/internal/api/handlers"
	"github.com/go-chi/chi/v5"
)

func Router() {
	r := chi.NewRouter()

	r.Get("/", handlers.HelloHandler)

	registerRegularMembersRoutes(r)

}
