package products

import (
	"api/internal/domain/products"
	dbPack "api/internal/infrastructure"
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/shopspring/decimal"
)

type ProductCategoryDB struct {
	Id   int32  `db:"id"`
	Name string `db:"name"`
}

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

func (r *PostgresRepo) GetAll(ctx context.Context) ([]products.Product, error) {
	var allProducts []products.Product
	query := `
	SELECT p.id, p.product_name, p.article, p.quantity, p.price, p.location, p.reserved_quantity,
    pc.id, pc.category_name
    FROM products p
	LEFT JOIN product_categories pc ON p.category_id = pc.id`
	rows, err := r.db.QueryxContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to get products: %v", err)
	}
	defer rows.Close()
	
	for rows.Next() {
		var productDB ProductDB
		var productCategoryDB ProductCategoryDB
		err := rows.StructScan(&productCategoryDB)
		if err != nil {
			return nil, fmt.Errorf("failed to scan product category row: %v", err)
		}
		err = rows.StructScan(&productDB)
		if err != nil {
			return nil, fmt.Errorf("failed to scan product row: %v", err)
		}
		
		productCategory, err := products.NewProductCategory(productCategoryDB.Id, productCategoryDB.Name)
		if err != nil {
			return nil, fmt.Errorf("failed to create product category: %v", err)
		}
		product, err := products.NewProduct(productDB.Id, productDB.Name, productDB.Article,
			*productCategory, productDB.Quantity, productDB.Price, productDB.Location, productDB.ReservedQuantity)
		if err != nil {
			return nil, fmt.Errorf("failed to create product: %v", err)
		}
		allProducts = append(allProducts, *product)
	}
	
	return allProducts, nil
}
