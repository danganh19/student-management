package router

import (
	"github.com/go-chi/chi"
	"github.com/swaggo/http-swagger"
	_ "ims-be/docs"
	"net/http"
)

const basePath string = "/api/tda/student-management"

func Router() http.Handler {
	r := chi.NewRouter()

	r.Get(basePath+"/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:3000/api/tda/student-management/swagger/doc.json"),
	))

	r.Route(basePath, func(r chi.Router) {
		r.Group(func(r chi.Router) {

		})

		r.Group(func(r chi.Router) {

		})

		r.Group(func(r chi.Router) {

		})

		r.Group(func(r chi.Router) {

		})
	})
	return r
}
