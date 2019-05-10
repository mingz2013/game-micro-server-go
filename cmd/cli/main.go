package main

import (
	"github.com/mingz2013/game-micro-server-go/internal/app/cli/actions/robot"
	"github.com/urfave/cli"
	"log"
	"os"
)

func main() {

	app := cli.NewApp()

	app.Commands = []cli.Command{

		{
			Name:  "robot",
			Usage: "robot manager",
			Subcommands: []cli.Command{
				{
					Name:  "start",
					Usage: "start robot actions...",
					Action: func(c *cli.Context) error {

						log.Println("robot start..")
						return robot.Start()
						//return nil
					},
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

}
