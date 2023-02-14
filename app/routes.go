package app

import (
	base "inventoryApp"
	"inventoryApp/app/actions/lines"
	"inventoryApp/app/actions/users"
	"inventoryApp/app/middleware"

	"github.com/gobuffalo/buffalo"
)

// SetRoutes for the application
func setRoutes(root *buffalo.App) {

	root.Use(middleware.Database)
	root.Use(middleware.ParameterLogger)
	root.Use(middleware.CSRF)
	//ROUTES LINE
	root.GET("/", lines.List)
	root.GET("/lines/{id}/show", lines.Show)
	root.POST("/create/line", lines.Create)
	root.GET(("/line/new"), lines.New)
	root.GET("/lines/{id}/edit", lines.Edit)
	root.DELETE("/lines/{id}/delete", lines.Delete)
	root.PUT("/lines/{id}/update", lines.Update)
	root.PUT("/lines/{id}/change-status", lines.ChangeStatus)

	//ROUTES USER

	root.GET("/users", users.List)
	root.GET("/user/new", users.User)
	root.POST("/create/user", users.Create)
	root.GET("/users/{id}/show", users.Show)
	root.GET("/users/{id}/edit", users.Edit)
	root.PUT("/users/{id}/update", users.Update)
	root.DELETE("/users/{id}/delete", users.Delete)
	//
	root.ServeFiles("/", base.Assets)
}
