package users

import (
	"strconv"

	"github.com/gogf/gf/container/gset"
	"github.com/ruinshe/acm-training/app/model/configurations"
	"github.com/ruinshe/acm-training/internal/api"
)

const (
	// ColumnID - the id column of entity.
	ColumnID = "id"
	// ColumnSeasonID - the season_id column of entity.
	ColumnSeasonID = "seasonId"
	// ColumnName - the name column of entity.
	ColumnName = "name"
	// ColumnEmail - the email column of entity.
	ColumnEmail = "email"
	// ColumnPhone - the phone column of entity.
	ColumnPhone = "phone"
	// ColumnPassword - the password column of entity.
	ColumnPassword = "password"
	// ColumnRating - the rating column of entity.
	ColumnRating = "rating"
)

var (
	// AllQueryColumns - all columns used for query, which doesn't include sensitive columns (e.g. password).
	AllQueryColumns = gset.NewFrom([]string{
		ColumnID, ColumnSeasonID, ColumnName, ColumnEmail, ColumnPhone, ColumnRating})
)

// Scan - scan the table to specific result, and apply default filter about current season id.
func Scan(requestPage *api.RequestPage, pointer interface{}, where interface{}, args ...interface{}) error {
	return Model.M.Page(requestPage.Page, requestPage.Size).
		Scan(pointer, applySeasonIDFilter(where, args...))
}

// Count - count the query results by applying default filter about current season id.
func Count(where interface{}, args ...interface{}) (int, error) {
	return Model.Count(applySeasonIDFilter(where, args...))
}

func applySeasonIDFilter(where interface{}, args ...interface{}) *arModel {
	return Model.Where(ColumnSeasonID+" = "+strconv.Itoa(configurations.GetCurrentSeasonID())).
		And(where, args...)
}
