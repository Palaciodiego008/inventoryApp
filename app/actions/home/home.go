package home

import (
	"inventoryApp/app/render"
)

var (
	// r is a buffalo/render Engine that will be used by actions
	// on this package to render render HTML or any other formats.
	r = render.Engine
)

// func Index(c buffalo.Context) error {
// 	return c.Render(http.StatusOK, r.HTML("home/index.plush.html"))
// }

// func IndexTest(c buffalo.Context) error {
// 	return c.Render(http.StatusOK, r.HTML("home/test.plush.html"))
// }
