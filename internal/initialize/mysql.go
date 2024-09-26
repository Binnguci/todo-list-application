package initialize

import (
	"fmt"
	"github.com/binnguci/todo-app/global"
	"github.com/binnguci/todo-app/internal/models"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

func checkErrorPanic(err error, errString string) {
	if err != nil {
		global.Logger.Error(errString, zap.Error(err))
		panic(errString)
	}
}

func InitMySQL() {
	m := global.Config.Mysql
	dsn := "%s:%s@tcp(%s:%v)/%s?charset=utf8mb4&parseTime=True&loc=Local"
	var s = fmt.Sprintf(dsn, m.Username, m.Password, m.Host, m.Port, m.Dbname)
	fmt.Println(s)
	db, err := gorm.Open(mysql.Open(s), &gorm.Config{
		SkipDefaultTransaction: false,
	})
	checkErrorPanic(err, "Initialize MySQL error")
	global.Logger.Info("Initialize MySQL success")
	global.Mdb = db
	SetPool()
	migrateTables()
}
func SetPool() {
	m := global.Config.Mysql
	sqlDB, err := global.Mdb.DB()
	if err != nil {
		global.Logger.Error("SetPool error", zap.Error(err))
	}
	sqlDB.SetMaxIdleConns(global.Config.Mysql.MaxIdleConns)
	sqlDB.SetMaxOpenConns(global.Config.Mysql.MaxOpenConns)
	sqlDB.SetConnMaxLifetime(time.Duration(m.ConnMaxLifeTime))
}

func migrateTables() {
	err := global.Mdb.AutoMigrate(
		&models.User{},
	)
	if err != nil {
		global.Logger.Error("Migrate tables error", zap.Error(err))
	}
}
