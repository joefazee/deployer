package main

import (
	"database/sql"
	"expvar"
	"log"
	"os"

	"github.com/hibiken/asynq"
	"github.com/joefazee/autodeploy/pkg/config"
	"github.com/joefazee/autodeploy/pkg/db"

	_ "github.com/lib/pq"
)

// build time variables
var (
	version string
)

func main() {

	expvar.NewString("version").Set(version)

	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	// load and bind config
	cfg, err := config.Load(cwd)
	if err != nil {
		log.Fatal(err)
	}

	// connect to the database
	conn, err := sql.Open(cfg.DBDriver, cfg.DBSource)
	if err != nil {
		log.Fatal(err)
	}

	err = conn.Ping()
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		err := conn.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	// initialzie the server
	server := newServer(db.NewDBStore(conn), cfg)
	if server == nil {
		log.Fatal("server cannot be nil")
	}

	// configure redis client
	redisOpt := asynq.RedisClientOpt{
		Addr:     cfg.RedisAddress,
		Username: cfg.RedisUsername,
		Password: cfg.RedisPassword,
		DB:       cfg.RedisDB,
	}

	server.asyncServer = asynq.NewServer(redisOpt, asynq.Config{
		Concurrency: 10,
	})

	server.asyncMux = asynq.NewServeMux()
	server.asyncClient = asynq.NewClient(redisOpt)

	// run the async server
	go func() {
		if err := server.asyncServer.Run(server.asyncMux); err != nil {
			log.Fatal(err)
		}
	}()

	// start the server
	if err := server.run(cfg.HttServerAddress); err != nil {
		log.Fatal(err)
	}

}
