package pets

import (
	"errors"
	"fmt"

	"github.com/ildx/breeders/models"
)

type AnimalInterface interface {
	Show() string
}

type DogFromFactory struct {
	Pet *models.Dog
}

func (f *DogFromFactory) Show() string {
	return fmt.Sprintf("This animal is a %s", f.Pet.Breed.Breed)
}

type CatFromFactory struct {
	Pet *models.Cat
}

func (f *CatFromFactory) Show() string {
	return fmt.Sprintf("This animal is a %s", f.Pet.Breed.Breed)
}

type PetFactoryInterface interface {
	newPet() AnimalInterface
}

type DogAbstractFactory struct{}

func (f *DogAbstractFactory) newPet() AnimalInterface {
	return &DogFromFactory{
		Pet: &models.Dog{},
	}
}

type CatAbstractFactory struct{}

func (f *CatAbstractFactory) newPet() AnimalInterface {
	return &CatFromFactory{
		Pet: &models.Cat{},
	}
}

func NewPetFromAbstractFactory(species string) (AnimalInterface, error) {
	switch species {
	case "dog":
		var dogFactory DogAbstractFactory
		dog := dogFactory.newPet()
		return dog, nil
	case "cat":
		var catFactory CatAbstractFactory
		cat := catFactory.newPet()
		return cat, nil
  default:
    return nil, errors.New("Invalid species supplied.")
	}
}
