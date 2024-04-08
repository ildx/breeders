package main

import (
	"os"
	"testing"

	"github.com/ildx/breeders/config"
	"github.com/ildx/breeders/models"
)

var testApp application

func TestMain(m *testing.M) {
	testBackend := &TestBackend{}
	testAdapter := &RemoteService{Remote: testBackend}

	testApp = application{
		App:        config.New(nil),
		catService: testAdapter,
	}

	os.Exit(m.Run())
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
