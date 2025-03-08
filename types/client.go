package types

import (
	"fmt"
	"time"
)

type CustomTime struct {
	Time time.Time
}

// UnmarshalJSON для десериализации из JSON
func (ct *CustomTime) UnmarshalJSON(data []byte) error {
	// Удаляем surrounding кавычки
	dateStr := string(data[1 : len(data)-1])

	// Проверяем, является ли строка пустой
	if dateStr == "" {
		ct.Time = time.Time{}
		return nil
	}

	// Парсим дату в формате "YYYY-MM-DD"
	parsedTime, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		return fmt.Errorf("failed to parse date %q: %w", dateStr, err)
	}
	ct.Time = parsedTime
	return nil
}

type Client struct {
	ID               int        `json:"id"`
	ClientName       string     `json:"client_name"`
	ClientSurname    string     `json:"client_surname"`
	Birthday         CustomTime `json:"birthday"`
	Gender           string     `json:"gender"`
	RegistrationDate time.Time  `json:"registration_date"`
	AddressID        int        `json:"address_id"`
}

// CreateClientRequest represents the request to create a new client
// @Description Client creation request object
// @Name CreateClientRequest
// @Id CreateClientRequest
// @Property client_name type string description="Client's first name" example="John" required=true
// @Property client_surname type string description="Client's last name" example="Doe" required=true
// @Property birthday type string description="Client's birthday (YYYY-MM-DD)" example="1990-01-01" required=true format=date
// @Property gender type string description="Gender (M/F)" example="M" required=true enum=["M", "F"]
// @Property country type string description="Country of residence" example="USA"
// @Property city type string description="City of residence" example="New York"
// @Property street type string description="Street address" example="Broadway 100"
type CreateClientRequest struct {
	ClientName    string     `json:"client_name"`
	ClientSurname string     `json:"client_surname"`
	Birthday      CustomTime `json:"birthday"`
	Gender        string     `json:"gender"`
	Country       string     `json:"country"`
	City          string     `json:"city"`
	Street        string     `json:"street"`
}

// SearchClientResponse represents client details
// @Description Client details response object
// @Name SearchClientResponse
// @Id SearchClientResponse
// @Property client_name type string description="Client's first name" example="John"
// @Property client_surname type string description="Client's last name" example="Doe"
// @Property birthday type string description="Client's birthday" example="1990-01-01" format=date
// @Property gender type string description="Gender (M/F)" example="M"
// @Property registration_date type string description="Registration date" example="2023-10-01T12:00:00Z" format=date-time
// @Property country type string description="Country of residence" example="USA"
// @Property city type string description="City of residence" example="New York"
// @Property street type string description="Street address" example="Broadway 100"
type SearchClientResponse struct {
	ClientName       string    `json:"client_name" db:"client_name"`
	ClientSurname    string    `json:"client_surname" db:"client_surname"`
	Birthday         time.Time `json:"birthday" db:"birthday"`
	Gender           string    `json:"gender" db:"gender"`
	RegistrationDate time.Time `json:"registration_date" db:"registration_date"`
	Country          string    `json:"country" db:"country"`
	City             string    `json:"city" db:"city"`
	Street           string    `json:"street" db:"street"`
}
