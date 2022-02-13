package message_broker

type Publisher struct {
}

func NewPublisher() *Publisher {
	return &Publisher{}
}

func (p *Publisher) Publish(url string, body interface{}) error {
	return nil
}
