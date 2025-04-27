package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mhakimsaputra17/tickitz-api-go/internal/repository"
)

// ScheduleHandler handles schedule-related endpoints
type ScheduleHandler struct {
    scheduleRepo *repository.ScheduleRepository
}

// NewScheduleHandler creates a new ScheduleHandler instance
func NewScheduleHandler(scheduleRepo *repository.ScheduleRepository) *ScheduleHandler {
    return &ScheduleHandler{
        scheduleRepo: scheduleRepo,
    }
}

// GetSchedules handles the GET /schedules endpoint
func (h *ScheduleHandler) GetSchedules(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
        "message": "Dummy GetSchedules endpoint",
        "status": "success",
    })
}

// GetShowtimeSeats handles the GET /schedules/:id/seats endpoint
func (h *ScheduleHandler) GetShowtimeSeats(c *gin.Context) {
    id := c.Param("id")
    c.JSON(http.StatusOK, gin.H{
        "message": "Dummy GetShowtimeSeats endpoint",
        "id": id,
        "status": "success",
    })
}