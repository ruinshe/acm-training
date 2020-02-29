package users

import (
	"github.com/gogf/gf/container/gset"
	model "github.com/ruinshe/acm-training/app/model/users"
)

// UserDto - the data transfer object for user entities.
type UserDto struct {
	ID       int64  `json:"id,omitempty"`
	Name     string `json:"name,omitempty"`
	Email    string `json:"email,omitempty"`
	Phone    string `json:"phone,omitempty"`
	Password string `json:"password,omitempty"`
	Rating   int    `json:"rating,omitempty"`
}

// Of - creates a data transfer object according to the DB entity.
func Of(entity *model.Entity, columns *gset.Set) UserDto {
	result := UserDto{}
	if columns.Contains(model.ColumnID) {
		result.ID = entity.Id
	}
	if columns.Contains(model.ColumnName) {
		result.Name = entity.Name
	}
	if columns.Contains(model.ColumnEmail) {
		result.Email = entity.Email
	}
	if columns.Contains(model.ColumnPhone) {
		result.Phone = entity.Phone
	}
	if columns.Contains(model.ColumnPassword) {
		result.Password = entity.Password
	}
	if columns.Contains(model.ColumnRating) {
		result.Rating = entity.Rating
	}
	return result
}
