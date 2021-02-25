package infrastructure

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/ivan-salazar14/apigo_products/infrastructure/database"
)

func RoutesProducts(conn *database.Data) http.Handler {
	router := chi.NewRouter()

	lr := handler.NewProductsHandler(conn) //domain

	router.Get("/")
	//router.Get("/{id}", handler.GetOneHandler)
	//	router.Post("/", handler.CreateHandler)

	//router.Put("/{id}", handler.UpdateHandler)
	//router.Delete("/{id}", handler.DeleteHandler)
	return router
}
