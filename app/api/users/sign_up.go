package users

import (
	service "github.com/UESTC-ACM/acm-training/app/service/users"
	"github.com/UESTC-ACM/acm-training/internal/api"
	"github.com/gogf/gf/net/ghttp"
)

// SignUp - sign up a new user according to the basic fields.
func SignUp(r *ghttp.Request) {
	var user *service.SignUpRequest
	if err := r.Parse(&user); err != nil {
		r.Response.WriteJsonExit(api.Response{
			ErrorCode:    "SYS_INVALID_REQUEST_BODY",
			ErrorMessage: err.Error(),
		})
	}

	r.Response.WriteJsonExit(service.SignUp(user))
}
