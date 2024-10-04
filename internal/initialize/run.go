package initialize

func Run() {
	LoadConfig()
	InitMySQL()
	r := InitRouter()
	r.Run(":8083")
}
