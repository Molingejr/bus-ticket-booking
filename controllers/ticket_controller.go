package controllers

import (
	"ticket-booking/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Booking a ticket
func BookTicket(c *gin.Context) {
	var ticket models.Ticket
	if err := c.ShouldBindJSON(&ticket); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid data"})
		return
	}

	// Save the ticket (this can be customized with more complex business logic)
	if err := ticket.Save(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to book ticket"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Ticket booked successfully", "ticket": ticket})
}
