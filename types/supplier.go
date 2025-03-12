package types

type Supplier struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	AddresId    int    `json:"address_id"`
	PhoneNumber string `json:"phone_number"`
}

// SupplierDTO represents supplier details
// @Description Supplier data transfer object
// @Name SupplierDTO
// @Id SupplierDTO
// @Property name type string description="Supplier name" example="Acme Co." required=true
// @Property country type string description="Country of operation" example="USA" required=true
// @Property city type string description="City of operation" example="New York" required=true
// @Property street type string description="Street address" example="Wall Street 123" required=true
// @Property phone_number type string description="Contact phone number" example="+1234567890" required=true
type SupplierDTO struct {
	Name        string `json:"name" db:"name"`
	Country     string `json:"country" db:"country"`
	City        string `json:"city" db:"city"`
	Street      string `json:"street" db:"street"`
	PhoneNumber string `json:"phone_number" db:"phone_number"`
}
