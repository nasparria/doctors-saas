// pkg/email/mailtrap_api_sender.go
package email

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/sirupsen/logrus"
)

type MailtrapAPISender struct {
	apiToken string
	from     string
}

type Sender interface {
	Send(to, subject, body string) error
}

// NewMailtrapAPISender creates a new instance of MailtrapAPISender using environment variables
func NewMailtrapAPISender() Sender {
	return &MailtrapAPISender{
		apiToken: os.Getenv("EMAIL_API_TOKEN"), // Loaded from .env
		from:     os.Getenv("EMAIL_FROM"),      // Loaded from .env
	}
}

func (s *MailtrapAPISender) Send(to, subject, body string) error {
	url := "https://send.api.mailtrap.io/api/send"
	method := "POST"

	payload := map[string]interface{}{
		"from": map[string]string{
			"email": s.from,
			"name":  "Mailtrap Test",
		},
		"to": []map[string]string{
			{"email": to},
		},
		"subject":  subject,
		"text":     body,
		"category": "Integration Test",
	}

	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	client := &http.Client{}
	req, err := http.NewRequest(method, url, bytes.NewBuffer(payloadBytes))
	if err != nil {
		return fmt.Errorf("failed to create new request: %w", err)
	}

	req.Header.Add("Authorization", "Bearer "+s.apiToken)
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send email: %w", err)
	}
	defer res.Body.Close()

	// Capture full response body for debugging
	bodyBytes, _ := io.ReadAll(res.Body)
	logrus.Printf("Mailtrap response: %s", string(bodyBytes))

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to send email: received status code %d, response: %s", res.StatusCode, string(bodyBytes))
	}

	return nil
}
