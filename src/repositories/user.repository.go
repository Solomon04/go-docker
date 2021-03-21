package repositories

import (

	"github.com/solomon04/go-docker/graph/model"
)

// https://github.com/holys/initials-avatar for creating an avatar
// https://pkg.go.dev/github.com/aws/aws-sdk-go/service/s3 for uploading the user avatar file
// https://github.com/oliveroneill/exponent-server-sdk-golang for handling notifications

func UpdateSettings(userSettingsInfo *model.UpdateSettings) {
	if userSettingsInfo.PrimaryCity != nil {
		// db.Queries.UpdateCity()
	}
}