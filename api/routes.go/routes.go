package routes

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/yantoledo/input-service/api/controllers"
)

func HandleRequest() {
	r := mux.NewRouter()
	r.HandleFunc("/", controllers.Ping)
	r.HandleFunc("/api/facebook", controllers.MessageEventFromFacebook).Methods("Post")
	r.HandleFunc("/api/facebook", controllers.FacebookAuthorization).Methods("Get")
	log.Fatal(http.ListenAndServe(":5003", r))
}
