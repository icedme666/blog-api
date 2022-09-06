package app

import (
	"github.com/astaxie/beego/validation"
	"xiamei.guo/blog-api/pkg/logging"
)

func MakeErrors(errors []*validation.Error) {
	for _, err := range errors {
		logging.Info(err.Key, err.Message)
	}
	return
}
