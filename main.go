package main

import (
	"workspace/chat/handler"
	"workspace/chat/util"

	"github.com/gin-gonic/gin"
)

func main() {

	util.Connect()
	db := util.GetDB()

	r := gin.Default()
	r.Static("/site", "./chat/templates/site")
	r.LoadHTMLGlob("./chat/templates/*.html")

	apiRoutes := r.Group("/api")
	{
		apiRoutes.POST("/login", handler.Login(&db))
		apiRoutes.POST("/logout", handler.Logout())
		apiRoutes.POST("/msg", handler.CreateMessage(&db))
		apiRoutes.GET("/chats", handler.GetChats(&db))

	}
	viewRoutes := r.Group("/view")
	{
		viewRoutes.GET("/panel", handler.ViewPanel())

		viewRoutes.GET("/login", handler.ViewLogin())

		viewRoutes.GET("/pv/:username", handler.ViewPv(&db))
	}

	r.Run("127.0.0.1:8080")
	defer db.Close()
}
