package postgres

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/lana-cnmd/backend2/types"
)

type ProductPostgresimpl struct {
	db *sqlx.DB
}

func NewProductPostgresImpl(db *sqlx.DB) *ProductPostgresimpl {
	return &ProductPostgresimpl{db: db}
}

func (r *ProductPostgresimpl) Create(product types.CreateProductRequest) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var productId, count int

	queerySupp := fmt.Sprintf("SELECT COUNT(*) FROM %s WHERE id = $1", suppliersTable)
	err = r.db.QueryRow(queerySupp, product.SupplierId).Scan(&count)
	if err != nil {
		return 0, fmt.Errorf("failed to check supplier existence: %w", err)
	}
	if count == 0 {
		return 0, fmt.Errorf("supplier with id %d does not exist", product.SupplierId)
	}

	query := fmt.Sprintf("INSERT INTO %s (name, category, price, available_stock, last_update_date, supplier_id, image_id) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id", productsTable)
	// log.Println(query)
	row := tx.QueryRow(query, product.Name, product.Category, product.Price, product.AvailableStock, time.Now(), product.SupplierId, product.ImageId)
	if err := row.Scan(&productId); err != nil {
		tx.Rollback()
		return 0, err
	}
	return productId, tx.Commit()
}

func (r *ProductPostgresimpl) GetProductById(productId int) (types.GetProductResponce, error) {
	var product types.GetProductResponce
	rows := "name, category, price, available_stock, last_update_date, supplier_id, image_id"
	query := fmt.Sprintf("SELECT %s FROM %s WHERE id = $1", rows, productsTable)
	err := r.db.Get(&product, query, productId)

	return product, err
}

func (r *ProductPostgresimpl) GetAllProducts() ([]types.GetProductResponce, error) {
	var products []types.GetProductResponce
	rows := "name, category, price, available_stock, last_update_date, supplier_id, image_id"
	query := fmt.Sprintf("SELECT %s FROM %s", rows, productsTable)
	err := r.db.Select(&products, query)

	return products, err
}

func (r *ProductPostgresimpl) DeleteProductById(productId int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id = $1", productsTable)
	_, err := r.db.Exec(query, productId)
	return err
}

func (r *ProductPostgresimpl) DecreaseProductAmount(productId int, decreaseAmount int) error {

	var currentStock int
	query := fmt.Sprintf("SELECT available_stock FROM %s WHERE id = $1 FOR UPDATE", productsTable)
	err := r.db.QueryRow(query, productId).Scan(&currentStock)
	if err != nil {
		if err == sql.ErrNoRows {
			return fmt.Errorf("product with id %d not found", productId)
		}
		return fmt.Errorf("failed to get current stock: %w", err)
	}

	if currentStock < decreaseAmount {
		return fmt.Errorf("not enough stock: current stock is %d", currentStock)
	}

	updateQuery := fmt.Sprintf("UPDATE %s SET available_stock = available_stock - $1 WHERE id = $2", productsTable)

	result, err := r.db.Exec(updateQuery, decreaseAmount, productId)
	if err != nil {
		return fmt.Errorf("failed to decrease product amount: %w", err)
	}

	// Проверяем, была ли обновлена хотя бы одна строка
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to check rows affected: %w", err)
	}
	if rowsAffected == 0 {
		return fmt.Errorf("product with id %d not found", productId)
	}

	return nil
}
