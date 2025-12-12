package tasks

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/hibiken/asynq"
)

const TypeEmailDeliveryTask = "email:delivery"

type EmailDeliveryPayload struct {
	ID    int
	Email string
	Text  string
}

func HandleEmailDeliveryTask(c context.Context, t *asynq.Task) error {

	var p EmailDeliveryPayload

	err := json.Unmarshal(t.Payload(), &p)
	if err != nil {
		return fmt.Errorf("json.Unmarshal failed: %v: %w", err, asynq.SkipRetry)
	}

	log.Printf("Sending Email to User: user_email=%s, text=%s", p.Email, p.Text)

	return nil
}
