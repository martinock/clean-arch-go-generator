package generator

import (
	"github.com/urfave/cli"
)

// Usecase contract functions for generator
type Usecase interface {
	GenerateModel(ctx *cli.Context) error
	GenerateRepository(ctx *cli.Context) error
	GenerateUsecase(ctx *cli.Context) error
	GenerateDelivery(ctx *cli.Context) error
}
