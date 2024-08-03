package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/hibiken/asynq"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/mustafayilmazdev/musarchive/api"
	db "github.com/mustafayilmazdev/musarchive/db/sqlc"
	localization "github.com/mustafayilmazdev/musarchive/locales"
	"github.com/mustafayilmazdev/musarchive/mail"
	"github.com/mustafayilmazdev/musarchive/worker"

	"golang.org/x/sync/errgroup"

	"github.com/mustafayilmazdev/musarchive/util"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

var interruptSignals = []os.Signal{
	os.Interrupt,
	syscall.SIGTERM,
	syscall.SIGINT,
}

func main() {

	config, err := util.LoadConfig(".")

	if err != nil {
		log.Fatal().Err(err).Msg("cannot load config: ")

	}

	if config.Environment == "development" {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	}

	// Initialize the LocalizationManager singleton
	if err := localization.Initialize(util.LocalizationPath + util.DefaultLocale + util.LocalizationType); err != nil {
		log.Fatal().Msg("Can not load localization")
	}

	lm := localization.GetInstance()

	ctx, stop := signal.NotifyContext(context.Background(), interruptSignals...)
	defer stop()

	connPool, err := pgxpool.New(ctx, config.DBSource)
	if err != nil {
		log.Fatal().Err(err).Msg(lm.Translate(util.DefaultLocale, localization.Errors_CanNotConnectToDb, err))
	}
	runDBMigration(config.MigrationUrl, config.DBSource)
	store := db.NewStore(connPool)
	waitGroup, ctx := errgroup.WithContext(ctx)

	redisOpt := asynq.RedisClientOpt{
		Addr: config.RedisAddress,
	}

	taskDistributor := worker.NewRedisTaskDistributor(redisOpt)
	go runTaskProcessor(config, ctx, waitGroup, redisOpt, store)
	runGinServer(config, ctx, waitGroup, store, taskDistributor)

	err = waitGroup.Wait()
	if err != nil {
		log.Fatal().Err(err).Msg(lm.Translate(util.DefaultLocale, localization.Errors_ErrorFromWaitGroup, err))
	}
}

func runDBMigration(migrationUrl, dbSource string) {
	lm := localization.GetInstance()
	migration, err := migrate.New(migrationUrl, dbSource)
	if err != nil {
		log.Fatal().Err(err).Msg(lm.Translate(util.DefaultLocale, localization.Errors_MigrationInstance))
	}
	if err = migration.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal().Err(err).Msg(lm.Translate(util.DefaultLocale, localization.Errors_MigrateUp))
	}
	log.Info().Msg(lm.Translate(util.DefaultLocale, localization.Success_Migrate))
}

func runTaskProcessor(config util.Config, ctx context.Context, waitGroup *errgroup.Group, redisOpt asynq.RedisClientOpt, store db.Store) {
	mailer := mail.NewGmailSender(config.EmailSenderName, config.EmailSenderAddress, config.EmailSenderPassword)
	taskProcessor := worker.NewRedisTaskProcessor(redisOpt, store, mailer)
	log.Info().Msg("start task processor")
	err := taskProcessor.Start()
	if err != nil {
		log.Fatal().Err(err).Msg("failed to start task processor")
	}

	waitGroup.Go(func() error {
		<-ctx.Done()
		log.Info().Msg("graceful shutdown task processor")

		taskProcessor.ShutDown()
		log.Info().Msg("task processor is stopped")
		return nil
	})
}

func runGinServer(config util.Config, ctx context.Context, waitGroup *errgroup.Group, store db.Store, taskDistributor worker.TaskDistributor) {

	server, err := api.NewServer(config, store, taskDistributor)
	if err != nil {
		fmt.Println(err)
	}

	// Create an http.Server instance
	httpServer := &http.Server{
		Addr:    config.HTTPServerAddress,
		Handler: server.Router,
	}

	waitGroup.Go(func() error {
		lm := localization.GetInstance()
		log.Info().Msgf(lm.Translate(util.DefaultLocale, localization.Info_StartHttp, httpServer.Addr))
		err = httpServer.ListenAndServe()
		if err != nil {
			if errors.Is(err, http.ErrServerClosed) {
				return nil
			}
			log.Fatal().Err(err).Msg(lm.Translate(util.DefaultLocale, localization.Errors_HttpGateway))
			return err
		}
		return nil
	})

	waitGroup.Go(func() error {
		lm := localization.GetInstance()
		<-ctx.Done()
		log.Info().Msg(lm.Translate("tr", localization.Info_StartHttp, httpServer.Addr))
		err := httpServer.Shutdown(context.Background())
		if err != nil {
			log.Error().Err(err).Msg(lm.Translate(util.DefaultLocale, localization.Errors_HttpGatewayShutdown))
			return err
		}
		log.Info().Msg(lm.Translate(util.DefaultLocale, localization.Info_StopHttp))
		return nil
	})

}
