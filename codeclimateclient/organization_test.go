package codeclimateclient

import (
	"fmt"
	"net/http"
	"testing"
)

const (
	organizationName           = "testorg2"
	expectedTestOrganizationId = "5f2176144f78f5011f002a89"
)

func TestOrganizationGetId(t *testing.T) {
	teardown := setup()
	defer teardown()

	mux.HandleFunc("/orgs", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, getFixture("organizations/organizations.json"))
	})

	organization, err := client.GetOrganization(organizationName)

	if err != nil {
		t.Fatal(err)
	}

	if organization.Id != expectedTestOrganizationId {
		t.Errorf("Expected organization_id to be '%s', got: '%s'", expectedTestOrganizationId, organization.Id)
	}
}
