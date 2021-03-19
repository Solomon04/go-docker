package entities

// Event mirrors a result row from its table
type Event struct {
	ID           int      `json:"id"`
	UUID         string   `json:"uuid"`
	CreatorUUID  *string  `json:"creatorUUID"`
	LocationUUID string   `json:"locationUUID"`
	StartTime    string   `json:"startTime"`
	EndTime      string   `json:"endTime"`
	Lat          float64  `json:"lat"`
	Long         float64  `json:"long"`
	Distance     *float64 `json:"distance"`
	LiveCount    int      `json:"liveCount"`
	CreatedAt    *string  `json:"createdAt"`
	UpdatedAt    *string  `json:"updatedAt"`
	DeletedAt    *string  `json:"deletedAt"`
}