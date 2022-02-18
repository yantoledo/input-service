package message_broker

type BrokerInterface interface {
	Publish(BrokerRequest) error
}
