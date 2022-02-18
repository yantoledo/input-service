package message_broker

type BrokerClientMock struct {
}

func NewBrokerClientMock() *BrokerClientMock {
	return &BrokerClientMock{}
}

func (p *BrokerClientMock) Publish(BrokerRequest) error {
	return nil
}
