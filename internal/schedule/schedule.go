package schedule

import (
	"api_shop/config"
	"api_shop/internal/tasks"
	"api_shop/internal/worker"
	"log"

	"github.com/goccy/go-json"
	"github.com/hibiken/asynq"
)

func StartScheduleServer() {
	config.Load("worker")

	scheduler := asynq.NewScheduler(
		asynq.RedisClientOpt{Addr: worker.RedisAddr},
		nil,
	)

	payload := tasks.EmailDeliveryPayload{
		ID:    42,
		Email: "abbasmomeny1994@gmail.com",
		Text:  "welcome-email",
	}

	jm, _ := json.Marshal(payload)

	_, err := scheduler.Register(
		"* * * * *", // every minute
		asynq.NewTask(tasks.TypeEmailDeliveryTask, jm),
	)

	if err != nil {
		log.Fatalf("could not register task: %v", err)
	}

	if err := scheduler.Run(); err != nil {
		log.Fatalf("scheduler failed: %v", err)
	}

}
