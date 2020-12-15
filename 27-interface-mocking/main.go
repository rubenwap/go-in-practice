package msg


type Message struct {
	email, subject string
	body []byte
}

func (m *Message) Send(email, subject string, body []byte) error {
	return nil
}

// write an interface that describes the methods of Message,
// so you can have your tests using that interface in its declarations
// instead of the real message type

type Messager interface {
	Send(email, subject string, body []byte) error
}

// function to be tested via mocks
func Alert(m Messager, problem []byte) error {
	return m.Send("noc@example.com", "Critical Error", problem)
}


