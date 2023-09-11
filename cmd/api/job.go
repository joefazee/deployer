package main

import (
	"encoding/json"

	"github.com/hibiken/asynq"
	"github.com/joefazee/autodeploy/pkg/domain"
)

func (srv *server) sendJob(name string, payload any, opts ...asynq.Option) *domain.AppError {
	taskPayload, err := json.Marshal(payload)
	if err != nil {
		return domain.NewAppError("failed to marshal task payload", err)
	}

	task := asynq.NewTask(name, taskPayload, opts...)

	_, err = srv.asyncClient.Enqueue(task, opts...)
	if err != nil {
		return domain.NewAppError("unable to enqeueu task", err)
	}

	return nil
}
