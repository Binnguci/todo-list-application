package global

import (
	"github.com/binnguci/todo-app/pkg/setting"
	"gorm.io/gorm"
)

var (
	Config setting.Config
	Mdb    *gorm.DB
)
