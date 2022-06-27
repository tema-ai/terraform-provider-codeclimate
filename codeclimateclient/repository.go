package codeclimateclient

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type Repository struct {
	Id                       string
	TestReporterId           string
	GithubSlug               string
	Organization             string
	RepositoryURL            string
	HumanName                string
	Branch                   string
	Attributes               string
	LinkServices             string
	LinkSelf                 string
	LinkWebCoverage          string
	LinkWebIssues            string
	LinkMaintainabilityBadge string
	LinkTestCoverageBadge    string
}

type BasicResponse struct {
	ID         string `json:"id"`
	Attributes struct {
		TestReporterID string `json:"test_reporter_id"`
		GithubSlug     string `json:"github_slug"`
		VCSHost        string `json:"vcs_host"`
		HumanName      string `json:"human_name"`
		Branch         string `json:"branch"`
	} `json:"attributes"`
	Links struct {
		Self                 string `json:"self"`
		Services             string `json:"services"`
		WebCoverage          string `json:"web_coverage"`
		WebIssues            string `json:"web_issues"`
		MaintainabilityBadge string `json:"maintainability_badge"`
		TestCoverageBadge    string `json:"test_coverage_badge"`
	} `json:"links"`
	Relationships struct {
		Account struct {
			Data struct {
				ID string `json:"id"`
			} `json:"data"`
		} `json:"account"`
	} `json:"relationships"`
}

// The structure describes just what we need from the response.
//  For the full description look at: https://developer.codeclimate.com/?shell#get-repository
type readMultiRepositoriesResponse struct {
	Data []struct {
		BasicResponse
	} `json:"data"`
}

type readRepositoriesResponse struct {
	Data struct {
		BasicResponse
	} `json:"data"`
}

type createRepositoryResponse struct {
	Data struct {
		ID         string `json:"id"`
		Attributes struct {
			TestReporterID string `json:"test_reporter_id"`
			GithubSlug     string `json:"github_slug"`
		} `json:"attributes"`
	} `json:"data"`
}

type errorResponse struct {
	Errors []struct {
		Detail string `json:"detail"`
		Title  string `json:"title"`
	} `json:"errors"`
}

func _UnpackRepository(data struct{ BasicResponse }) *Repository {
	repository := &Repository{
		Id:                       data.ID,
		TestReporterId:           data.Attributes.TestReporterID,
		GithubSlug:               data.Attributes.GithubSlug,
		HumanName:                data.Attributes.HumanName,
		Branch:                   data.Attributes.Branch,
		Organization:             data.Relationships.Account.Data.ID,
		RepositoryURL:            data.Attributes.VCSHost + "/" + data.Attributes.GithubSlug,
		LinkServices:             data.Links.Services,
		LinkSelf:                 data.Links.Self,
		LinkWebCoverage:          data.Links.WebCoverage,
		LinkWebIssues:            data.Links.WebIssues,
		LinkMaintainabilityBadge: data.Links.MaintainabilityBadge,
		LinkTestCoverageBadge:    data.Links.TestCoverageBadge,
	}
	return repository
}

func _ProcessMultiRepository(data []byte, err error) (*Repository, error) {
	var repositoryData readMultiRepositoriesResponse
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, &repositoryData)
	if err != nil {
		return nil, err
	}
	numberOfReposFound := len(repositoryData.Data)
	if numberOfReposFound != 1 {
		return nil, fmt.Errorf(
			"The response returned %v repositories (should have been 1)",
			numberOfReposFound,
		)
	}
	var BasicRepositoryData struct{ BasicResponse } = repositoryData.Data[0]
	repository := _UnpackRepository(BasicRepositoryData)
	return repository, nil
}

func _ProcessRepository(data []byte, err error) (*Repository, error) {
	var repositoryData readRepositoriesResponse
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, &repositoryData)
	if err != nil {
		return nil, err
	}
	var BasicRepositoryData struct{ BasicResponse } = repositoryData.Data
	repository := _UnpackRepository(BasicRepositoryData)
	return repository, nil
}

func (client *Client) GetRepository(repositorySlug string) (*Repository, error) {
	data, err := client.makeRequest(http.MethodGet, fmt.Sprintf("repos?github_slug=%s", repositorySlug), nil)
	return _ProcessMultiRepository(data, err)
}

func (client *Client) GetRepositoryById(repositoryId string) (*Repository, error) {
	data, err := client.makeRequest(http.MethodGet, fmt.Sprintf("repos/%s", repositoryId), nil)
	return _ProcessRepository(data, err)
}

func (client *Client) CreateOrganizationRepository(organizationID string, url string) (*Repository, error) {
	var repositoryData createRepositoryResponse

	payload := fmt.Sprintf(`
		{
			"data": {
				"type": "repos",
				"attributes": {
					"url": "%s"
				}
			}
		}`, url)

	data, err := client.makeRequest(http.MethodPost, fmt.Sprintf("orgs/%s/repos", organizationID), strings.NewReader(payload))
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data, &repositoryData)

	if err != nil {
		return nil, err
	}

	if repositoryData.Data.ID == "" {
		return nil, handleErrorResponse(data)
	}

	repository := &Repository{
		Id:             repositoryData.Data.ID,
		TestReporterId: repositoryData.Data.Attributes.TestReporterID,
		GithubSlug:     repositoryData.Data.Attributes.GithubSlug,
	}

	return repository, nil

}

func (client *Client) DeleteOrganizationRepository(repositoryID string) error {
	_, err := client.makeRequest(http.MethodDelete, fmt.Sprintf("repos/%s", repositoryID), nil)

	if err != nil {
		return fmt.Errorf("The repository couldn't be deleted: %s", err)
	}

	return nil
}

func handleErrorResponse(data []byte) error {
	var errorResponse errorResponse
	err := json.Unmarshal(data, &errorResponse)
	if err != nil {
		return err
	}
	return fmt.Errorf("Invalid json response %s", errorResponse.Errors)
}
