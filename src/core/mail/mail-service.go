package mail

import "fmt"

func StartMailService() {
	fmt.Println("Starting Mail Service...")
}

// SendPrompt sends a prompt to a specific agent via email (mocked for now)
func SendPrompt(config MailConfig, to string, subject string, body string) error {
	fmt.Printf("MOCK MAIL: Sending prompt to %s, Subject: %s, Body: %s\n", to, subject, body)
	return nil
}
