package main

import (
	"log"

	"github.com/superbet-group/code-cadets-2021/homework_4/betacceptance/cmd/bootstrap"
	"github.com/superbet-group/code-cadets-2021/homework_4/betacceptance/cmd/config"
	"github.com/superbet-group/code-cadets-2021/homework_4/betacceptance/internal/tasks"
)

func main() {
	log.Println("Bootstrap initiated")

	config.Load()

	rabbitMqChannel := bootstrap.RabbitMq()
	signalHandler := bootstrap.SignalHandler()
	api := bootstrap.Api(rabbitMqChannel)

	log.Println("Bootstrap finished. Bet Acceptance API is starting")

	tasks.RunTasks(signalHandler, api)

	log.Println("Bet Acceptance API finished gracefully")
}
