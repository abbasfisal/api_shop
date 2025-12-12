package worker

import (
	"api_shop/config"
	"api_shop/internal/tasks"
	"fmt"
	"log"

	"github.com/hibiken/asynq"
)

// todo: get from config
const RedisAddr = "127.0.0.1:6379"

func StartWorker() {

	config.Load("worker")

	srv := asynq.NewServer(
		asynq.RedisClientOpt{Addr: RedisAddr},
		asynq.Config{
			// Specify how many concurrent workers to use
			Concurrency: 10,
			// Optionally specify multiple queues with different priority.
			Queues: map[string]int{
				"critical": 6,
				"default":  3,
				"low":      1,
			},
		},
	)

	// mux maps a type to a handler
	mux := asynq.NewServeMux()

	mux.HandleFunc(tasks.TypeEmailDeliveryTask, tasks.HandleEmailDeliveryTask)

	fmt.Println("worker started ")
	if err := srv.Run(mux); err != nil {
		log.Fatalf("could not run server: %v", err)
	}

}
