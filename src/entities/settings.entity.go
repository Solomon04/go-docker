package entities

// Settings mirrors a result row from UserSettings table
// holds the user setting configurations
type Settings struct {
	ID                     int     `json:"id"`
	UserUUID               string  `json:"userUUID"`
	PreferredLocationTypes *string `json:"preferredLocationTypes"`
	PreferredDistance      int     `json:"preferredDistance"`
	PrimaryCity      *string  `json:"primaryCity"`
	PrimaryState     *string  `json:"primaryState"`
	PrimaryZipCode   *string  `json:"primaryZipCode"`
	PrimaryLat       *float64 `json:"primaryLat"`
	PrimaryLong      *float64 `json:"primaryLong"`
	NotifyComments         bool    `json:"notifyComments"`
	NotifyEvents           bool    `json:"notifyEvents"`
	NotifyRecommendations  bool    `json:"notifyRecommendations"`
	NotifyFollowerActivity bool    `json:"notifyFollowerActivity"`
	CreatedAt              *string `json:"createdAt"`
	UpdatedAt              *string `json:"updatedAt"`
	DeletedAt              *string `json:"deletedAt"`
}