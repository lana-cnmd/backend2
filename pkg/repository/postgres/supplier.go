package postgres

import (
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/lana-cnmd/backend2/types"
)

type SupplierPostgresImpl struct {
	db *sqlx.DB
}

func NewSupplierPostgresImpl(db *sqlx.DB) *SupplierPostgresImpl {
	return &SupplierPostgresImpl{
		db: db,
	}
}

func (r *SupplierPostgresImpl) Create(supplier types.SupplierDTO) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var address_id int
	createAddressQuery := fmt.Sprintf("INSERT INTO %s (country, city, street) VALUES ($1, $2, $3) RETURNING id", addressesTable)
	addressRow := tx.QueryRow(createAddressQuery, supplier.Country, supplier.City, supplier.Street)
	if err := addressRow.Scan(&address_id); err != nil {
		tx.Rollback()
		return 0, err
	}

	var supplierId int
	createSupplierQuery := fmt.Sprintf("INSERT INTO %s (name, address_id, phone_number) VALUES ($1, $2, $3) RETURNING id", suppliersTable)
	row := tx.QueryRow(createSupplierQuery, supplier.Name, address_id, supplier.PhoneNumber)
	if err := row.Scan(&supplierId); err != nil {
		tx.Rollback()
		return 0, err
	}

	return supplierId, tx.Commit()
}

// select name, country, city, street, phone_number from suppliers join addresses on suppliers.address_id = addresses.id;
func (r *SupplierPostgresImpl) GetSupplierById(supplierId int) (types.SupplierDTO, error) {
	var supplier types.SupplierDTO
	rows := "name, country, city, street, phone_number"
	query := fmt.Sprintf("SELECT %s FROM %s JOIN %s ON %s.address_id = %s.id WHERE %s.id = $1", rows, suppliersTable, addressesTable, suppliersTable, addressesTable, suppliersTable)
	err := r.db.Get(&supplier, query, supplierId)

	return supplier, err
}

func (r *SupplierPostgresImpl) GetAllSuppliers() ([]types.SupplierDTO, error) {
	var suppliers []types.SupplierDTO
	rows := "name, country, city, street, phone_number"
	query := fmt.Sprintf("SELECT %s FROM %s JOIN %s ON %s.address_id = %s.id", rows, suppliersTable, addressesTable, suppliersTable, addressesTable)
	err := r.db.Select(&suppliers, query)

	return suppliers, err
}

// delete from addresses where id in (select address_id from suppliers where id = 2);
func (r *SupplierPostgresImpl) DeleteSupplierById(supplierId int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id IN (SELECT address_id FROM %s WHERE id = $1)", addressesTable, suppliersTable)
	_, err := r.db.Exec(query, supplierId)
	return err
}

func (r *SupplierPostgresImpl) UpdateSupplierAddress(supplierId int, newAddress types.UpdateAddressInput) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if newAddress.Country != nil {
		setValues = append(setValues, fmt.Sprintf("country=$%d", argId))
		args = append(args, *newAddress.Country)
		argId++
	}

	if newAddress.City != nil {
		setValues = append(setValues, fmt.Sprintf("city=$%d", argId))
		args = append(args, *newAddress.City)
		argId++
	}

	if newAddress.Street != nil {
		setValues = append(setValues, fmt.Sprintf("street=$%d", argId))
		args = append(args, *newAddress.Street)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf("UPDATE %s SET %s WHERE id=(SELECT address_id FROM %s WHERE %s.id=$%d)", addressesTable, setQuery, suppliersTable, suppliersTable, argId)
	args = append(args, supplierId)
	// slog.Info(query, args)

	_, err := r.db.Exec(query, args...)
	return err
}
