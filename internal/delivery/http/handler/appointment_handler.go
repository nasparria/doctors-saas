// internal/delivery/http/handler/appointment_handler.go
package handler

import (
	"doctors/internal/domain"
	"doctors/internal/usecase"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type AppointmentHandler struct {
	appointmentUseCase usecase.AppointmentUseCase
}

func NewAppointmentHandler(appointmentUseCase usecase.AppointmentUseCase) *AppointmentHandler {
	return &AppointmentHandler{appointmentUseCase: appointmentUseCase}
}

func (h *AppointmentHandler) CreateAppointment(c *gin.Context) {
	var appointmentRequest struct {
		PatientID uint      `json:"patient_id" binding:"required"`
		DateTime  time.Time `json:"date_time" binding:"required"`
		Notes     string    `json:"notes"`
	}

	if err := c.ShouldBindJSON(&appointmentRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	appointment := domain.Appointment{
		PatientID: appointmentRequest.PatientID,
		DateTime:  appointmentRequest.DateTime,
		Notes:     appointmentRequest.Notes,
	}

	if err := h.appointmentUseCase.CreateAppointment(c.Request.Context(), &appointment); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create appointment"})
		return
	}

	c.JSON(http.StatusCreated, appointment)
}

func (h *AppointmentHandler) GetAppointment(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid appointment ID"})
		return
	}

	appointment, err := h.appointmentUseCase.GetAppointment(c.Request.Context(), uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Appointment not found"})
		return
	}

	c.JSON(http.StatusOK, appointment)
}

func (h *AppointmentHandler) UpdateAppointment(c *gin.Context) {
	var appointment domain.Appointment
	if err := c.ShouldBindJSON(&appointment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid appointment ID"})
		return
	}
	appointment.ID = uint(id)

	if err := h.appointmentUseCase.UpdateAppointment(c.Request.Context(), &appointment); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update appointment"})
		return
	}

	c.JSON(http.StatusOK, appointment)
}

func (h *AppointmentHandler) DeleteAppointment(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid appointment ID"})
		return
	}

	if err := h.appointmentUseCase.DeleteAppointment(c.Request.Context(), uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete appointment"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Appointment deleted successfully"})
}

func (h *AppointmentHandler) GetAppointmentsByDate(c *gin.Context) {
	dateStr := c.Query("date")
	date, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid date format"})
		return
	}

	appointments, err := h.appointmentUseCase.GetAppointmentsByDate(c.Request.Context(), date)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch appointments"})
		return
	}

	c.JSON(http.StatusOK, appointments)
}
