// Code generated by sqlc. DO NOT EDIT.
// source: user_settings_queries.sql

package db

import (
	"context"
)

const updateCity = `-- name: UpdateCity :exec
UPDATE user_settings SET primary_city = $2
WHERE id = $1
`

func (q *Queries) UpdateCity(ctx context.Context) error {
	_, err := q.db.ExecContext(ctx, updateCity)
	return err
}

const updateCommentPreferences = `-- name: UpdateCommentPreferences :exec
UPDATE user_settings SET notify_comments = $2
WHERE id = $1
`

func (q *Queries) UpdateCommentPreferences(ctx context.Context) error {
	_, err := q.db.ExecContext(ctx, updateCommentPreferences)
	return err
}

const updateEventPreferences = `-- name: UpdateEventPreferences :exec
UPDATE user_settings SET notify_events = $2
WHERE id = $1
`

func (q *Queries) UpdateEventPreferences(ctx context.Context) error {
	_, err := q.db.ExecContext(ctx, updateEventPreferences)
	return err
}

const updateFollowerPreferences = `-- name: UpdateFollowerPreferences :exec
UPDATE user_settings SET notify_follower_activity = $2
WHERE id = $1
`

func (q *Queries) UpdateFollowerPreferences(ctx context.Context) error {
	_, err := q.db.ExecContext(ctx, updateFollowerPreferences)
	return err
}

const updatePreferredDistance = `-- name: UpdatePreferredDistance :exec
UPDATE user_settings SET preferred_distance = $2
WHERE id = $1
`

func (q *Queries) UpdatePreferredDistance(ctx context.Context) error {
	_, err := q.db.ExecContext(ctx, updatePreferredDistance)
	return err
}

const updatePreferredLocationTypes = `-- name: UpdatePreferredLocationTypes :exec
UPDATE user_settings SET preferred_location_types = $2
WHERE id = $1
`

func (q *Queries) UpdatePreferredLocationTypes(ctx context.Context) error {
	_, err := q.db.ExecContext(ctx, updatePreferredLocationTypes)
	return err
}

const updateRecommendationPreferences = `-- name: UpdateRecommendationPreferences :exec
UPDATE user_settings SET notify_recommendations = $2
WHERE id = $1
`

func (q *Queries) UpdateRecommendationPreferences(ctx context.Context) error {
	_, err := q.db.ExecContext(ctx, updateRecommendationPreferences)
	return err
}

const updateState = `-- name: UpdateState :exec
UPDATE user_settings SET primary_state = $2
WHERE id = $1
`

func (q *Queries) UpdateState(ctx context.Context) error {
	_, err := q.db.ExecContext(ctx, updateState)
	return err
}

const updateZipCode = `-- name: UpdateZipCode :exec
UPDATE user_settings SET primary_zip_code = $2
WHERE id = $1
`

func (q *Queries) UpdateZipCode(ctx context.Context) error {
	_, err := q.db.ExecContext(ctx, updateZipCode)
	return err
}
