package bootstrap

import (
	"github.com/superbet-group/code-cadets-2021/homework_4/betsapi/cmd/config"
	"github.com/superbet-group/code-cadets-2021/homework_4/betsapi/internal/api"
	"github.com/superbet-group/code-cadets-2021/homework_4/betsapi/internal/api/controllers"
	"github.com/superbet-group/code-cadets-2021/homework_4/betsapi/internal/domain/mappers"
	"github.com/superbet-group/code-cadets-2021/homework_4/betsapi/internal/domain/services"
	"github.com/superbet-group/code-cadets-2021/homework_4/betsapi/internal/infrastructure/sqlite"
)

func newBetMapper() *mappers.BetMapper {
	return mappers.NewBetMapper()
}

func newBetRepository(dbExecutor sqlite.DatabaseExecutor, betMapper sqlite.BetMapper) *sqlite.BetRepository {
	return sqlite.NewBetRepository(dbExecutor, betMapper)
}

func newBetService(betRepository services.BetRepository, betMapper services.BetMapper) *services.BetService {
	return services.NewBetService(betRepository, betMapper)
}

func newController(betService controllers.BetService) *controllers.Controller {
	return controllers.NewController(betService)
}

// Api bootstraps the http server.
func Api(dbExecutor sqlite.DatabaseExecutor) *api.WebServer {
	betMapper := newBetMapper()
	betRepository := newBetRepository(dbExecutor, betMapper)
	betService := newBetService(betRepository, betMapper)
	controller := newController(betService)

	return api.NewServer(config.Cfg.Api.Port, config.Cfg.Api.ReadWriteTimeoutMs, controller)
}
