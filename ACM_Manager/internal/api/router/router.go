package router

import (
	"acmmanager/internal/api/handlers"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func Router() http.Handler {
	r := chi.NewRouter()

	r.Get("/", handlers.HelloHandler)

	registerRegularMembersRoutes(r)
	registerTaskRoutes(r)

	return r

}
