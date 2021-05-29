package bootstrap

import "github.com/superbet-group/code-cadets-2021/homework_4/betsapi/internal/tasks"

// SignalHandler bootstraps the signal handler.
func SignalHandler() *tasks.SignalHandler {
	return tasks.NewSignalHandler()
}
