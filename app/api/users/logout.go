package users

import (
	"github.com/gogf/gf/net/ghttp"
	"github.com/ruinshe/acm-training/internal/api"
	"github.com/ruinshe/acm-training/internal/token"
)

// Logout - log out from the system.
func Logout(r *ghttp.Request) {
	claims := r.Session.Get(token.SessionAuthorizationClaimsKey).(token.AuthorizationClaims)
	token.RemoveUserToken(claims.Phone)
	r.Response.WriteJsonExit(api.SuccessfulResponse)
}
