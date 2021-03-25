package commands

import (
	"github.com/solomon04/go-docker/database/seeds"
	"github.com/solomon04/go-docker/src/db"
	"github.com/urfave/cli/v2"
	"os"
)

var SeederCommand = cli.Command{
	Name:  "db:seed",
	Usage: "Seed the database",
	Flags: []cli.Flag{
		&cli.StringFlag{Name: "seeder", Aliases: []string{"s"}, Usage: "Specify the database seeder to run (e.g. UserSeed)"},
	},
	Action: func(c *cli.Context) error {
		return handle(c)
	},
}

func handle(c *cli.Context) error {
	seeder := c.String("seeder")

	if seeder != "" {
		println(seeder)
		seeds.Execute(db.Db,seeder)
		os.Exit(0)
	}

	return nil
}


