package main

import (
	"os"
	"testing"

	"github.com/rogpeppe/go-internal/testscript"
)

func TestMain(m *testing.M) {
	exitVal := testscript.RunMain(m, map[string]func() int{
		"greta": Main,
	})

	os.Exit(exitVal)
}

func TestGreta(t *testing.T) {
	testscript.Run(t, testscript.Params{
		Dir: "./testdata",
	})
}
