package bootstrap

import "code-cadets-2021/homework_2/offerfeed/internal/infrastructure/queue"

func NewOrderedQueue() *queue.OrderedQueue {
	return queue.NewOrderedQueue()
}
