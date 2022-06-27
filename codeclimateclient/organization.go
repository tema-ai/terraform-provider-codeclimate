package codeclimateclient

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Organization struct {
	Id   string
	Name string
}

type readOrganizationsResponse struct {
	Data []struct {
		ID         string `json:"id"`
		Attributes struct {
			Name string `json:"name"`
		} `json:"attributes"`
	} `json:"data"`
}

func (client *Client) GetOrganization(organizationName string) (*Organization, error) {
	var organizationsData readOrganizationsResponse

	data, err := client.makeRequest(http.MethodGet, "orgs", nil)

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data, &organizationsData)
	if err != nil {
		return nil, err
	}

	for _, org := range organizationsData.Data {
		if org.Attributes.Name == organizationName {
			organization := &Organization{
				Id:   org.ID,
				Name: org.Attributes.Name,
			}
			return organization, nil
		}
	}
	return nil, fmt.Errorf("The Organization could not be found")
}
