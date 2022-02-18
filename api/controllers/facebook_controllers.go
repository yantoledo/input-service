package controllers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	event "github.com/yantoledo/input-service/api/entity/facebook_events"
	service "github.com/yantoledo/input-service/api/service/customer_service"
	"github.com/yantoledo/input-service/api/service/message_service"
	"github.com/yantoledo/input-service/api/usecase/process_customer"
	"github.com/yantoledo/input-service/api/usecase/process_message"
	"github.com/yantoledo/input-service/infra/http_protocol"
	"github.com/yantoledo/input-service/infra/message_broker"
)

func MessageEventFromFacebook(w http.ResponseWriter, r *http.Request) {
	event := event.Event{}
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatal(err)
	}

	if err := json.Unmarshal(body, &event); err != nil {
		log.Fatal(err)
	}

	httpClient := http_protocol.NewHttpClient()
	customerService := service.NewCustomerService(httpClient)
	customerUseCase := process_customer.NewProcessCustomer(customerService)

	uniqueID, err := strconv.Atoi(event.Entry[0].Messaging[0].Sender.ID)

	if err != nil {
		log.Panic(err)
	}

	uniqueClientID, err := strconv.Atoi(event.Entry[0].Messaging[0].Recipient.ID)
	if err != nil {
		log.Panic(err)
	}

	customerProcessed, err := customerUseCase.Execute(process_customer.CustomerDtoInput{
		UniqueID:       uniqueID,
		UniqueClientID: uniqueClientID,
		Source:         1,
	})
	if err != nil {
		log.Panic(err)
	}

	BrokerClient := message_broker.NewBrokerClient()

	messageService := message_service.NewMessageService(BrokerClient)

	messageUseCase := process_message.NewProcessMessage(messageService)
	out, err := messageUseCase.Execute(process_message.MessageDtoInput{
		Text:     event.Entry[0].Messaging[0].Message.Text,
		Type:     "text",
		MediaUrl: "",
		Customer: customerProcessed,
	})

	if err != nil {
		log.Panic(err)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(out)

}

func FacebookAuthorization(w http.ResponseWriter, r *http.Request) {

	urlParams := r.URL.Query()
	//TODO: Need checks secret key before return the response

	resp, err := strconv.Atoi(urlParams["hub.challenge"][0])
	if err != nil {
		panic(err)
	}
	json.NewEncoder(w).Encode(resp)
}

func Ping(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("pong")
}
