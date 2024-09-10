package route

import (
	"master-proof-api/controller"
	"master-proof-api/middleware"

	"github.com/gofiber/fiber/v2"
)

func SetupRoute(app *fiber.App, userController controller.UserController, learningMaterialController controller.LearningMaterialController) {
	api := app.Group("/api")

	//USERS
	api.Post("/users", userController.Create)
	api.Get("users/profile", middleware.FirebaseAuthMiddleware(), userController.Find)
	api.Post("users/login", userController.Login)
	api.Post("/users/reset-password", userController.ResetPassword)

	//LEARNING_MATERIAL
	api.Get("/learning-materials", middleware.FirebaseAuthMiddleware(), learningMaterialController.FindAll)
}
