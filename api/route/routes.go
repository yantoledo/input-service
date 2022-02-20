package route

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/yantoledo/input-service/api/controller"
)

func HandleRequest() {
	r := mux.NewRouter()
	r.HandleFunc("/", controller.Ping)

	facebookRoutes(r)

	log.Fatal(http.ListenAndServe(":5003", r))
}

func facebookRoutes(r *mux.Router) {
	r.HandleFunc("/api/facebook", controller.MessageEventFromFacebook).Methods("Post")
	r.HandleFunc("/api/facebook", controller.FacebookAuthorization).Methods("Get")
}
