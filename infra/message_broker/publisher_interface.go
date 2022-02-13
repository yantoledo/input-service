package message_broker

type PublisherInterface interface {
	Publish() error
}
