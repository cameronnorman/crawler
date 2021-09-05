package main

import (
	"context"
	"crawler/internal"
	"crawler/internal/workers"
	"os"

	worker "github.com/contribsys/faktory_worker_go"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "crawler-api",
		Usage: "used to crawl websites and download their raw HTML",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "database_host",
				EnvVars: []string{"DATABASE_HOST"},
				Usage:   "The URL to use for connecting to the database",
			},
			&cli.StringFlag{
				Name:    "database_port",
				EnvVars: []string{"DATABASE_PORT"},
				Usage:   "The URL to use for connecting to the database",
			},
			&cli.StringFlag{
				Name:    "database_user",
				EnvVars: []string{"DATABASE_USER"},
				Usage:   "The URL to use for connecting to the database",
			},
			&cli.StringFlag{
				Name:    "database_password",
				EnvVars: []string{"DATABASE_PASSWORD"},
				Usage:   "The URL to use for connecting to the database",
			},
			&cli.StringFlag{
				Name:    "database_name",
				EnvVars: []string{"DATABASE_DB"},
				Usage:   "The URL to use for connecting to the database",
			},
			&cli.StringFlag{
				Name:    "app_env",
				EnvVars: []string{"APP_ENV"},
				Usage:   "The environment the application is running in",
			},
			&cli.StringFlag{
				Name:     "worker_function",
				Usage:    "The worker function you wish to run",
				Required: true,
			},
		},
		Action: run,
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func run(c *cli.Context) error {
	log.SetFormatter(&log.JSONFormatter{})
	logger := log.New()

	host := c.String("database_host")
	port := c.String("database_port")
	user := c.String("database_user")
	password := c.String("database_password")
	DBName := c.String("database_name")
	_, err := internal.NewDBConnection(host, port, user, password, DBName, logger)
	if err != nil {
		return err
	}

	workerFunctions := map[string]interface{}{
		"get_html_worker": workers.GetHtmlWorker,
	}

	f := workerFunctions[c.String("worker_function")]
	mgr := worker.NewManager()

	convertedFunction := f.(func(context.Context, ...interface{}) error)
	mgr.Register(c.String("worker_function"), convertedFunction)

	mgr.Concurrency = 20

	mgr.Run()
	return nil
}
