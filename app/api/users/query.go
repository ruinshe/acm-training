package users

import (
	"github.com/gogf/gf/net/ghttp"
	model "github.com/ruinshe/acm-training/app/model/users"
	service "github.com/ruinshe/acm-training/app/service/users"
	"github.com/ruinshe/acm-training/internal/api"
	apiUtils "github.com/ruinshe/acm-training/internal/utils/api"
)

const (
	requestParameterPhone = "phone"
	requestParameterName  = "name"
)

// GetUsers - queries users according to specific search conditions.
func GetUsers(r *ghttp.Request) {
	page, apiError := apiUtils.ParseRequestPage(r)
	if apiError != nil {
		r.Response.WriteJsonExit(apiError)
		return
	}
	phone := r.Get(requestParameterPhone)
	name := r.Get(requestParameterName)

	where := model.Model.Where("1=1")
	if phone != nil {
		where = model.Model.Where(model.ColumnPhone+" like ?", "%"+phone.(string)+"%")
	} else if name != nil {
		where = model.Model.Where(model.ColumnName+" like ?", "%"+name.(string)+"%")
	}

	results, err := service.Query(where, page, model.AllQueryColumns)
	if err != nil {
		r.Response.WriteJsonExit(api.Response{
			Error: &api.Error{
				ErrorCode:    "SYS_INTERNAL_ERROR_OCCURS",
				ErrorMessage: "internal error occurs when perform query",
			},
		})
		return
	}
	r.Response.WriteJsonExit(api.Response{
		Data: results,
	})
}
