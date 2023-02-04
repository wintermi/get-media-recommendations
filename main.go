// Copyright 2022-2023, Matthew Winter
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/rs/zerolog"
)

var logger zerolog.Logger
var applicationText = "%s 0.1.0%s"
var copyrightText = "Copyright 2022-2023, Matthew Winter\n"
var indent = "..."

var helpText = `
A command line application designed to provide a simple method to request
recommendations from a Google Cloud Discovery Engine API model for the
parameters contained within the input file.

Use --help for more details.


USAGE:
    get-media-recommendations -p PROJECT_NUMBER -s SERVING_CONFIG -i INPUT_FILE

ARGS:
`

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, applicationText, filepath.Base(os.Args[0]), "\n")
		fmt.Fprint(os.Stderr, copyrightText)
		fmt.Fprint(os.Stderr, helpText)
		flag.PrintDefaults()
	}

	// Define the Long CLI flag names
	var projectNumber = flag.String("p", "", "Project Number  (Required)")
	var location = flag.String("l", "global", "Location")
	var dataStore = flag.String("d", "default_data_store", "Data Store")
	var branch = flag.String("b", "0", "Branch")
	var servingConfig = flag.String("s", "", "Serving Config  (Required)")
	var parameterInputFile = flag.String("i", "", "Parameter Input File  (Required)")
	var numberResults = flag.Int("n", 5, "Number of Results, 1 to 100")
	var filterString = flag.String("f", "", "Filter String")
	var verbose = flag.Bool("v", false, "Output Verbose Detail")

	// Parse the flags
	flag.Parse()

	// Validate the Required Flags
	if *projectNumber == "" || *location == "" || *dataStore == "" || *branch == "" || *servingConfig == "" || *parameterInputFile == "" {
		flag.Usage()
		os.Exit(1)
	}

	// Verify Number of Results is between 1 and 100
	if *numberResults < 1 || *numberResults > 100 {
		flag.Usage()
		os.Exit(1)
	}

	// Setup Zero Log for Consolo Output
	output := zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC3339}
	logger = zerolog.New(output).With().Timestamp().Logger()
	zerolog.TimeFieldFormat = "2006-01-02 15:04:05.000"
	zerolog.DurationFieldUnit = time.Millisecond
	zerolog.DurationFieldInteger = true
	if *verbose {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	} else {
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	}

	// Output Header
	logger.Info().Msgf(applicationText, filepath.Base(os.Args[0]), "")
	logger.Info().Msg("Arguments")
	logger.Info().Str("Project Number", *projectNumber).Msg(indent)
	logger.Info().Str("Location", *location).Msg(indent)
	logger.Info().Str("Data Store", *dataStore).Msg(indent)
	logger.Info().Str("Branch", *branch).Msg(indent)
	logger.Info().Str("Serving Config", *servingConfig).Msg(indent)
	logger.Info().Str("Parameter Input File", *parameterInputFile).Msg(indent)
	logger.Info().Int("Number of Results", *numberResults).Msg(indent)
	logger.Info().Str("Filter String", *filterString).Msg(indent)
	logger.Info().Msg("Begin")

	// Setup the Recommendation Request
	var recommendation = NewRecommendation(*projectNumber, *location, *dataStore, *branch, *servingConfig, *numberResults, *filterString)

	//  Load the Parameter Input File
	logger.Info().Msg("Loading Parameter Input File")
	err := recommendation.LoadParameters(*parameterInputFile)
	if err != nil {
		logger.Error().Err(err).Msg("Load Parameter Input File Failed")
		os.Exit(1)
	}

	// Request to get a Recommendation from the Discovery Engine API
	err = recommendation.ExecuteRequests()
	if err != nil {
		logger.Error().Err(err).Msg("Recommendation Request Failed")
		os.Exit(1)
	}
	logger.Info().Msg("End")

}
