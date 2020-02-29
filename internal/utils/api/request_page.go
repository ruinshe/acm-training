package api

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

// ParseRequestPage - create a ParseRequestPage object from the request parameter.
func ParseRequestPage(r *ghttp.Request) (*api.RequestPage, *api.Error) {
	pageVar := r.Get(requestParameterPage)
	sizeVar := r.Get(requestParameterSize)
	page := defaultPage
	var err error
	if pageVar != nil {
		if page, err = strconv.Atoi(pageVar.(string)); err != nil || page <= 0 {
			return nil, &api.Error{
				ErrorCode:    "SYS_INVALID_REQUEST_PARAMETER",
				ErrorMessage: fmt.Sprintf("Invalid request parameter value of parameter 'page': %v", pageVar),
			}
		}
	}
	size := defaultSize
	if sizeVar != nil {
		if size, err = strconv.Atoi(pageVar.(string)); err != nil || size <= 0 {
			return nil, &api.Error{
				ErrorCode:    "SYS_INVALID_REQUEST_PARAMETER",
				ErrorMessage: fmt.Sprintf("Invalid request parameter value of parameter 'size': %v", sizeVar),
			}
		}
	}
	return &api.RequestPage{Page: page, Size: size}, nil
}
