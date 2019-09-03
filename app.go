package main

import (
	"os"

	"github.com/martinock/gocleanarch/internal/delivery/cli"
	"github.com/martinock/gocleanarch/internal/usecase/generator"
)

func main() {
	generatorUsecase := generator.NewGeneratorUsecase()
	cli.Initialize(os.Args, generatorUsecase)
}
