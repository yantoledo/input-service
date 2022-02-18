package main

import (
	"fmt"

	"github.com/yantoledo/input-service/api/routes.go"
)

func main() {

	fmt.Println("Initialiazing server..")
	routes.HandleRequest()
}
