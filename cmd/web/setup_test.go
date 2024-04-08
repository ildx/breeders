package main

import (
	"os"
	"testing"

	"github.com/ildx/breeders/config"
)

var testApp application

func TestMain(m *testing.M) {
	testApp = application{
		App: config.New(nil),
	}
	os.Exit(m.Run())
}
