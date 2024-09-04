# Doctor SaaS

## Table of Contents
- [Doctor SaaS](#doctor-saas)
  - [Table of Contents](#table-of-contents)
  - [Project Overview](#project-overview)
  - [Features](#features)
  - [Technologies Used](#technologies-used)
  - [Installation](#installation)
  - [Configuration](#configuration)
  - [Usage](#usage)
  - [API Endpoints](#api-endpoints)
  - [Contributing](#contributing)
  - [License](#license)

## Project Overview
Doctor SaaS is a web-based service for managing doctor appointments and patients. It provides a RESTful API to manage patients, appointments, and email notifications for appointment confirmations. The application is built using Go, PostgreSQL, and integrates Kafka for messaging.

## Features
- **Patient Management**: Create, update, delete, and list patients.
- **Appointment Management**: Schedule, update, delete, and list appointments.
- **Email Notifications**: Sends appointment confirmation emails to patients using Mailtrap API.
- **Kafka Integration**: Message consumption from Kafka for various application events.

## Technologies Used
- **Backend**: Go 
- **Database**: PostgreSQL
- **Messaging**: Kafka
- **Email Service**: Mailtrap API
- **Web Framework**: Gin (for REST API)
- **Containerization**: Docker
- **Configuration**: Viper  for environment variable management

## Installation
1. **Clone the Repository**:
   ```
   git clone https://github.com/nasparria/doctors-saas.git
   cd doctors-saas
   ```

2. **Install Dependencies**:
   Ensure you have Go installed on your machine. You can download it from [here](https://golang.org/dl/).
   Install dependencies:
   ```
   go mod download
   ```

3. **Run the Application Locally**:
   You can run the application locally with:
   ```
   go run cmd/api/main.go
   ```

## Configuration

1. **Environment Variables**:
   The application requires a few environment variables. Create a `.env` file in the project root with the following structure:

   ```
   SERVER_PORT=8080
   DB_HOST=localhost
   DB_PORT=5432
   DB_USER=postgres
   DB_PASSWORD=yourpassword
   DB_NAME=doctor_saas

   EMAIL_API_TOKEN=your-mailtrap-api-token
   EMAIL_FROM=your-email@example.com

   KAFKA_BROKERS=localhost:9092
   KAFKA_GROUP_ID=doctor_saas_group
   ```

2. **Docker**:
   To run the application using Docker, use the following commands:

   ```
   docker-compose up --build
   ```

   This will spin up the application along with the PostgreSQL and Kafka containers.

## Usage

Once the application is running, you can interact with the API through HTTP requests using tools like Postman or curl.

Example: Creating an appointment using Postman:

- Endpoint: `POST /api/v1/appointments`
- Body:
  ```json
  {
    "patient_id": 1,
    "date_time": "2023-09-16T14:30:00Z",
    "notes": "Regular checkup"
  }
  ```

## API Endpoints

[List your API endpoints here]

## Contributing

Contributions are welcome! Please follow these steps to contribute:

1. Fork the repository.
2. Create a new branch (`git checkout -b feature-branch`).
3. Make your changes.
4. Commit your changes (`git commit -m 'Add new feature'`).
5. Push to the branch (`git push origin feature-branch`).
6. Open a Pull Request.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.