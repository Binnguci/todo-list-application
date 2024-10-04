package initialize

import (
	"fmt"
	"github.com/binnguci/todo-app/global"
	"github.com/binnguci/todo-app/internal/model"
	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
	"time"
)

func checkErrorPanic(err error, errString string) {
	if err != nil {
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
	fmt.Println("Connect to MySQL success")
	global.Mdb = db
	SetPool()
	genTableDAO()
}

func SetPool() {
	m := global.Config.Mysql
	sqlDB, err := global.Mdb.DB()
	if err != nil {
		panic("Get DB error")
	}
	sqlDB.SetMaxIdleConns(global.Config.Mysql.MaxIdleConns)
	sqlDB.SetMaxOpenConns(global.Config.Mysql.MaxOpenConns)
	sqlDB.SetConnMaxLifetime(time.Duration(m.ConnMaxLifeTime))
}

// gen db sang model
func genTableDAO() {
	g := gen.NewGenerator(gen.Config{
		OutPath: "./internal/model",
		Mode:    gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface,
	})
	g.UseDB(global.Mdb)
	g.GenerateAllTable()
	g.Execute()
}

// gen model sang db
func migrateTables() {
	err := global.Mdb.AutoMigrate(
		&model.User{},
		&model.Role{},
	)
	if err != nil {
		panic("Migrate tables error")
	}
}
