package repositories

import (

	"github.com/solomon04/go-docker/graph/model"
)

func UpdateSettings(userSettingsInfo *model.UpdateSettings) {
	if userSettingsInfo.PrimaryCity != nil {
		// db.Queries.UpdateCity()
	}
}