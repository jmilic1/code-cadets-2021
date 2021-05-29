package bootstrap

import (
	"code-cadets-2021/lecture_2/06_offerfeed/internal/domain/models"
	"code-cadets-2021/lecture_2/06_offerfeed/internal/infrastructure/queue"
)

func OrderedQueue(source chan models.Odd) *queue.OrderedQueue {
	return queue.NewOrderedQueue(source)
}
