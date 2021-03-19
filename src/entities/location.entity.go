package entities

// Location mirrors a result row from its table
type Location struct {
	ID                  int      `json:"id"`
	UUID                string   `json:"uuid"`
	Name                string   `json:"name"`
	Slug                *string  `json:"slug"`
	Avatar              string   `json:"avatar"`
	Description         *string  `json:"description"`
	LocationType        string   `json:"locationType"`
	Street              string   `json:"street"`
	City                string   `json:"city"`
	State               string   `json:"state"`
	Zip                 string   `json:"zip"`
	Lat                 float64  `json:"lat"`
	Long                float64  `json:"long"`
	Distance            *float64 `json:"distance"`
	ScheduleRule        *string  `json:"scheduleRule"`
	IsActive            bool     `json:"isActive"`
	HasCovidRestriction bool     `json:"hasCovidRestriction"`
	AppointmentTypeID   string   `json:"appointmentTypeID"`
	CalendarID          string   `json:"calendarID"`
	CreatedAt           string   `json:"createdAt"`
	UpdatedAt           string   `json:"updatedAt"`
	DeletedAt           string   `json:"deletedAt"`
}
