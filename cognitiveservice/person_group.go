package cognitiveservice

import (
	"bytes"
	"encoding/json"
	"net/http"
)

// PersonGroupAPI is responsible for person group creation, query and delete
type PersonGroupAPI interface {
	Create(personGroupObject PersonGroupObject) error
	Query(personGroupID string) (PersonGroupObject, error)
	Delete(personGroupID string) error
}

// MicrosoftPersonGroupAPI is responsible for person group creation and lookup
type MicrosoftPersonGroupAPI struct {
	BaseURL string
	APIKey  string
}

type personGroupRequestBody struct {
	Name     string `json:"name"`
	UserData string `json:"userData"`
}

// PersonGroupObject represents a unit of information about a person group
type PersonGroupObject struct {
	PersonGroupID string `json:"personGroupId"`
	Name          string `json:"name"`
	UserData      string `json:"userData"`
}

// Delete deletes a person group
func (api MicrosoftPersonGroupAPI) Delete(personGroupID string) error {
	requestURL := api.BaseURL + personGroupID

	err := commonHTTPRequest(http.MethodDelete, requestURL, api.APIKey, nil, nil)

	return err
}

// Create creates a person group by calling Microsoft cognitive service API
func (api MicrosoftPersonGroupAPI) Create(personGroupObject PersonGroupObject) error {
	requestURL := api.BaseURL + "/" + personGroupObject.PersonGroupID

	personGroupRequestBody := personGroupRequestBody{personGroupObject.Name, personGroupObject.UserData}

	marsheledRequestBody, _ := json.Marshal(personGroupRequestBody)

	err := commonHTTPRequest(http.MethodPut, requestURL, api.APIKey, bytes.NewReader(marsheledRequestBody), nil)

	return err
}

// Query implements PersonGroupAPI.Query
func (api MicrosoftPersonGroupAPI) Query(personGroupID string) (PersonGroupObject, error) {
	requestURL := api.BaseURL + personGroupID

	personGroupObject := PersonGroupObject{}

	err := commonHTTPRequest(http.MethodGet, requestURL, api.APIKey, nil, &personGroupObject)

	return personGroupObject, err
}

// NewMicrosoftPersonGroupAPI supports person group creation and lookup with Microsoft cognitive API
// NewMicrosoftPersonGroupAPI returns MicrosoftPersonGroupAPI that implements PersonGroupApi
func NewMicrosoftPersonGroupAPI(baseURL string, apiKey string) MicrosoftPersonGroupAPI {
	return MicrosoftPersonGroupAPI{baseURL, apiKey}
}
