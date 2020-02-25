package router_test

import (
	"fmt"
	"time"

	"github.com/UESTC-ACM/acm-training/internal/token"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

const port = 47088

var _ = Describe("authorization token interceptor", func() {
	SetDefaultEventuallyTimeout(5 * time.Second)
	s := g.Server("authorization-token-interceptor")
	s.Group("/router-test/authorization-token-interceptor", func(group *ghttp.RouterGroup) {
		group.GET("/ping", func(r *ghttp.Request) {
			r.Response.WriteJsonExit(g.Map{
				"message": "pong",
			})
		})
		group.POST("/login", func(r *ghttp.Request) {
			tokenString := token.CreateTokenForUser("phone")
			// Purpose for easy testing, so we return the token as response directly.
			r.Response.WriteExit(tokenString)
		})
	})
	s.Group("/router-test/authorization-token-interceptor", func(group *ghttp.RouterGroup) {
		group.Middleware(token.AuthenticationInterceptor)
		group.GET("/need-login", func(r *ghttp.Request) {
			r.Response.WriteJsonExit(g.Map{
				"message": "ok",
			})
		})
	})
	s.SetPort(port)
	s.Start()
	defer s.Shutdown()

	Context("basic router rules for non-authorization API.", func() {
		It("should return normally without authorization error.", func() {
			client := ghttp.NewClient()
			client.SetPrefix(fmt.Sprintf("http://127.0.0.1:%v", port))
			Ω(client.GetContent("/router-test/authorization-token-interceptor/ping")).
				Should(MatchJSON(`{"message": "pong"}`))
		})
	})

	Context("router rules for authorization needed API.", func() {
		It("should return authorization error when user not login.", func() {
			client := ghttp.NewClient()
			client.SetPrefix(fmt.Sprintf("http://127.0.0.1:%v", port))
			Ω(client.GetContent("/router-test/authorization-token-interceptor/need-login")).
				Should(MatchJSON(`{
					"error_code":    "SYS_UNAURHORIZED",
					"error_message": "Should login to gain permission for this API."
				}`))
		})
		It("should return normally wihout authorization error after login.", func() {
			client := ghttp.NewClient()
			client.SetPrefix(fmt.Sprintf("http://127.0.0.1:%v", port))

			token := client.PostContent("/router-test/authorization-token-interceptor/login")
			Ω(token).ShouldNot(BeEmpty())

			client.SetHeader("Authorization", "Bearer "+token)
			Ω(client.GetContent("/router-test/authorization-token-interceptor/need-login")).
				Should(MatchJSON(`{"message": "ok"}`))
		})
	})
})
