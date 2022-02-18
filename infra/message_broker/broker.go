package message_broker

type BrokerClient struct {
}

func NewBrokerClient() *BrokerClient {
	return &BrokerClient{}
}

func (p *BrokerClient) Publish(BrokerRequest) error {
	return nil
}
