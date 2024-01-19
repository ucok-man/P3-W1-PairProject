package main

import (
	"log"

	"github.com/ucok-man/P3-W1-PairProject/api"
	_ "github.com/ucok-man/P3-W1-PairProject/docs"
	"github.com/ucok-man/P3-W1-PairProject/internal/scheduler"
)

// @title Transaction API
// @version 1.0
// @description Documentation for Transaction API
// @termsOfService http://swagger.io/terms/

// @contact.name ucok-man
// @contact.email ucokkocu411@gmail.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host pixelrental-production.up.railway.app:8080
// @BasePath /v1
func main() {
	go func() {
		scheduler.StartScheduler()
	}()
	if err := api.New().Serve(); err != nil {
		log.Fatal(err)
	}
}
