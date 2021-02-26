package v1

import (
	"fmt"
	"goapi_admin_products/domain/products/domain/model"
	repoDomain "goapi_admin_products/domain/products/domain/repository"
	"goapi_admin_products/domain/products/infrastructure/persistence"
	"goapi_admin_products/infrastructure/database"
	"goapi_admin_products/infrastructure/middleware"
	"net/http"
)

// ProductRouter aaa
type ProductRouter struct {
	Repo repoDomain.ProductRepository
}

// NewProductHandler  a
func NewProductHandler(db *database.Data) *ProductRouter {
	return &ProductRouter{
		Repo: persistence.NewProductRepository(db),
	}
}

// CreateProductHandler aa
func (prod *ProductRouter) CreateProductHandler(w http.ResponseWriter, r *http.Request) {
	var product model.Product

	ctx := r.Context()

	result, err := prod.Repo.CreateProductHandler(ctx, &product)
	if err != nil {

		_ = middleware.HTTPError(w, r, http.StatusConflict, "conflict", err.Error())
		return
	}

	w.Header().Add("Location", fmt.Sprintf("%s%s", r.URL.String(), result))
	_ = middleware.JSON(w, r, http.StatusCreated, result)
}
