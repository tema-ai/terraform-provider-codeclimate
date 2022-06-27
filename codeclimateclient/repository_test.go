package codeclimateclient

import (
	"fmt"
	"net/http"
	"testing"
)

const (
	repositorySlug                   = "testorg/testarepo"
	expectedTestReporterID           = "0c89092bc2c088d667612ddd1a992ec62f643ded331f40783bcf6b847561234d"
	organizationID                   = "testorg"
	repositoryUrl                    = "https://github.com/testorg/testarepo"
	repositoryID                     = "696a76232df2736347000001"
	expectedHumanName                = "testarepo"
	expectedBranch                   = "master"
	expectedSelfLink                 = "https://codeclimate.com/repos/696a76232df2736347000001"
	expectedServicesLink             = "https://api.codeclimate.com/v1/repos/696a76232df2736347000001/services"
	expectedWebCoverageLink          = "https://codeclimate.com/repos/696a76232df2736347000001/coverage"
	expectedWeIssuesLink             = "https://codeclimate.com/repos/696a76232df2736347000001/issues"
	expectedMaintainabilityBadgeLink = "https://api.codeclimate.com/v1/badges/d6dd39dd84d8575d1ddd/maintainability"
	expectedTestCoverageBadge        = "https://api.codeclimate.com/v1/badges/d6dd39dd84d8575d1ddd/test_coverage"
)

func _testDataGetRepository(t *testing.T, repository *Repository, err error) {
	if err != nil {
		t.Fatal(err)
	}

	if repository.TestReporterId != expectedTestReporterID {
		t.Errorf("Expected test_reporter_id to be '%s', got: '%s'", expectedTestReporterID, repository.TestReporterId)
	}

	if repository.GithubSlug != repositorySlug {
		t.Errorf("Expected github slug to be '%s', got: '%s'", repositorySlug, repository.GithubSlug)
	}

	if repository.HumanName != expectedHumanName {
		t.Errorf("Expected human_name to be '%s', got: '%s'", expectedHumanName, repository.HumanName)
	}

	if repository.Branch != expectedBranch {
		t.Errorf("Expected branch to be '%s', got: '%s'", expectedBranch, repository.Branch)
	}

	if repository.LinkSelf != expectedSelfLink {
		t.Errorf("Expected self link to be '%s', got: '%s'", expectedSelfLink, repository.LinkSelf)
	}

	if repository.LinkServices != expectedServicesLink {
		t.Errorf("Expected servicees link to be '%s', got: '%s'", expectedServicesLink, repository.LinkServices)
	}

	if repository.LinkWebCoverage != expectedWebCoverageLink {
		t.Errorf("Expected web coverage link to be '%s', got: '%s'", expectedWebCoverageLink, repository.LinkWebCoverage)
	}

	if repository.LinkWebIssues != expectedWeIssuesLink {
		t.Errorf("Expected web issues link to be '%s', got: '%s'", expectedWeIssuesLink, repository.LinkWebIssues)
	}

	if repository.LinkMaintainabilityBadge != expectedMaintainabilityBadgeLink {
		t.Errorf("Expected maintainability badge link to be '%s', got: '%s'", expectedMaintainabilityBadgeLink, repository.LinkMaintainabilityBadge)
	}

	if repository.LinkTestCoverageBadge != expectedTestCoverageBadge {
		t.Errorf("Expected coverage badge link to be '%s', got: '%s'", expectedTestCoverageBadge, repository.LinkTestCoverageBadge)
	}
}

func TestClient_GetRepository(t *testing.T) {
	teardown := setup()
	defer teardown()

	mux.HandleFunc("/repos", func(w http.ResponseWriter, r *http.Request) {
		if r.FormValue("github_slug") != repositorySlug {
			t.Fatal(fmt.Errorf("received slug doesnt match `%s`", repositorySlug))
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, getFixture("repositories/multi_repository.json"))
	})

	repository, err := client.GetRepository(repositorySlug)
	_testDataGetRepository(t, repository, err)

}

func TestClient_GetRepositoryById(t *testing.T) {
	teardown := setup()
	defer teardown()

	mux.HandleFunc(fmt.Sprintf("/repos/%s", repositoryID), func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != fmt.Sprintf("/repos/%s", repositoryID) {
			t.Fatal(fmt.Errorf("received id doesnt match `%s`", repositoryID))
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, getFixture("repositories/repository.json"))
	})

	repository, err := client.GetRepositoryById(repositoryID)
	_testDataGetRepository(t, repository, err)

}

func TestClient_CreateOrganizationRepository(t *testing.T) {
	teardown := setup()
	defer teardown()

	mux.HandleFunc(fmt.Sprintf("/orgs/%s/repos", organizationID), func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, getFixture("repositories/create_repository_response.json"))
	})

	repository, err := client.CreateOrganizationRepository(organizationID, repositoryUrl)

	if err != nil {
		t.Fatal(err)
	}

	if repository.TestReporterId != expectedTestReporterID {
		t.Errorf("Expected test_reporter_id to be '%s', got: '%s'", expectedTestReporterID, repository.TestReporterId)
	}

	if repository.GithubSlug != repositorySlug {
		t.Errorf("Expected github slug to be '%s', got: '%s'", repositorySlug, repository.GithubSlug)
	}
}
func TestClient_DeleteOrganizationRepository(t *testing.T) {
	teardown := setup()
	defer teardown()

	mux.HandleFunc(fmt.Sprintf("repos/%s", repositoryID), func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNoContent)
	})

	err := client.DeleteOrganizationRepository(repositoryID)

	if err != nil {
		t.Fatal(err)
	}

}
