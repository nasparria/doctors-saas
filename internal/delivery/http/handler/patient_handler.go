// internal/delivery/http/handler/patient_handler.go
package handler

import (
	"net/http"
	"strconv"

	"doctors/internal/domain"
	"doctors/internal/usecase"
	"github.com/gin-gonic/gin"
)

type PatientHandler struct {
	patientUseCase usecase.PatientUseCase
}

func NewPatientHandler(patientUseCase usecase.PatientUseCase) *PatientHandler {
	return &PatientHandler{
		patientUseCase: patientUseCase,
	}
}

func (h *PatientHandler) CreatePatient(c *gin.Context) {
	var patient domain.Patient
	if err := c.ShouldBindJSON(&patient); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.patientUseCase.CreatePatient(c.Request.Context(), &patient); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create patient"})
		return
	}

	c.JSON(http.StatusCreated, patient)
}

func (h *PatientHandler) GetPatient(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid patient ID"})
		return
	}

	patient, err := h.patientUseCase.GetPatient(c.Request.Context(), uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Patient not found"})
		return
	}

	c.JSON(http.StatusOK, patient)
}

func (h *PatientHandler) UpdatePatient(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid patient ID"})
		return
	}

	var patient domain.Patient
	if err := c.ShouldBindJSON(&patient); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	patient.ID = uint(id)

	if err := h.patientUseCase.UpdatePatient(c.Request.Context(), &patient); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update patient"})
		return
	}

	c.JSON(http.StatusOK, patient)
}

func (h *PatientHandler) DeletePatient(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid patient ID"})
		return
	}

	if err := h.patientUseCase.DeletePatient(c.Request.Context(), uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete patient"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Patient deleted successfully"})
}

func (h *PatientHandler) ListPatients(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))

	patients, totalCount, err := h.patientUseCase.ListPatients(c.Request.Context(), page, pageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to list patients"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"patients":  patients,
		"total":     totalCount,
		"page":      page,
		"page_size": pageSize,
	})
}
