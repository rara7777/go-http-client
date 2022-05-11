package examples

import (
	"fmt"
)

type Endpoints struct {
	CurrentUserUrl    string `json:"current_user_url"`
	AuthorizationsUrl string `json:"authorizations_url"`
	RepositoryUrl     string `json:"repository_url"`
}

type EmptyRequest struct{}

func GetEndpoints(request EmptyRequest) (*Endpoints, error) {
	response, err := httpClient.Get("https://api.github.com", request)
	if err != nil {
		// Deal with the error as you need:
		return nil, err
	}

	fmt.Printf("Status Code: %d", response.StatusCode)
	fmt.Printf("Status: %s", response.Status)
	fmt.Printf("Body: %s\n", response.String())

	var endpoints Endpoints
	if err := response.UnmarshalJson(&endpoints); err != nil {
		// Deal with the unmarshal error as you need:
		return nil, err
	}

	fmt.Printf("Repository URL: %s", endpoints.RepositoryUrl)
	return &endpoints, nil
}
