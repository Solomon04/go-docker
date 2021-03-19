package entities

// Follow mirrors a result row from UserFollow table
type Follow struct {
	ID                int    `json:"id"`
	FollowerUserUUID  string `json:"followerUserUUID"`
	FollowingUserUUID string `json:"followingUserUUID"`
	AcceptedAt        string `json:"acceptedAt"`
	CreatedAt         string `json:"createdAt"`
	UpdatedAt         string `json:"updatedAt"`
}