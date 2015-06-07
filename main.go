package main

import (
	"log"
	"os"
	"path/filepath"

	sql "github.com/aodin/aspect"
	_ "github.com/aodin/aspect/postgres"
	"github.com/aodin/volta/config"
	"github.com/codegangsta/cli"
	"github.com/pgrunde/just-auth/server"
)

func main() {
	app := cli.NewApp()
	app.Name = "auth"
	app.Usage = "Just authentication"
	app.Action = startServer
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "log, l",
			Value: "",
			Usage: "Sets the log output file path",
		},
		cli.StringFlag{
			Name:  "config, c",
			Value: "./settings.json",
			Usage: "Sets config file",
		},
	}
	app.Run(os.Args)
}

func startServer(c *cli.Context) {
	logFP := c.String("log")
	file := c.String("config")
	if logFP != "" {
		if err := os.MkdirAll(filepath.Dir(logFP), 0776); err != nil {
			log.Panic(err)
		}
		lg, err := os.OpenFile(logFP, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
		if err != nil {
			log.Panic(err)
		}
		defer lg.Close()
		log.SetOutput(lg)
	}
	conf, conn := loadSettings(file)
	log.Fatal(server.New(conf, conn).ListenAndServe())
}

func loadSettings(file string) (config.Config, *sql.DB) {
	conf, err := config.ParseFile(file)
	if err != nil {
		log.Panicf("Could not parse configuration file: %s", err)
	}

	conn, err := sql.Connect(conf.Database.Driver, conf.Database.Credentials())
	if err != nil {
		log.Panicf("Could not connect to database: %s", err)
	}
	return conf, conn
}
