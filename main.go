package main

import (
    "log"
    "ticket-booking/routes"
    "ticket-booking/utils"
    "ticket-booking/models"
    "github.com/gin-gonic/gin"
)

func main() {
    utils.InitDB() // Initialize the database

    // Automatically migrate your models
	utils.DB.AutoMigrate(&models.User{}, &models.Ticket{})

    router := gin.Default()

    // Register routes
    routes.AuthRoutes(router)
    routes.TicketRoutes(router)

    // Start the server
    if err := router.Run(":8080"); err != nil {
        log.Fatal("Error starting server:", err)
    }
}
