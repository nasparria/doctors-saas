package main

import (
	"context"
	"doctors/config"
	"doctors/internal/delivery/http"
	"doctors/internal/infrastracture/database"
	"doctors/internal/infrastracture/messaging"
	"doctors/internal/repository"
	"doctors/internal/usecase"
	"doctors/pkg/email"
	"fmt"
	"log"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	emailSender := email.NewMailtrapAPISender()

	db, err := database.NewPostgresDB(cfg.DBHost, cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBPort)
	if err != nil {
		log.Fatalf("Failed to setup database: %v", err)
	}

	kafkaClient, err := messaging.NewKafkaClient(cfg.KafkaBrokers, "doctor_saas_topic", cfg.KafkaGroupID)
	if err != nil {
		log.Fatalf("Failed to create Kafka client: %v", err)
	}
	defer kafkaClient.Close()

	patientRepo := repository.NewPatientRepository(db)
	appointmentRepo := repository.NewAppointmentRepository(db)
	doctorRepo := repository.NewDoctorRepository(db)

	patientUseCase := usecase.NewPatientUseCase(patientRepo)
	appointmentUseCase := usecase.NewAppointmentUseCase(appointmentRepo, patientRepo, doctorRepo, emailSender)

	router := http.NewRouter(patientUseCase, appointmentUseCase)

	go func() {
		err := kafkaClient.ConsumeMessages(context.Background(), func(msg []byte) error {
			log.Printf("Received message: %s", string(msg))
			return nil
		})
		if err != nil {
			log.Printf("Error consuming Kafka messages: %v", err)
		}
	}()
	serverAddr := fmt.Sprintf("0.0.0.0:%d", cfg.ServerPort)
	log.Printf("Server starting on %s", serverAddr)
	if err := router.Run(serverAddr); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
