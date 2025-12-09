package routes

import (
	"database/sql"
	"uas/app/services"
	"uas/middleware"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

func SetupRoutes(app *fiber.App, postgreSQL *sql.DB, mongoDB *mongo.Database) {

	api := app.Group("/api/v1") // (tidak perlu login)

	// Autentikasi & Otorisasi 
	auth := api.Group("/auth")
	auth.Post("/login", services.Login)
	auth.Post("/refresh", services.Refresh)
	auth.Get("/profile", middleware.AuthRequired(), services.GetProfile)

	// Protected routes (perlu login) 
	protected := api.Group("", middleware.AuthRequired()) 
	userService := services.NewUserService(postgreSQL)

	// Users (Admin)
	protected.Post("/users", userService.CreateUser)
	protected.Get("/users", userService.GetAllUsers)
	protected.Get("/users/:id", userService.GetUserByID)
	protected.Put("/users/:id", userService.UpdateUser)
	protected.Delete("/users/:id", userService.DeleteUser)
	protected.Put("/users/:id/role", userService.UpdateUserRole)
}