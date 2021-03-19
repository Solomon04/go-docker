-- name: UpdateCity :exec
UPDATE user_settings SET primary_city = $2
WHERE id = $1;

-- name: UpdateState :exec
UPDATE user_settings SET primary_state = $2
WHERE id = $1;

-- name: UpdateZipCode :exec
UPDATE user_settings SET primary_zip_code = $2
WHERE id = $1;

-- name: UpdatePreferredLocationTypes :exec
UPDATE user_settings SET preferred_location_types = $2
WHERE id = $1;

-- name: UpdatePreferredDistance :exec
UPDATE user_settings SET preferred_distance = $2
WHERE id = $1;

-- name: UpdateCommentPreferences :exec
UPDATE user_settings SET notify_comments = $2
WHERE id = $1;

-- name: UpdateEventPreferences :exec
UPDATE user_settings SET notify_events = $2
WHERE id = $1;

-- name: UpdateRecommendationPreferences :exec
UPDATE user_settings SET notify_recommendations = $2
WHERE id = $1;

-- name: UpdateFollowerPreferences :exec
UPDATE user_settings SET notify_follower_activity = $2
WHERE id = $1;