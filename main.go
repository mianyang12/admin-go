package main

import (
	"github.com/gin-gonic/gin"
	comnon "naive-admin-go/common"
	"naive-admin-go/config"
	"naive-admin-go/db"
	"naive-admin-go/router"
	"time"
)

func init() {
	//模板加载
	comnon.LoadTemplate()
}
func main() {
	var Loc, _ = time.LoadLocation("Asia/Shanghai")
	time.Local = Loc

	app := gin.Default()

	config.Init()
	db.Init()
	router.Init(app)
	app.Run(":8800")
}
