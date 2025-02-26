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

type CreateClientRequest struct {
	ClientName    string     `json:"client_name"`
	ClientSurname string     `json:"client_surname"`
	Birthday      CustomTime `json:"birthday"`
	Gender        string     `json:"gender"`
	Country       string     `json:"country"`
	City          string     `json:"city"`
	Street        string     `json:"street"`
}

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
