package configurations

import (
	"strconv"

	"github.com/gogf/gf/frame/g"
)

const (
	// ColumnIdentifier - the identifier column of entity.
	ColumnIdentifier = "identifier"
	// ColumnValue - the key column of entity.
	ColumnValue = "value"

	// keyCurrentSeasonID - the configuration key for current season id.
	keyCurrentSeasonID = "current_season_id"
)

var logger = g.Log("persistence")

// GetCurrentSeasonID - get currrent season id for other services.
func GetCurrentSeasonID() int {
	var row *Entity
	if err := Model.M.Scan(&row, ColumnIdentifier+" like ?", keyCurrentSeasonID); err != nil {
		logger.Fatalf("error occurs when finding current season id: %v", err.Error())
		return -1
	}
	id, err := strconv.Atoi(row.Value)
	if err != nil {
		logger.Fatalf("error occurs when parsing current season id: %v", err.Error())
		return -1
	}
	return id
}
