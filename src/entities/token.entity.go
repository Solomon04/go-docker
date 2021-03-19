package entities

// PersonalAccessToken mirrors a result row from its table
type PersonalAccessToken struct {
	ID         int     `json:"id"`
	UserID   string  `json:"userID"`
	Token      string  `json:"token"`
	Abilities  string  `json:"abilities"`
	LastUsedAt string  `json:"lastUsedAt"`
	CreatedAt  *string `json:"createdAt"`
	UpdatedAt  *string `json:"updatedAt"`
}