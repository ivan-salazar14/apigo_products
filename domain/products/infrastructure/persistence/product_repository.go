package persistence

import (
	"context"
	"database/sql"
	"goapi_admin_products/domain/products/application/v1/response"
	"goapi_admin_products/domain/products/domain/model"
	repoDomain "goapi_admin_products/domain/products/domain/repository"
	"goapi_admin_products/infrastructure/database"
)

type sqlProductRepo struct {
	Conn *database.Data
}

// NewProductRepository aa
func NewProductRepository(Conn *database.Data) repoDomain.ProductRepository {
	return &sqlProductRepo{
		Conn: Conn,
	}
}

// CreateProductHandler aa
func (sr *sqlProductRepo) CreateProductHandler(ctx context.Context, product *model.Product) (*response.ProductCreateResponse, error) {

	stmt, err := sr.Conn.DB.PrepareContext(ctx, insertProduct)
	if err != nil {
		return &response.ProductCreateResponse{}, err
	}
	defer stmt.Close()

	row := stmt.QueryRowContext(ctx, &product.ProductoID, &product.ProductoNombre, &product.ProductoCantidad, &product.ProductoUsercreacion, &product.ProductoUserModificacion)

	var idResult string
	err = row.Scan(&idResult)
	if err != sql.ErrNoRows {
		return &response.ProductCreateResponse{}, err
	}

	ProductResponse := response.ProductCreateResponse{
		Message: "product created",
	}

	return &ProductResponse, nil

}
