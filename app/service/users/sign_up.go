package users

import (
	"database/sql"

	model "github.com/UESTC-ACM/acm-training/app/model/users"
	"github.com/UESTC-ACM/acm-training/internal/api"
	"github.com/UESTC-ACM/acm-training/internal/rating"
	apiUtils "github.com/UESTC-ACM/acm-training/internal/utils/api"
	"github.com/UESTC-ACM/acm-training/internal/utils/encrypt"
	"github.com/gogf/gf/frame/g"
)

var (
	logger      = g.Log("service")
	signUpError = api.Response{
		ErrorCode:    "USER_SIGN_UP_FAILED",
		ErrorMessage: "sign up process failed, pleas have another try",
	}
)

//SignUpRequest - the request body for sign up API
type SignUpRequest struct {
	Name              string `p:"name" v:"required|length:1,24#Plese input name|The length of name should be between :min and :max."`
	Email             string `p:"email" v:"required|length:1,100#Plese input email|The length of email should be between :min and :max."`
	Phone             string `p:"phone" v:"required|length:11,11#Plese input phone|The length of phone should be 11."`
	Password          string `p:"password" v:"required|length:8,30#Please input password|The password length should be in range between :min and :max."`
	ConfirmedPassword string `p:"confirmed_password" v:"required|length:8,30|same:password#Please input confirmed_password|The confirmed_password is not same as password."`
}

// SignUp - sign up a new user according to the request input.
func SignUp(user *SignUpRequest) api.Response {
	var persisted []model.Entity
	if err := model.Model.M.Structs(&persisted, model.Model.Where("email like ?", user.Email).Or("phone", user.Phone)); err != nil {
		if err != sql.ErrNoRows {
			return apiUtils.RecordDBError(logger, err)
		}
	} else {
		for _, userInDB := range persisted {
			if userInDB.Email == user.Email {
				return api.Response{
					ErrorCode:    "USER_EMAIL_IN_USED",
					ErrorMessage: "the email is in used for other user",
				}
			} else if userInDB.Phone == user.Phone {
				return api.Response{
					ErrorCode:    "USER_PHONE_IN_USED",
					ErrorMessage: "the phone is in used for other user",
				}
			}
		}
	}

	encryptedPassword, err := encrypt.Pasword(user.Password)
	if err != nil {
		logger.Fatalf("Error occurs when encrypting user password for user: %v", user)
		return signUpError
	}
	entity := model.Entity{
		Name:     user.Name,
		Email:    user.Email,
		Phone:    user.Phone,
		Password: encryptedPassword,
		Rating:   rating.DefaultRating,
	}
	result, err := model.Save(entity)
	if err != nil {
		return apiUtils.RecordDBError(logger, err)
	} else if affected, err := result.RowsAffected(); affected == 0 || err != nil {
		return signUpError
	}

	id, err := result.LastInsertId()
	if err != nil {
		return signUpError
	}

	return api.Response{
		Data: g.Map{
			"id": id,
		},
	}
}
