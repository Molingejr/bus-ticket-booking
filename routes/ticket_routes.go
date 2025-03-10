package routes

import (
	"ticket-booking/controllers"
	"github.com/gin-gonic/gin"
)

func TicketRoutes(r *gin.Engine) {
	r.POST("/book-ticket", controllers.BookTicket)
}
