package repository

import (
	"context"

	"goapi_admin_products/domain/products/application/v1/response"
	"goapi_admin_products/domain/products/domain/model"
)

// ProductRepository a
type ProductRepository interface {
	CreateProductHandler(ctx context.Context, product *model.Product) (*response.ProductCreateResponse, error)
}
