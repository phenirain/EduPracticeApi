package products

import (
	"api/internal/domain/products"
	dbPack "api/internal/infrastructure"
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/shopspring/decimal"
)

type ProductDB struct {
	Id               int32           `db:"id"`
	Name             string          `db:"name"`
	Article          string          `db:"article"`
	Category         int32           `db:"category"`
	Quantity         int32           `db:"quantity"`
	Price            decimal.Decimal `db:"price"`
	Location         string          `db:"location"`
	ReservedQuantity int32           `db:"reserved_quantity"`
}

func (p *ProductDB) FromModelToDB(product *products.Product) {
	p.Id = product.Id
	p.Name = product.Name
	p.Article = product.Article
	p.Category = product.Category.Id
	p.Quantity = product.Quantity
	p.Price = product.Price
	p.Location = product.Location
	p.ReservedQuantity = product.ReservedQuantity
}

func (p *ProductDB) TableName() string {
	return "products"
}

func (p *ProductDB) ID() int32 {
	return p.Id
}

type PostgresRepo struct {
	*dbPack.Repository[*ProductDB, *products.Product]
	db *sqlx.DB
}

func NewPostgresRepo(db *sqlx.DB) *PostgresRepo {
	baseRepo := dbPack.NewRepository[*ProductDB, *products.Product](db)
	return &PostgresRepo{
		Repository: baseRepo,
		db:         db,
	}
}

func (r *PostgresRepo) GetById(ctx context.Context, id int32) (*products.Product, error) {
	var productView ProductView
	err := r.db.GetContext(ctx, &productView.View, productView.Query+"WHERE p.id = $1", id)
	if err != nil {
		return nil, err
	}

	productCategory, err := products.NewProductCategory(productView.View.Category.Id,
		productView.View.Category.Name)
	if err != nil {
		return nil, fmt.Errorf("failed to create product category: %v", err)
	}
	product, err := products.NewProduct(productView.View.Id, productView.View.Name, productView.View.Article,
		*productCategory, productView.View.Quantity, productView.View.Price, productView.View.Location,
		productView.View.ReservedQuantity)
	if err != nil {
		return nil, fmt.Errorf("failed to create product: %v", err)
	}
	return product, nil
}

func (r *PostgresRepo) GetAll(ctx context.Context) ([]*products.Product, error) {
	var allProducts []*products.Product
	productView := MustNewProductView()
	rows, err := r.db.QueryxContext(ctx, productView.Query)
	if err != nil {
		return nil, fmt.Errorf("failed to get products: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.StructScan(&productView.View)
		if err != nil {
			return nil, fmt.Errorf("failed to scan product row: %v", err)
		}

		productCategory, err := products.NewProductCategory(productView.View.Category.Id,
			productView.View.Category.Name)
		if err != nil {
			return nil, fmt.Errorf("failed to create product category: %v", err)
		}
		product, err := products.NewProduct(productView.View.Id, productView.View.Name, productView.View.Article,
			*productCategory, productView.View.Quantity, productView.View.Price, productView.View.Location,
			productView.View.ReservedQuantity)
		if err != nil {
			return nil, fmt.Errorf("failed to create product: %v", err)
		}
		allProducts = append(allProducts, product)
	}

	return allProducts, nil
}
