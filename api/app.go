package api

import (
	"log"

	"github.com/yantoledo/input-service/api/route"
)

type App struct {
}

func NewApp() *App {
	return &App{}
}

func (a App) Run() {

	log.Println("Initialiazing server..")
	route.HandleRequest()
}
