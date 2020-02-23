package users

import (
	"database/sql"

	userModel "github.com/UESTC-ACM/acm-training/app/model/user"
	"github.com/UESTC-ACM/acm-training/internal/api"
	"github.com/UESTC-ACM/acm-training/internal/token"
	"github.com/UESTC-ACM/acm-training/internal/utils/encrypt"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

// Login - login into the system.
func Login(r *ghttp.Request) {
	var request *struct {
		Phone    string `p:"phone" v:"required|length:11,11#Plese input phone|The length of phone should be 11."`
		Password string `p:"password" v:"required|length:8,30#Please input password|The password length should be in range between :min and :max."`
	}
	if err := r.Parse(&request); err != nil {
		r.Response.WriteJsonExit(api.Response{
			ErrorCode:    "SYS_INVALID_REQUEST_BODY",
			ErrorMessage: err.Error(),
		})
	}
	var persisted *userModel.Entity
	if err := userModel.Model.M.Struct(&persisted, userModel.Model.Where("phone like ?", request.Phone)); err != nil {
		if err == sql.ErrNoRows {
			r.Response.WriteJsonExit(api.Response{
				ErrorCode:    "USER_NOT_FOUND",
				ErrorMessage: "The user is not found in the system.",
			})
		} else {
			g.Log("api").Fatalf("Error occurs when finding existing user by phone: %v", err.Error())
			r.Response.WriteJsonExit(api.Response{
				ErrorCode:    "SYS_INTERNAL_ERROR_OCCURS",
				ErrorMessage: "Internal server error occurs.",
			})
		}
	}

	if encrypted, err := encrypt.Pasword(request.Password); err != nil || encrypted != persisted.Password {
		r.Response.WriteJsonExit(api.Response{
			ErrorCode:    "USER_INVALID_PASSOWRD",
			ErrorMessage: "The password is incorrect.",
		})
	}

	authorizationToken := token.CreateTokenForUser(persisted.Phone)
	r.Response.WriteJsonExit(api.Response{
		Data: struct {
			Token string `json:"token"`
		}{
			Token: authorizationToken,
		},
	})
}
