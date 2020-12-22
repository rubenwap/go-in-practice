package main

// Message describes the message to be sent
type Message struct {
	email, subject string
	body []byte
}

// Send sends a message
func (m *Message) Send(email, subject string, body []byte) error {
	return nil
}

// write an interface that describes the methods of Message,
// so you can have your tests using that interface in its declarations
// instead of the real message type

// Messager is the interface that sends emails
type Messager interface {
	Send(email, subject string, body []byte) error
}

// Alert is the function to be tested via mocks
func Alert(m Messager, problem []byte) error {
	return m.Send("noc@example.com", "Critical Error", problem)
}


