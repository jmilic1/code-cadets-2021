package bootstrap

import (
	"github.com/superbet-group/code-cadets-2021/lecture_3/03_project/calculator/cmd/config"
	"github.com/superbet-group/code-cadets-2021/lecture_3/03_project/calculator/internal/domain/mappers"
	"github.com/superbet-group/code-cadets-2021/lecture_3/03_project/calculator/internal/engine"
	"github.com/superbet-group/code-cadets-2021/lecture_3/03_project/calculator/internal/engine/consumer"
	"github.com/superbet-group/code-cadets-2021/lecture_3/03_project/calculator/internal/engine/handler"
	"github.com/superbet-group/code-cadets-2021/lecture_3/03_project/calculator/internal/engine/publisher"
	"github.com/superbet-group/code-cadets-2021/lecture_3/03_project/calculator/internal/infrastructure/rabbitmq"
	"github.com/superbet-group/code-cadets-2021/lecture_3/03_project/calculator/internal/infrastructure/sqlite"
)

func newBetConsumer(channel rabbitmq.Channel) *rabbitmq.BetConsumer {
	betReceivedConsumer, err := rabbitmq.NewBetConsumer(
		channel,
		rabbitmq.ConsumerConfig{
			Queue:             config.Cfg.Rabbit.ConsumerBetQueue,
			DeclareDurable:    config.Cfg.Rabbit.DeclareDurable,
			DeclareAutoDelete: config.Cfg.Rabbit.DeclareAutoDelete,
			DeclareExclusive:  config.Cfg.Rabbit.DeclareExclusive,
			DeclareNoWait:     config.Cfg.Rabbit.DeclareNoWait,
			DeclareArgs:       nil,
			ConsumerName:      config.Cfg.Rabbit.ConsumerBetReceivedName,
			AutoAck:           config.Cfg.Rabbit.ConsumerAutoAck,
			Exclusive:         config.Cfg.Rabbit.ConsumerExclusive,
			NoLocal:           config.Cfg.Rabbit.ConsumerNoLocal,
			NoWait:            config.Cfg.Rabbit.ConsumerNoWait,
			Args:              nil,
		},
	)
	if err != nil {
		panic(err)
	}
	return betReceivedConsumer
}

func newEventSettledConsumer(channel rabbitmq.Channel) *rabbitmq.EventSettledConsumer {
	betCalculatedConsumer, err := rabbitmq.NewEventSettledConsumer(
		channel,
		rabbitmq.ConsumerConfig{
			Queue:             config.Cfg.Rabbit.ConsumerEventUpdatesQueue,
			DeclareDurable:    config.Cfg.Rabbit.DeclareDurable,
			DeclareAutoDelete: config.Cfg.Rabbit.DeclareAutoDelete,
			DeclareExclusive:  config.Cfg.Rabbit.DeclareExclusive,
			DeclareNoWait:     config.Cfg.Rabbit.DeclareNoWait,
			DeclareArgs:       nil,
			ConsumerName:      config.Cfg.Rabbit.ConsumerBetCalculatedName,
			AutoAck:           config.Cfg.Rabbit.ConsumerAutoAck,
			Exclusive:         config.Cfg.Rabbit.ConsumerExclusive,
			NoLocal:           config.Cfg.Rabbit.ConsumerNoLocal,
			NoWait:            config.Cfg.Rabbit.ConsumerNoWait,
			Args:              nil,
		},
	)
	if err != nil {
		panic(err)
	}
	return betCalculatedConsumer
}

func newConsumer(betReceivedConsumer consumer.BetConsumer, betCalculatedConsumer consumer.EventSettledConsumer) *consumer.Consumer {
	return consumer.New(betReceivedConsumer, betCalculatedConsumer)
}

func newBetMapper() *mappers.BetCalculatedMapper {
	return mappers.NewBetCalculatedMapper()
}

func newBetRepository(dbExecutor sqlite.DatabaseExecutor, betMapper sqlite.BetCalculatedMapper) *sqlite.BetCalculatedRepository {
	return sqlite.NewBetCalculatedRepository(dbExecutor, betMapper)
}

func newHandler(betRepository handler.BetCalculatedRepository) *handler.Handler {
	return handler.New(betRepository)
}

func newBetCalculatedPublisher(channel rabbitmq.Channel) *rabbitmq.BetCalculatedPublisher {
	betPublisher, err := rabbitmq.NewBetCalculatedPublisher(
		channel,
		rabbitmq.PublisherConfig{
			Queue:             config.Cfg.Rabbit.PublisherBetCalculatedQueue,
			DeclareDurable:    config.Cfg.Rabbit.DeclareDurable,
			DeclareAutoDelete: config.Cfg.Rabbit.DeclareAutoDelete,
			DeclareExclusive:  config.Cfg.Rabbit.DeclareExclusive,
			DeclareNoWait:     config.Cfg.Rabbit.DeclareNoWait,
			DeclareArgs:       nil,
			Exchange:          config.Cfg.Rabbit.PublisherExchange,
			Mandatory:         config.Cfg.Rabbit.PublisherMandatory,
			Immediate:         config.Cfg.Rabbit.PublisherImmediate,
		},
	)
	if err != nil {
		panic(err)
	}
	return betPublisher
}

func newPublisher(betPublisher publisher.BetCalculatedPublisher) *publisher.Publisher {
	return publisher.New(betPublisher)
}

func Engine(rabbitMqChannel rabbitmq.Channel, dbExecutor sqlite.DatabaseExecutor) *engine.Engine {
	betConsumer := newBetConsumer(rabbitMqChannel)
	eventSettledConsumer := newEventSettledConsumer(rabbitMqChannel)
	consumer := newConsumer(betConsumer, eventSettledConsumer)

	betMapper := newBetMapper()
	betRepository := newBetRepository(dbExecutor, betMapper)
	handler := newHandler(betRepository)

	betPublisher := newBetCalculatedPublisher(rabbitMqChannel)
	publisher := newPublisher(betPublisher)

	return engine.New(consumer, handler, publisher)
}
