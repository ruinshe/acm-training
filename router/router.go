package router

import (
	"github.com/UESTC-ACM/acm-training/app/api/users"
	"github.com/UESTC-ACM/acm-training/internal/api"
	"github.com/UESTC-ACM/acm-training/internal/token"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

func init() {
	s := g.Server()
	s.Group("/api/v1/", func(group *ghttp.RouterGroup) {
		group.Middleware(token.AuthenticationInterceptor)
		group.GET("/ping", func(r *ghttp.Request) {
			r.Response.WriteJsonExit(api.Response{
				Data: "pong",
			})
		})
	})
	s.Group("/api/v1/users", func(group *ghttp.RouterGroup) {
		group.POST("/login", users.Login)
		if g.Cfg().GetBool("api.signUpEnabled") {
			group.POST("/signup", users.SignUp)
		}
	})
}
