package types

type UpdateAddressInput struct {
	Country *string `json:"country" db:"country"`
	City    *string `json:"city" db:"city"`
	Street  *string `json:"street" db:"street"`
}
