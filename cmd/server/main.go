package main

import (
	cfg "github.com/StepanShevelev/library/config"
	mydb "github.com/StepanShevelev/library/db"
	protobuff "github.com/StepanShevelev/library/internal/api/protobuff"
	"log"
	"os"
	"os/signal"
)

func main() {

	config := cfg.New()
	if err := config.Load("./configs", "config", "yml"); err != nil {
		log.Fatal(err)
	}

	_, err := mydb.FillDbWithTestData(config)
	if err != nil {
		log.Fatal(err)
	}

	grpcSrv := protobuff.NewServer(config)
	protobuff.RunServer(grpcSrv, config)

	c := make(chan os.Signal, 1)
	// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	// SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
	signal.Notify(c, os.Interrupt)

	// Block until we receive our signal
	<-c

	//api.Shutdown()
	grpcSrv.Stop()

	//log.Info("Good bye!")
	os.Exit(0)
}
