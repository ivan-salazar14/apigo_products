package infrastructure

import (
	"net/http"

	v1Product "goapi_admin_products/domain/products/application/v1"
	"goapi_admin_products/infrastructure/database"

	"github.com/go-chi/chi"
)

func RoutesProducts(conn *database.Data) http.Handler {
	router := chi.NewRouter()

	products := v1Product.NewProductHandler(conn) //domain
	router.Mount("/products", routesProduct(products))
	return router
}

//Router user
func routesProduct(handler *v1Product.ProductRouter) http.Handler {
	router := chi.NewRouter()
	router.Post("/", handler.CreateProductHandler)
	return router
}
