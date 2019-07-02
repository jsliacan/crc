// +build integration

/*
Copyright (C) 2019 Red Hat, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package integration

import (
	"os"
	"strings"
	"testing"

	"github.com/DATA-DOG/godog"
	"github.com/code-ready/clicumber/testsuite"
	"github.com/code-ready/crc/test/integration/crcsuite"
)

func TestMain(m *testing.M) {

	parseFlags()

	status := godog.RunWithOptions("crc", func(s *godog.Suite) {
		getFeatureContext(s)
	}, godog.Options{
		Format:              testsuite.GodogFormat,
		Paths:               strings.Split(testsuite.GodogPaths, ","),
		Tags:                testsuite.GodogTags,
		ShowStepDefinitions: testsuite.GodogShowStepDefinitions,
		StopOnFailure:       testsuite.GodogStopOnFailure,
		NoColors:            testsuite.GodogNoColors,
	})

	os.Exit(status)
}

func getFeatureContext(s *godog.Suite) {

	// NOTE:
	// Preserve the order in which FeatureContexts are loaded.

	testsuite.FeatureContext(s) // Clicumber step definitions

	crcsuite.FeatureContext(s) // CRC specific step definitions

}

func parseFlags() {

	// NOTE:
	// testsuite.ParseFlags() needs to be last - it calls flag.Parse()
	crcsuite.ParseFlags()
	testsuite.ParseFlags()

}
