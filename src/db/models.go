// Code generated by sqlc. DO NOT EDIT.

package db

import (
	"database/sql"
	"time"
)

type Event struct {
	ID         int64
	Uuid       string
	CreatorID  sql.NullInt64
	LocationID int64
	StartTime  time.Time
	EndTime    time.Time
	Lat        float64
	Long       float64
	LiveCount  int32
	CreatedAt  sql.NullTime
	UpdatedAt  sql.NullTime
}

type Location struct {
	ID                  int64
	Uuid                string
	Name                string
	Slug                sql.NullString
	Image               string
	Description         sql.NullString
	LocationType        string
	Street              string
	City                string
	State               string
	Zip                 string
	Lat                 float64
	Long                float64
	ScheduleRule        sql.NullString
	IsActive            bool
	HasCovidRestriction bool
	AppointmentTypeID   sql.NullString
	CalendarID          sql.NullString
	CreatedAt           sql.NullTime
	UpdatedAt           sql.NullTime
}

type PersonalAccessToken struct {
	ID         int64
	UserID     int64
	Token      string
	Abilities  sql.NullString
	LastUsedAt sql.NullTime
	CreatedAt  sql.NullTime
	UpdatedAt  sql.NullTime
}

type User struct {
	ID               int64
	Uuid             string
	FirstName        sql.NullString
	LastName         sql.NullString
	Email            string
	OauthProvider    sql.NullString
	OauthIdentifier  sql.NullString
	Password         sql.NullString
	Avatar           string
	LastSeenAt       sql.NullTime
	LastSeenLocation sql.NullString
	CreatedAt        sql.NullTime
	UpdatedAt        sql.NullTime
	DeletedAt        sql.NullTime
}

type UserEvent struct {
	ID          int64
	EventID     int64
	UserID      int64
	RespondedAt time.Time
	IsGoing     bool
	AttendedAt  sql.NullTime
}

type UserFollow struct {
	ID              int32
	FollowerUserID  int64
	FollowingUserID int64
	AcceptedAt      sql.NullTime
	CreatedAt       sql.NullTime
	UpdatedAt       sql.NullTime
}

type UserLocation struct {
	ID         int64
	LocationID int64
	UserID     int64
	JoinedOn   time.Time
}

type UserSetting struct {
	ID                           int64
	UserID                       int64
	PreferredLocationTypes       sql.NullString
	PreferredDistance            int32
	PrimaryCity                  sql.NullString
	PrimaryState                 sql.NullString
	PrimaryZipCode               sql.NullString
	PrimaryLat                   sql.NullFloat64
	PrimaryLong                  sql.NullFloat64
	NotifyComments               bool
	NotifyEvents                 bool
	NotifyRecommendations        bool
	NotificationFollowerActivity bool
	CreatedAt                    sql.NullTime
	UpdatedAt                    sql.NullTime
}
