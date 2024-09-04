// internal/delivery/http/router.go
package http

import (
	"doctors/internal/delivery/http/handler"
	"doctors/internal/usecase"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func NewRouter(patientUseCase usecase.PatientUseCase, appointmentUseCase usecase.AppointmentUseCase) *gin.Engine {
	router := gin.New()

	// Add logging middleware
	router.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC1123),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}))
	router.Use(gin.Recovery())

	// Add a root route for basic testing
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Welcome to the Doctor SaaS API"})
	})

	patientHandler := handler.NewPatientHandler(patientUseCase)
	appointmentHandler := handler.NewAppointmentHandler(appointmentUseCase)

	v1 := router.Group("/api/v1")
	{
		patients := v1.Group("/patients")
		{
			patients.POST("/", patientHandler.CreatePatient)
			patients.GET("/:id", patientHandler.GetPatient)
			patients.PUT("/:id", patientHandler.UpdatePatient)
			patients.DELETE("/:id", patientHandler.DeletePatient)
			patients.GET("/", patientHandler.ListPatients) // Add this line
		}

		appointments := v1.Group("/appointments")
		{
			appointments.POST("/", appointmentHandler.CreateAppointment)
			appointments.GET("/:id", appointmentHandler.GetAppointment)
			appointments.PUT("/:id", appointmentHandler.UpdateAppointment) // Changed from patients to appointments
			appointments.DELETE("/:id", appointmentHandler.DeleteAppointment)
			appointments.GET("/", appointmentHandler.GetAppointmentsByDate)
		}
	}

	// Add a catch-all route for debugging
	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"message": "Route not found", "path": c.Request.URL.Path})
	})

	return router
}
