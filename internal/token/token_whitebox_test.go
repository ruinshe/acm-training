package token

import (
	"time"

	"bou.ke/monkey"
	"github.com/dgrijalva/jwt-go"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("create token for user", func() {
	It("should gnerated correct JWT string", func() {
		mockTime := time.Date(2020, 2, 23, 0, 0, 0, 0, time.UTC)
		monkey.Patch(time.Now, func() time.Time {
			return mockTime
		})
		defer monkey.UnpatchAll()

		phone := "13122222222"
		tokenString := CreateTokenForUser(phone)
		Ω(tokenString).Should(Equal("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9." +
			"eyJnZW5lcmF0ZWQiOjE1ODI0MTYwMDAsInBob25lIjoiMTMxMjIyMjIyMjIifQ." +
			"akiJs0lJb3fifksZN3bL6IT4KGV57yhiohDajLp-a60"))

		claims := AuthorizationClaims{}
		_, err := jwt.ParseWithClaims(tokenString, &claims, func(*jwt.Token) (interface{}, error) {
			return jwtSecret, nil
		})
		Ω(err).Should(BeNil())
		Ω(claims).Should(Equal(AuthorizationClaims{
			Generated: mockTime.Unix(),
			Phone:     phone,
		}))
	})
})
