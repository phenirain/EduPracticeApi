package products

import (
	"api/internal/domain/products"
	domProduct "api/internal/domain/products"
	"context"
	"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/shopspring/decimal"
	"reflect"
	"strings"
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

func (p *ProductDB) TableName() string {
	return "products"
}

func (p *ProductDB) ID() int32 {
	return p.Id
}

type PostgresRepo struct {
	db *sqlx.DB
}

func NewPostgresRepo(db *sqlx.DB) *PostgresRepo {
	return &PostgresRepo{
		db: db,
	}
}

func (r *PostgresRepo) GetAllCategories(ctx context.Context) ([]*products.ProductCategory, error) {
	categories := make([]*products.ProductCategory, 0, 25)
	rows, err := r.db.QueryxContext(ctx, "SELECT id as category_id, category_name FROM product_categories")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var category ProductCategory
		err := rows.StructScan(&category)
		if err != nil {
			return nil, fmt.Errorf("failed to scan product category row: %v", err)
		}
		productCategory, err := products.NewProductCategory(category.Id, category.Name)
		if err != nil {
			return nil, fmt.Errorf("failed to create product category: %v", err)
		}
		categories = append(categories, productCategory)
	}
	return categories, nil
}

func (r *PostgresRepo) GetById(ctx context.Context, id int32) (*products.Product, error) {
	var productView ProductView
	err := r.db.GetContext(ctx, &productView.View, productView.Query+"WHERE p.id = $1", id)
	if err != nil {
		return nil, err
	}

	productCategory, err := products.NewProductCategory(productView.View.CategoryId,
		productView.View.CategoryName)
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

func (r *PostgresRepo) GetAll(ctx context.Context) ([]**products.Product, error) {
	var allProducts []**products.Product
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

		productCategory, err := products.NewProductCategory(productView.View.CategoryId,
			productView.View.CategoryName)
		if err != nil {
			return nil, fmt.Errorf("failed to create product category: %v", err)
		}
		product, err := products.NewProduct(productView.View.Id, productView.View.Name, productView.View.Article,
			*productCategory, productView.View.Quantity, productView.View.Price, productView.View.Location,
			productView.View.ReservedQuantity)
		if err != nil {
			return nil, fmt.Errorf("failed to create product: %v", err)
		}
		allProducts = append(allProducts, &product)
	}

	return allProducts, nil
}
func (r *PostgresRepo) Create(ctx context.Context, model *domProduct.Product) (*domProduct.Product, error) {
	productDB := &ProductDB{
		Name:             model.Name,
		Article:          model.Article,
		Category:         model.Category.Id,
		Quantity:         model.Quantity,
		Price:            model.Price,
		Location:         model.Location,
		ReservedQuantity: model.ReservedQuantity,
	}

	val := reflect.ValueOf(*productDB)
	typ := reflect.TypeOf(*productDB)
	fields := make([]string, 0, typ.NumField()-1)
	args := make([]interface{}, 0, typ.NumField()-1)
	argsIds := make([]string, 0, typ.NumField()-1)

	for i := 0; i < typ.NumField(); i++ {
		if typ.Field(i).Name == "Id" {
			continue
		}
		fields = append(fields, typ.Field(i).Name)
		argsIds = append(argsIds, fmt.Sprintf("$%d", len(args)+1))
		args = append(args, val.Field(i).Interface())
	}
	query := fmt.Sprintf(`INSERT INTO %s (%s) VALUES (%s)`, productDB.TableName(), strings.Join(fields, ", "+
		""), strings.Join(argsIds, ", "))

	var id int32
	err := r.db.QueryRowxContext(ctx, query, args...).Scan(&id)
	if err != nil {
		// must return model, because i cannot return nil due all interfaces must can operate with pointer
		//instead copy of struct
		return model, fmt.Errorf("failed to insert to %s: %v", productDB.TableName(), err)
	}
	model.SetId(id)
	return model, nil
}

func (r *PostgresRepo) ExistsById(ctx context.Context, id int32) (bool, error) {
	productDB := &ProductDB{}
	query := fmt.Sprintf(`SELECT 1 FROM %s WHERE id = $1`, productDB.TableName())
	var result int32
	err := r.db.QueryRowxContext(ctx, query, id).Scan(&result)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, fmt.Errorf("failed to check existence: %v", err)
	}
	return true, nil
}

func (r *PostgresRepo) Update(ctx context.Context, model *domProduct.Product) error {
	productDB := &ProductDB{
		Id:               model.Id,
		Name:             model.Name,
		Article:          model.Article,
		Category:         model.Category.Id,
		Quantity:         model.Quantity,
		Price:            model.Price,
		Location:         model.Location,
		ReservedQuantity: model.ReservedQuantity,
	}

	val := reflect.ValueOf(*productDB)
	typ := reflect.TypeOf(*productDB)
	fields := make([]string, 0, typ.NumField()-1)
	args := make([]interface{}, 0, typ.NumField()-1)

	for i := 0; i < typ.NumField(); i++ {
		if typ.Field(i).Name == "Id" {
			continue
		}
		fields = append(fields, fmt.Sprintf("%s = $%d", typ.Field(i).Name, len(args)+1))
		args = append(args, val.Field(i).Interface())
	}

	query := fmt.Sprintf(`UPDATE %s SET %s WHERE id = $%d`, productDB.TableName(), strings.Join(fields, ", "), productDB.ID())

	_, err := r.db.ExecContext(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("failed to update %s with id = %d: %v", productDB.TableName(), productDB.ID(), err)
	}
	return nil
}

func (r *PostgresRepo) Delete(ctx context.Context, id int32) error {
	productDB := &ProductDB{}
	query := fmt.Sprintf(`DELETE FROM %s WHERE id = $1`, productDB.TableName())
	_, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return fmt.Errorf("failed to delete %s with id = %d: %v", productDB.TableName(), id, err)
	}
	return nil
}
