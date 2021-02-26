package v1

import (
	repoDomain "goapi_admin_products/domain/products/domain/repository"
)

// ProductRouter aaa
type ProductRouter struct {
	Repo repoDomain.ProductRepository
}
