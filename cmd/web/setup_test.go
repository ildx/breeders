package main

import (
	"os"
	"testing"

	"github.com/ildx/breeders/adapters"
	"github.com/ildx/breeders/config"
)

var testApp application

func TestMain(m *testing.M) {
	testBackend := &adapters.TestBackend{}
	testAdapter := &adapters.RemoteService{Remote: testBackend}

	testApp = application{
		App: config.New(nil, testAdapter),
	}

	os.Exit(m.Run())
}
