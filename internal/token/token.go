package token

import (
	"strings"
	"time"

	"github.com/UESTC-ACM/acm-training/internal/api"
	"github.com/dgrijalva/jwt-go"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/gcache"
)

const (
	expireTime          = 5 * 24 * time.Hour // 5 days
	headerAuthorization = "Authorization"
	headerBearer        = "Bearer"
	configSecret        = "token.secret"

	claimKeyGenerated = "generated"
	claimKeyPhone     = "phone"
)

var (
	logger = g.Log("token")
	// Since we don't allow multi-login so we need a memory cache for all users.
	tokenCache = gcache.New()
	jwtSecret  []byte
)

func init() {
	jwtSecret = []byte(g.Config().GetString(configSecret))
}

// getTokenForUser - get token for the user, or exists = false for non existing case.
func getTokenForUser(phone string) (token string, exists bool) {
	cached := gcache.Get(phone)
	if cached == nil {
		exists = false
	} else {
		token = cached.(string)
		exists = true
	}
	return
}

// CreateTokenForUser - creates the token of specific user, by their phone.
// Since we don't support multi-login, we will replace the existing cached
// token of the phone when this function called.
func CreateTokenForUser(phone string) string {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := make(jwt.MapClaims)
	claims[claimKeyGenerated] = time.Now().Unix()
	claims[claimKeyPhone] = phone

	token.Claims = claims
	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		logger.Fatalf("Error occurs when sign secret for phone: %v", phone)
	}

	tokenCache.Set(phone, tokenString, expireTime)
	return tokenString
}

// AuthenticationInterceptor - the interceptor for all need authorization APIs.
func AuthenticationInterceptor(r *ghttp.Request) {
	authorizationHeader := r.Header.Get(headerAuthorization)
	if authorizationHeader != "" {
		parts := strings.SplitN(authorizationHeader, " ", 2)
		if len(parts) != 2 || parts[0] != headerBearer || parts[1] == "" {
			logger.Debugf("Invalid token set from user: %v", authorizationHeader)
			r.Response.WriteJsonExit(api.Response{
				ErrorCode:    "SYS_UNAURHORIZED",
				ErrorMessage: "Should login to gain permission for this API.",
			})
			return
		}

		token := parts[1]
		parseAuth, err := jwt.Parse(token, func(*jwt.Token) (interface{}, error) {
			return jwtSecret, nil
		})
		if err != nil {
			r.Response.WriteJsonExit(api.Response{
				ErrorCode:    "SYS_INVALID_AUTHORIZATION_TOKEN",
				ErrorMessage: "Invalid authorization token, please re-login to correct it.",
			})
			return
		}

		//将token中的内容存入parmMap
		claims := parseAuth.Claims.(jwt.MapClaims)
		var phone interface{} = nil
		for key, value := range claims {
			if key == claimKeyPhone {
				phone = value
			}
		}
		cached := tokenCache.Get(phone)
		if cached == nil || cached.(string) != token {
			r.Response.WriteJsonExit(api.Response{
				ErrorCode:    "SYS_INVALID_AUTHORIZATION_TOKEN",
				ErrorMessage: "Invalid authorization token, please re-login to correct it.",
			})
			return
		}

		r.Middleware.Next()
	}
}
