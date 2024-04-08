package adapters

import (
	"encoding/json"
	"encoding/xml"
	"io"
	"net/http"

	"github.com/ildx/breeders/models"
)

type CatBreedsInterface interface {
	GetAllCatBreeds() ([]*models.CatBreed, error)
	GetCatBreedByName(b string) (*models.CatBreed, error)
}

type RemoteService struct {
	Remote CatBreedsInterface
}

func (rs *RemoteService) GetAllBreeds() ([]*models.CatBreed, error) {
	return rs.Remote.GetAllCatBreeds()
}

type JSONBackend struct{}

func (j *JSONBackend) GetAllCatBreeds() ([]*models.CatBreed, error) {
	res, err := http.Get("http://localhost:8081/api/cat-breeds/all/json")
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var breeds []*models.CatBreed
	err = json.Unmarshal(body, &breeds)
	if err != nil {
		return nil, err
	}

	return breeds, nil
}

func (j *JSONBackend) GetCatBreedByName(b string) (*models.CatBreed, error) {
	res, err := http.Get("http://localhost:8081/api/cat-breeds/" + b + "/json")
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var breed models.CatBreed
	err = json.Unmarshal(body, &breed)
	if err != nil {
		return nil, err
	}

	return &breed, nil
}

type XMLBackend struct{}

func (x *XMLBackend) GetAllCatBreeds() ([]*models.CatBreed, error) {
	res, err := http.Get("http://localhost:8081/api/cat-breeds/all/xml")
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	type catBreeds struct {
		XMLName struct{}           `xml:"cat-breeds"`
		Breeds  []*models.CatBreed `xml:"cat-breed"`
	}

	var breeds catBreeds
	err = xml.Unmarshal(body, &breeds)
	if err != nil {
		return nil, err
	}

	return breeds.Breeds, nil
}

func (x *XMLBackend) GetCatBreedByName(b string) (*models.CatBreed, error) {
	res, err := http.Get("http://localhost:8081/api/cat-breeds/" + b + "/xml")
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var breed models.CatBreed
	err = xml.Unmarshal(body, &breed)
	if err != nil {
		return nil, err
	}

	return &breed, nil
}

type TestBackend struct{}

func (b *TestBackend) GetAllCatBreeds() ([]*models.CatBreed, error) {
	breeds := []*models.CatBreed{
		{
			ID:      1,
			Breed:   "Tomcat",
			Details: "Some details",
		},
	}
	return breeds, nil
}

func (b *TestBackend) GetCatBreedByName(n string) (*models.CatBreed, error) {
	return nil, nil
}
