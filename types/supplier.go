package types

type Supplier struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	AddresId    int    `json:"address_id"`
	PhoneNumber string `json:"phone_number"`
}

type SupplierDTO struct {
	Name        string `json:"name" db:"name"`
	Country     string `json:"country" db:"country"`
	City        string `json:"city" db:"city"`
	Street      string `json:"street" db:"street"`
	PhoneNumber string `json:"phone_number" db:"phone_number"`
}

// CREATE TABLE IF NOT EXISTS suppliers
// (
//     id SERIAL PRIMARY KEY,
//     name VARCHAR(100) NOT NULL,
//     address_id INT NOT NULL,
//     phone_number VARCHAR(50) NOT NULL,
//     FOREIGN KEY (address_id) REFERENCES addresses(id) ON DELETE CASCADE
// );
