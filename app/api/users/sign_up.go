package users

import (
	"github.com/gogf/gf/net/ghttp"
	service "github.com/ruinshe/acm-training/app/service/users"
	"github.com/ruinshe/acm-training/internal/api"
)

// SignUp - sign up a new user according to the basic fields.
func SignUp(r *ghttp.Request) {
	var user *service.SignUpRequest
	if err := r.Parse(&user); err != nil {
		r.Response.WriteJsonExit(api.Response{
			Error: &api.Error{
				ErrorCode:    "SYS_INVALID_REQUEST_BODY",
				ErrorMessage: err.Error(),
			},
		})
	}

	r.Response.WriteJsonExit(service.SignUp(user))
}
