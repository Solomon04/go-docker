package main

import (
	"flag"
	_ "github.com/go-sql-driver/mysql"
	"github.com/solomon04/go-docker/bootstrap"
	"github.com/solomon04/go-docker/database/seeds"
	"github.com/solomon04/go-docker/src/console/commands"
	"github.com/solomon04/go-docker/src/db"
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

func main() {
	bootstrap.Boot()
	app := &cli.App{
		Name: "HoopSpots",
		Usage: "Run console commands for the API",
		Commands: []*cli.Command{
			&commands.SeederCommand,
			// Live Look Command
			// Recommendation Command
			// Auto Populate Command
			// Generate Auth Token Command
			// Database Commands: Fresh, Seed, Migrate
			// Cron Scheduler Command
			// https://github.com/urfave/cli/blob/master/docs/v2/manual.md#arguments
			// https://medium.com/easyread/how-i-seed-my-database-with-go-27488d2e6a75
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func handleArgs() {
	flag.Parse()
	args := flag.Args()

	if len(args) >= 1 {
		switch args[0] {
		case "seed":
			seeds.Execute(db.Db, args[1:]...)
			os.Exit(0)
		}
	}
}
