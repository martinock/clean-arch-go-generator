package generator

type generatorUsecase struct{}

// NewGeneratorUsecase a function to create a new generator usecase
func NewGeneratorUsecase() Usecase {
	return generatorUsecase{}
}
