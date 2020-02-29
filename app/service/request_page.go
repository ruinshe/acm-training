package service

import (
	"fmt"
	"strconv"

	"github.com/gogf/gf/net/ghttp"
	"github.com/ruinshe/acm-training/internal/api"
)

const (
	requestParameterPage = "page"
	requestParameterSize = "size"

	defaultPage = 0
	defaultSize = 50
)

// RequestPage - the page informatin passed from UI side.
type RequestPage struct {
	Page int `json:"page"`
	Size int `json:"size"`
}

// ParseRequestPage - create a ParseRequestPage object from the request parameter.
func ParseRequestPage(r *ghttp.Request) (*RequestPage, *api.Error) {
	pageVar := r.Get(requestParameterPage)
	sizeVar := r.Get(requestParameterSize)
	page := defaultPage
	var err error
	if pageVar != nil {
		if page, err = strconv.Atoi(pageVar.(string)); err != nil {
			return nil, &api.Error{
				ErrorCode:    "SYS_INVALID_REQUEST_PARAMETER",
				ErrorMessage: fmt.Sprintf("Invalid request parameter value of parameter 'page': %v", pageVar),
			}
		}
	}
	size := defaultSize
	if sizeVar != nil {
		if size, err = strconv.Atoi(pageVar.(string)); err != nil {
			return nil, &api.Error{
				ErrorCode:    "SYS_INVALID_REQUEST_PARAMETER",
				ErrorMessage: fmt.Sprintf("Invalid request parameter value of parameter 'size': %v", sizeVar),
			}
		}
	}
	return &RequestPage{Page: page, Size: size}, nil
}
