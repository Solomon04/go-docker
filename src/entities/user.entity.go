package entities

// User mirrors a result row from its table
type User struct {
	ID               int      `json:"id"`
	UUID             string   `json:"uuid"`
	FirstName        *string  `json:"firstName"`
	LastName         *string  `json:"lastName"`
	Avatar           string   `json:"avatar"`
	Email            string   `json:"email"`
	Password         *string  `json:"password"`
	LastSeenAt       *string  `json:"lastSeenAt"`
	LastSeenLocation *string  `json:"lastSeenLocation"`
	CreatedAt        *string  `json:"createdAt"`
	UpdatedAt        *string  `json:"updatedAt"`
	DeletedAt        *string  `json:"deletedAt"`
}
