package routes

import (
	"fiberTest/controller"
	"fiberTest/middlewares"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App)  {


	app.Post("/register",controller.Register)
	app.Post("/login",controller.Login)


	app.Use(middlewares.IsAuthenticated)

	app.Get("/user",controller.User)
	app.Get("/logout",controller.Logout)

	users := app.Group("/users", middlewares.IsAuthenticated)
	users.Get("/getall",controller.AllUsers)
	users.Get("/getone",controller.GetUser)
	users.Post("/update",controller.UpdateUser)
	users.Delete("/delete",controller.DeleteUser)
	users.Post("/create",controller.CreateUser)
	users.Get("/alert",controller.Alert)

	objects := app.Group("/objects", middlewares.IsAuthenticated)
	objects.Get("/getall",controller.AllObjects)
	objects.Get("/getallout",controller.AllObjectsOut)
	objects.Get("/getone",controller.GetObject)
	objects.Post("/update",controller.UpdateObject)
	objects.Delete("/delete",controller.DeleteObject)
	objects.Post("/create",controller.CreateObject)
	objects.Post("/out",controller.OutObject)
}