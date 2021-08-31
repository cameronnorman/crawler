package main

import (
	"crawler/internal"
	"os"

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
		},
		Action: run,
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func run(c *cli.Context) error {
	// err := sentry.Init(sentry.ClientOptions{
	// 	Dsn:         c.String("sentry_dsn"),
	// 	Environment: c.String("app_env"),
	// 	Debug:       false,
	// })
	// if err != nil {
	// 	log.Fatalf("Failed to initialize sentry: %s", err)
	// }
	// defer sentry.Flush(2 * time.Second)

	log.SetFormatter(&log.JSONFormatter{})
	logger := log.New()

	// tracer.Start(
	// 	tracer.WithService(os.Getenv("DD_SERVICE_NAME")),
	// )
	// defer tracer.Stop()

	host := c.String("database_host")
	port := c.String("database_port")
	user := c.String("database_user")
	password := c.String("database_password")
	DBName := c.String("database_name")
	db, err := internal.NewDBConnection(host, port, user, password, DBName, logger)
	if err != nil {
		return err
	}

	app := internal.NewApp(db, true)
	return app.Run()
}
