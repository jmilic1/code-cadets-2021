package main

import (
	"fmt"
	"log"

	"code-cadets-2021/lecture_2/06_offerfeed/cmd/bootstrap"
	"code-cadets-2021/lecture_2/06_offerfeed/internal/domain/models"
	"code-cadets-2021/lecture_2/06_offerfeed/internal/tasks"
)

func main() {
	httpClient := bootstrap.HttpClient()
	updatesChannel := make(chan models.Odd)
	queueChannel := make(chan models.Odd)

	offerFeed := bootstrap.AxilisOfferFeed(httpClient, updatesChannel)
	homeworkFeed := bootstrap.HomeworkOfferFeed(httpClient, updatesChannel)
	feedComponent, err := bootstrap.FeedComponent(updatesChannel, offerFeed, homeworkFeed)
	if err != nil {
		log.Fatalln(err, "error while constructing OfferFeedComponent")
	}

	queue := bootstrap.OrderedQueue(queueChannel)

	service := bootstrap.FeedProcessingService(feedComponent, queue)

	signalHandler := bootstrap.SignalHandler()

	tasks.RunTasks(signalHandler, feedComponent, service, queue)
	fmt.Println("program finished gracefully")
}
