package facebook

import (
	"log"
	"strconv"

	"github.com/yantoledo/input-service/api/entity/facebook_events"
	"github.com/yantoledo/input-service/api/service/customer_service"
	"github.com/yantoledo/input-service/api/service/message_service"
	"github.com/yantoledo/input-service/api/usecase/process_customer"
	"github.com/yantoledo/input-service/api/usecase/process_message"
	"github.com/yantoledo/input-service/infra/http_protocol"
	"github.com/yantoledo/input-service/infra/message_broker"
)

type ProcessMessageEvent struct {
}

func NewProcessMessageEvent() *ProcessMessageEvent {
	return &ProcessMessageEvent{}
}

func (p *ProcessMessageEvent) Execute(event facebook_events.Event) error {

	customerProcessed, err := p.processCustomer(event)
	if err != nil {
		log.Panic("Fail to process Customer from facebook")
		return err
	}

	err = p.processMessage(event, customerProcessed)
	if err != nil {
		log.Panic("Fail to process Text Message from Facebook")
		return err
	}

	return nil
}

func (p *ProcessMessageEvent) processCustomer(event facebook_events.Event) (process_customer.CustomerDtoOutput, error) {
	httpClient := http_protocol.NewHttpClient()
	customerService := customer_service.NewCustomerService(httpClient)
	customerUseCase := process_customer.NewProcessCustomer(customerService)

	uniqueID, err := strconv.Atoi(event.Entry[0].Messaging[0].Sender.ID)

	if err != nil {
		return process_customer.CustomerDtoOutput{}, err
	}

	uniqueClientID, err := strconv.Atoi(event.Entry[0].Messaging[0].Recipient.ID)
	if err != nil {
		return process_customer.CustomerDtoOutput{}, err
	}

	customerProcessed, err := customerUseCase.Execute(process_customer.CustomerDtoInput{
		UniqueID:       uniqueID,
		UniqueClientID: uniqueClientID,
		Source:         1,
	})
	if err != nil {
		return process_customer.CustomerDtoOutput{}, err
	}

	return customerProcessed, nil
}

func (p *ProcessMessageEvent) processMessage(event facebook_events.Event, customerProcessed process_customer.CustomerDtoOutput) error {

	BrokerClient := message_broker.NewBrokerClient()

	messageService := message_service.NewMessageService(BrokerClient)

	messageUseCase := process_message.NewProcessMessage(messageService)
	err := messageUseCase.Execute(process_message.MessageDtoInput{
		Text:     event.Entry[0].Messaging[0].Message.Text,
		Type:     "text",
		MediaUrl: "",
		Customer: customerProcessed,
	})

	if err != nil {
		return err
	}

	return nil
}
