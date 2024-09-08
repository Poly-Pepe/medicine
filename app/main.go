package main

import (
	"context"
	"fmt"
	"runtime"

	"medicine/config"
	"medicine/examination/injection"

	"github.com/jackc/pgx/v5/pgxpool"
	log "github.com/sirupsen/logrus"
)

func initLogger(logLevel string) error {
	level, err := log.ParseLevel(logLevel)
	if err != nil {
		return err
	}

	log.SetLevel(level)

	log.SetReportCaller(true)

	log.SetFormatter(&log.JSONFormatter{
		CallerPrettyfier: func(frame *runtime.Frame) (function string, file string) {
			return fmt.Sprintf("%s:%d", frame.Function, frame.Line), ""
		},
	})
	return nil
}

func main() {
	ctx := context.Background()

	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatal(err)
	}

	if err := initLogger(cfg.LogLevel); err != nil {
		log.Fatal(err)
	}

	pool, err := pgxpool.New(ctx, cfg.DataBaseDNS)
	if err != nil {
		log.Fatal(err)
	}

	examination := injection.InitializeExaminationFront(pool)

	log.Fatal(examination.App.Listen(fmt.Sprintf(":%d", cfg.AppPort)))
}
