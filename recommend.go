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
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	discoveryengine "cloud.google.com/go/discoveryengine/apiv1beta"
	discoveryenginepb "cloud.google.com/go/discoveryengine/apiv1beta/discoveryenginepb"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/types/known/structpb"
)

type Recommendation struct {
	Project       string
	Location      string
	DataStore     string
	Branch        string
	ServingConfig string
	PageSize      int32
	Filter        string
	UserEvents    []interface{}
}

type Results struct {
	Id    string `json:"id,omitempty"`
	Title string `json:"title,omitempty"`
	Score string `json:"score,omitempty"`
}

//---------------------------------------------------------------------------------------

// Create a new Recommendation struct populated
func NewRecommendation(project string, location string, dataStore string, branch string, servingConfig string, numberResults int, filter string) Recommendation {
	return Recommendation{
		Project:       project,
		Location:      location,
		DataStore:     dataStore,
		Branch:        branch,
		ServingConfig: servingConfig,
		PageSize:      int32(numberResults),
		Filter:        filter,
	}
}

//---------------------------------------------------------------------------------------

// Execute the Requests to get a Recommendation from the Discovery Engine API
// for each of the UserEvent objects
func (recommendation *Recommendation) ExecuteRequests() error {

	// Establish a Discovery Engine API Recommendation Client
	logger.Info().Msg("Establishing a Discovery Engine Recommendation Client")
	ctx := context.Background()
	client, err := discoveryengine.NewRecommendationClient(ctx)
	if err != nil {
		return fmt.Errorf("Failed to establish a Discovery Engine Recommendation Client: %w", err)
	}
	defer client.Close()

	servingConfig := fmt.Sprintf("projects/%s/locations/%s/dataStores/%s/servingConfigs/%s",
		recommendation.Project, recommendation.Location, recommendation.DataStore, recommendation.ServingConfig)

	// Iterate through the User Events
	for i := 0; i < len(recommendation.UserEvents); i++ {

		// Encode the raw JSON User Event
		rawUserEvent, err := json.Marshal(recommendation.UserEvents[i])
		if err != nil {
			return fmt.Errorf("Encoding the User Event as JSON Failed: %w", err)
		}

		logger.Info().Int("Number", i+1).Msg("Initiating Recommendation Request")
		logger.Info().RawJSON("Parameters", rawUserEvent).Msg(indent)

		// Encode the Protobuf User Event using the `protojson` package
		pbUserEvent := new(discoveryenginepb.UserEvent)
		err = protojson.Unmarshal(rawUserEvent, pbUserEvent)
		if err != nil {
			return fmt.Errorf("Encoding the User Event as Protobuf Failed: %w", err)
		}

		// Populate Request Parameters
		request := discoveryenginepb.RecommendRequest{
			ServingConfig: servingConfig,
			UserEvent:     pbUserEvent,
			PageSize:      recommendation.PageSize,
			Filter:        recommendation.Filter,
			ValidateOnly:  false,
			Params:        make(map[string]*structpb.Value, 3),
		}
		request.Params["disableFallback"] = structpb.NewBoolValue(true)
		request.Params["strictFiltering"] = structpb.NewBoolValue(true)
		request.Params["returnDocument"] = structpb.NewBoolValue(true)
		request.Params["returnScore"] = structpb.NewBoolValue(true)

		// Raise the Recommendation Request
		response, err := client.Recommend(ctx, &request)
		if err != nil {
			return fmt.Errorf("Recommendation Request Failed: %w", err)
		}

		// Iterate through the results and get the Document Title
		for r := 0; r < len(response.Results); r++ {

			fields, err := recommendation.GetFields(response.Results[r].Document)
			if err != nil {
				return fmt.Errorf("Failed to get the Document Data Fields: %w", err)
			}

			results := Results{
				Id:    response.Results[r].Id,
				Title: fields["title"].GetStringValue(),
				Score: fmt.Sprintf("%.5f", response.Results[r].Metadata["score"].GetNumberValue()),
			}

			// Encode the Response Result before sending to the log
			rawResults, err := json.Marshal(&results)
			if err != nil {
				return fmt.Errorf("Encoding the Response Results as JSON Failed: %w", err)
			}
			logger.Info().RawJSON("Results", rawResults).Msg(indent)

		}

	}

	return nil
}

//---------------------------------------------------------------------------------------

// Walk the provided Parameter Input File GLOB and load all parameters
func (recommendation *Recommendation) LoadParameters(inputFile string) error {

	inputFile, _ = filepath.Abs(inputFile)
	buf, err := os.ReadFile(inputFile)
	if err != nil {
		return fmt.Errorf("Reading the Parameter Input File Failed: %w", err)
	}

	if err = json.Unmarshal(buf, &recommendation.UserEvents); err != nil {
		return fmt.Errorf("Parsing the Parameter Input File Failed: %w", err)
	}

	return nil
}

//---------------------------------------------------------------------------------------

// Execute a Requests to get a Document Title from the Discovery Engine API
func (recommendation *Recommendation) GetDocumentTitle(id string) (*string, error) {

	// Establish a Discovery Engine API Document Client
	ctx := context.Background()
	client, err := discoveryengine.NewDocumentClient(ctx)
	if err != nil {
		return nil, fmt.Errorf("Failed to establish a Discovery Engine Document Client: %w", err)
	}
	defer client.Close()

	// Populate Request Parameters
	request := discoveryenginepb.GetDocumentRequest{
		Name: fmt.Sprintf("projects/%s/locations/%s/dataStores/%s/branches/%s/documents/%s",
			recommendation.Project, recommendation.Location, recommendation.DataStore, recommendation.Branch, id),
	}

	// Raise the Get Document Request
	response, err := client.GetDocument(ctx, &request)
	if err != nil {
		return nil, fmt.Errorf("Document Request Failed: %w", err)
	}

	fields, err := recommendation.GetFields(response)
	if err != nil {
		return nil, fmt.Errorf("Failed to get the Document Data Fields: %w", err)
	}

	title := fields["title"].GetStringValue()
	return &title, nil
}

//---------------------------------------------------------------------------------------

// Execute a Requests to get a Document Title from the Discovery Engine API
func (recommendation *Recommendation) GetFields(document *discoveryenginepb.Document) (map[string]*structpb.Value, error) {

	// Is the Document Data held in JSON format
	if jd := document.GetJsonData(); jd != "" {
		var fields map[string]*structpb.Value

		if err := json.Unmarshal([]byte(jd), &fields); err != nil {
			return nil, fmt.Errorf("Failed to Unmarshal the Document JSON Data: %w", err)
		}
		return fields, nil
	}

	// Is the Document Data held in Struct format
	if sd := document.GetStructData(); sd != nil {
		return sd.Fields, nil
	}

	return nil, nil
}
