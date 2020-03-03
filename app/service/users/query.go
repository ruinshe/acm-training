package users

import (
	"github.com/gogf/gf/container/gset"
	"github.com/gogf/gf/util/gpage"
	model "github.com/ruinshe/acm-training/app/model/users"
	"github.com/ruinshe/acm-training/internal/api"
)

// Query - queries matching users according to search condition and request page.
// Since we need to hide many fields inside the entities, we wrap this function to
// avoid sending sensitive information to the frontend size.
func Query(where interface{}, requestPage *api.RequestPage, columns *gset.Set) (*api.Page, error) {
	var users []model.Entity
	err := model.Scan(requestPage, &users, where)
	if err != nil {
		logger.Fatalf("Error occurs when perform user query: %v", err.Error())
		return nil, err
	}

	count, err := model.Count(where)
	if err != nil {
		logger.Fatalf("Error occurs when figuring out the count of query: %v", err.Error())
		return nil, err
	}

	elements := make([]UserDto, len(users))
	for index, user := range users {
		elements[index] = Of(&user, columns)
	}

	totalPages := gpage.New(count, requestPage.Size, requestPage.Page, "").TotalPage

	return &api.Page{
		CurrentPage:   requestPage.Page,
		PageSize:      requestPage.Size,
		TotalPages:    totalPages,
		TotalElements: count,
		Elements:      elements,
	}, nil
}
