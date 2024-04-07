package main

import (
	"os"
	"testing"

	"github.com/ildx/breeders/models"
)

var testApp application

func TestMain(m *testing.M) {
	testApp = application{
		Models: *models.New(nil),
	}
	os.Exit(m.Run())
}
