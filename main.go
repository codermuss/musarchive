package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/golang-migrate/migrate"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/mustafayilmazdev/musarchive/api"
	db "github.com/mustafayilmazdev/musarchive/db/sqlc"

	"github.com/mustafayilmazdev/musarchive/util"
)

var interruptSignals = []os.Signal{
	os.Interrupt,
	syscall.SIGTERM,
	syscall.SIGINT,
}

func main() {
	config, err := util.LoadConfig(".")

	if err != nil {
		fmt.Println(err)

	}

	ctx, stop := signal.NotifyContext(context.Background(), interruptSignals...)
	defer stop()

	connPool, err := pgxpool.New(ctx, config.DBSource)
	if err != nil {
		fmt.Println(err)
	}

	store := db.NewStore(connPool)

	runGinServer(config, store)
}

func runDBMigration(migrationUrl, dbSource string) {
	migration, err := migrate.New(migrationUrl, dbSource)
	if err != nil {
		fmt.Println(err)
	}
	if err = migration.Up(); err != nil && err != migrate.ErrNoChange {
		fmt.Println(err)
	}
}

func runGinServer(config util.Config, store db.Store) {

	server, err := api.NewServer(config, store)
	if err != nil {
		fmt.Println(err)
	}

	err = server.Start(config.HTTPServerAddress)
	if err != nil {
		fmt.Println(err)
	}
}
