package types

// UpdateAddressInput represents the request to update client's address
// @Description Address update request object (supports partial updates)
// @Name UpdateAddressInput
// @Id UpdateAddressInput
// @Property country type string description="New country" example="Canada" required=false
// @Property city type string description="New city" example="Toronto" required=false
// @Property street type string description="New street address" example="Main Street 45" required=false
type UpdateAddressInput struct {
	Country *string `json:"country" db:"country"`
	City    *string `json:"city" db:"city"`
	Street  *string `json:"street" db:"street"`
}
