package users

import (
	"github.com/gogf/gf/net/ghttp"
	"github.com/ruinshe/acm-training/app/service"
	"github.com/ruinshe/acm-training/internal/api"
)

// GetUsers - queries users according to specific search conditions.
func GetUsers(r *ghttp.Request) {
	page, err := service.ParseRequestPage(r)
	if err != nil {
		r.Response.WriteJsonExit(err)
		return
	}
	r.Response.WriteJsonExit(api.Response{
		Data: api.Page{
			CurrentPage:   page.Page,
			PageSize:      page.Size,
			TotalPages:    1,
			TotalElements: 0,
			Elements:      make([]interface{}, 0),
		},
	})
}
