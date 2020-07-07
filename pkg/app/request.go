package app

import (
	"github.com/astaxie/beego/validation"
	"github.com/evanxzj/go-gin-example/pkg/logger"
)

func MarkError(errors []*validation.Error) {
	for _, err := range errors {
		logger.Info(err.Key, err.Message)
	}

	return
}
