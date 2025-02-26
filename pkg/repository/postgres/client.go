package postgres

import (
	"fmt"
	"strings"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/lana-cnmd/backend2/types"
)

type ClientPostgresImpl struct {
	db *sqlx.DB
}

func NewClientPostgresImpl(db *sqlx.DB) *ClientPostgresImpl {
	return &ClientPostgresImpl{db: db}
}

func (r *ClientPostgresImpl) Create(client types.CreateClientRequest) (int, error) {

	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}
	// log.Println((client))
	var addressId int
	createAddressQuery := fmt.Sprintf("INSERT INTO %s (country, city, street) VALUES ($1, $2, $3) RETURNING id", addressesTable)
	// log.Println(createAddressQuery)
	addressRow := tx.QueryRow(createAddressQuery, client.Country, client.City, client.Street)
	if err := addressRow.Scan(&addressId); err != nil {
		tx.Rollback()
		return 0, err
	}

	var clientId int
	createClientQuery := fmt.Sprintf("INSERT INTO %s (client_name, client_surname, birthday, gender, registration_date, address_id) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id", clientsTable)
	// log.Println(createClientQuery)
	row := tx.QueryRow(createClientQuery, client.ClientName, client.ClientSurname, client.Birthday.Time, client.Gender, time.Now(), addressId)
	if err := row.Scan(&clientId); err != nil {
		tx.Rollback()
		return 0, err
	}
	return clientId, tx.Commit()
}

// select client_name, client_surname, birthday, gender, registration_date, country, city, street from clients join addresses on clients.address_id = addresses.id where clients.client_name = 'Slava' and clients.client_surname = 'Stepanov' limit 1;
func (r *ClientPostgresImpl) SearchClientByName(firstName, lastName string) (types.SearchClientResponse, error) {
	var client types.SearchClientResponse
	rows := "client_name, client_surname, birthday, gender, registration_date, country, city, street"
	query := fmt.Sprintf("SELECT %s FROM %s JOIN %s ON %s.address_id = %s.id WHERE %s.client_name = $1 AND %s.client_surname = $2 LIMIT 1", rows, clientsTable, addressesTable, clientsTable, addressesTable, clientsTable, clientsTable)
	err := r.db.Get(&client, query, firstName, lastName)

	return client, err
}

// DELETE FROM addresses WHERE id IN (SELECT address_id FROM clients WHERE id = 8);
func (r *ClientPostgresImpl) DeleteClient(clientId int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id IN (SELECT address_id FROM %s WHERE id = $1) ", addressesTable, clientsTable)
	_, err := r.db.Exec(query, clientId)
	return err
}

// select client_name, client_surname, birthday, gender, registration_date, country, city, street from clients join addresses on clients.address_id = addresses.id
func (r *ClientPostgresImpl) GetAllClients(limit, offset int) ([]types.SearchClientResponse, error) {
	var clients []types.SearchClientResponse
	rows := "client_name, client_surname, birthday, gender, registration_date, country, city, street"
	query := fmt.Sprintf("SELECT %s FROM %s JOIN %s ON %s.address_id = %s.id limit $1 offset $2", rows, clientsTable, addressesTable, clientsTable, addressesTable)
	err := r.db.Select(&clients, query, limit, offset)
	return clients, err
}

func (r *ClientPostgresImpl) UpdateClientAddress(clientId int, newAddress types.UpdateAddressInput) error {
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

	query := fmt.Sprintf("UPDATE %s SET %s WHERE id=(SELECT address_id FROM %s WHERE %s.id=$%d)", addressesTable, setQuery, clientsTable, clientsTable, argId)
	args = append(args, clientId)
	// slog.Info(query, args)

	_, err := r.db.Exec(query, args...)
	return err
}
