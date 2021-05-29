package main

import (
	"log"

	"github.com/superbet-group/code-cadets-2021/homework_4/betsapi/cmd/bootstrap"
	"github.com/superbet-group/code-cadets-2021/homework_4/betsapi/cmd/config"
	"github.com/superbet-group/code-cadets-2021/homework_4/betsapi/internal/tasks"
)

func main() {
	log.Println("Bootstrap initiated")

	config.Load()

	signalHandler := bootstrap.SignalHandler()
	db := bootstrap.Sqlite()
	api := bootstrap.Api(db)

	log.Println("Bootstrap finished. Event API is starting")

	tasks.RunTasks(signalHandler, api)

	log.Println("Event API finished gracefully")
}
