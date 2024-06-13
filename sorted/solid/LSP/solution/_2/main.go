package main

import "fmt"

// Parent
type Notifier interface {
	Send(message string) error
}

// Sub
type EmailNotifier struct{}

func (e EmailNotifier) Send(message string) error {
	fmt.Println("Sending email:", message)
	return nil // Simplified for example purposes
}

// Sub
type SMSNotifier struct{}

func (s SMSNotifier) Send(message string) error {
	fmt.Println("Sending SMS:", message)
	return nil
}

// Sub
type PushNotifier struct{}

func (p PushNotifier) Send(message string) error {
	fmt.Println("Sending push notification:", message)
	return nil
}

// Util
func SendNotification(notifier Notifier, message string) {
	err := notifier.Send(message)
	if err != nil {
		fmt.Println("Error sending notification:", err)
	}
}

func main() {
	emailNotifier := EmailNotifier{}
	smsNotifier := SMSNotifier{}
	pushNotifier := PushNotifier{}

	// Sending different types of notifications.
	SendNotification(emailNotifier, "Welcome to our service!")
	SendNotification(smsNotifier, "Your verification code is 123456")
	SendNotification(pushNotifier, "You have a new message")
}
