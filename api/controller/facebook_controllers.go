package controller

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"sync"

	"github.com/yantoledo/input-service/api/entity/facebook_events"
	"github.com/yantoledo/input-service/api/usecase/facebook"
)

func MessageEventFromFacebook(w http.ResponseWriter, r *http.Request) {
	event := facebook_events.Event{}
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatal(err)
	}

	if err := json.Unmarshal(body, &event); err != nil {
		log.Fatal(err)
	}

	var wg sync.WaitGroup
	wg.Add(2)

	go func(event facebook_events.Event) {
		processEvent := facebook.NewProcessMessageEvent()
		if err := processEvent.Execute(event); err != nil {
			log.Fatal("Failed to process facebook event")
		}
		wg.Done()
	}(event)

	go func() {
		json.NewEncoder(w).Encode("Success")
		wg.Done()
	}()

	wg.Wait()

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
