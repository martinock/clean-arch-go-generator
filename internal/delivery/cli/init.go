package cli

import (
	"log"

	"github.com/martinock/gocleanarch/internal/usecase/generator"
	"github.com/urfave/cli"
)

// Initialize init cli app with args input from user
func Initialize(args []string, usecase generator.Usecase) {
	app := cli.NewApp()
	app.Name = "gocleanarch"
	app.Usage = "Generate Golang clean arch boilerplate for you"
	app.Commands = []cli.Command{
		{
			Name:    "model",
			Aliases: []string{"m"},
			Usage:   "Create a model folder",
			Action:  usecase.GenerateModel,
		},
		{
			Name:    "repo",
			Aliases: []string{"r"},
			Usage:   "Create a repository folder",
			Action:  usecase.GenerateRepository,
		},
		{
			Name:    "usecase",
			Aliases: []string{"u"},
			Usage:   "Create a usecase folder",
			Action:  usecase.GenerateUsecase,
		},
		{
			Name:    "delivery",
			Aliases: []string{"d"},
			Usage:   "Create a delivery folder",
			Action:  usecase.GenerateDelivery,
		},
	}

	err := app.Run(args)
	if err != nil {
		log.Fatal(err)
	}
}
