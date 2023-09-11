package main

import (
	"context"
	"log"

	"github.com/hibiken/asynq"
	"github.com/joefazee/autodeploy/pkg/tasks"
)

func (srv *server) registerAsyncHandlers() {

	srv.asyncMux.HandleFunc(tasks.TaskUserCreated, func(ctx context.Context, t *asynq.Task) error {

		log.Printf("Processing task: %s with payload %s\n", t.Type(), t.Payload())

		log.Println("done with the job!!!")

		return nil
	})
}
