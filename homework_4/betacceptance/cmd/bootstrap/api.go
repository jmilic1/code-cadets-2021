package bootstrap

import (
	"github.com/streadway/amqp"

	"github.com/superbet-group/code-cadets-2021/homework_4/betacceptance/cmd/config"
	"github.com/superbet-group/code-cadets-2021/homework_4/betacceptance/internal/api"
	"github.com/superbet-group/code-cadets-2021/homework_4/betacceptance/internal/api/controllers"
	"github.com/superbet-group/code-cadets-2021/homework_4/betacceptance/internal/api/controllers/validators"
	"github.com/superbet-group/code-cadets-2021/homework_4/betacceptance/internal/domain/services"
	"github.com/superbet-group/code-cadets-2021/homework_4/betacceptance/internal/infrastructure"
	"github.com/superbet-group/code-cadets-2021/homework_4/betacceptance/internal/infrastructure/rabbitmq"
)

func newBetRequestValidator() *validators.BetRequestValidator {
	return validators.NewBetRequestValidator()
}

func newIdGenerator() *infrastructure.IdGenerator {
	return infrastructure.NewIdGenerator()
}

func newBetReceivedPublisher(publisher rabbitmq.QueuePublisher) *rabbitmq.BetReceivedPublisher {
	return rabbitmq.NewBetReceivedPublisher(
		config.Cfg.Rabbit.PublisherExchange,
		config.Cfg.Rabbit.PublisherBetReceivedQueue,
		config.Cfg.Rabbit.PublisherMandatory,
		config.Cfg.Rabbit.PublisherImmediate,
		publisher,
	)
}

func newBetService(publisher services.BetReceivedPublisher, idGenerator services.IdGenerator) *services.BetService {
	return services.NewBetService(publisher, idGenerator)
}

func newController(betRequestValidator controllers.BetRequestValidator, betService controllers.BetService) *controllers.Controller {
	return controllers.NewController(betRequestValidator, betService)
}

// Api bootstraps the http server.
func Api(rabbitMqChannel *amqp.Channel) *api.WebServer {
	betRequestValidator := newBetRequestValidator()
	betReceivedPublisher := newBetReceivedPublisher(rabbitMqChannel)
	idGenerator := newIdGenerator()

	eventService := newBetService(betReceivedPublisher, idGenerator)
	controller := newController(betRequestValidator, eventService)

	return api.NewServer(config.Cfg.Api.Port, config.Cfg.Api.ReadWriteTimeoutMs, controller)
}
