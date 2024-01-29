package main

import (
	"github.com/cucumber/godog"
	steps "github.com/tbrito/ms-lanchonetedarua-clientes/tests/bdd/steps"
	"testing"
)

func TestFeature(t *testing.T) {
	suite := godog.TestSuite{
		ScenarioInitializer: func(sc *godog.ScenarioContext) {
			steps.InitializeScenario(sc, t)
		},
		Options: &godog.Options{
			Format:   "pretty",
			Paths:    []string{"../../tests/bdd/features"},
			TestingT: t,
		},
	}

	if suite.Run() != 0 {
		t.Fatal("non-zero status returned, failed to run features tests")
	}
}
