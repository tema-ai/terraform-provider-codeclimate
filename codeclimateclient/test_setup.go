package codeclimateclient

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
)

const (
	fixturesPath = "testdata/fixtures/"
)

var (
	mux    *http.ServeMux
	server *httptest.Server
	client *Client
)

func setup() func() {
	const mockAPIKey = "mockApiKey"

	mux = http.NewServeMux()
	server = httptest.NewServer(mux)
	client = &Client{ApiKey: mockAPIKey, BaseUrl: server.URL}

	return func() {
		server.Close()
	}
}

func getFixture(path string) string {
	b, err := ioutil.ReadFile(fixturesPath + path)
	if err != nil {
		panic(err)
	}
	return string(b)
}
