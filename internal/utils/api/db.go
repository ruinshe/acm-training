package api

import (
	"github.com/gogf/gf/os/glog"
	"github.com/ruinshe/acm-training/internal/api"
)

// RecordDBError - records a DB error and return corresponding error to user side.
func RecordDBError(logger *glog.Logger, err error) api.Response {
	logger.Fatalf("error occurs when query against databse: %v", err.Error())
	return api.Response{
		ErrorCode:    "SYS_DB_ERROR",
		ErrorMessage: "error occurs when query against databse.",
	}
}
