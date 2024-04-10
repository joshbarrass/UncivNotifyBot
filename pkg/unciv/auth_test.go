package unciv

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	RESPONSE_OLD_SERVER = "true"
	RESPONSE_NO_AUTH    = "{authVersion: 0}"
	RESPONSE_BASIC_AUTH = "{authVersion: 1}"
)

type GetAuthVersionTest struct {
	Response string
	Expected int
}

func TestGetAuthVersion(t *testing.T) {
	var responseMatrix map[string]GetAuthVersionTest = map[string]GetAuthVersionTest{
		"OldServer": {RESPONSE_OLD_SERVER, AUTH_VERSION_NO_AUTH},
		"NoAuth":    {RESPONSE_NO_AUTH, AUTH_VERSION_NO_AUTH},
		"BasicAuth": {RESPONSE_BASIC_AUTH, AUTH_VERSION_BASIC},
	}
	for name, test := range responseMatrix {
		t.Run(name, func(t *testing.T) {
			serveMux := http.NewServeMux()
			serveMux.HandleFunc(SERVER_LIVE_ROUTE, func(w http.ResponseWriter, _ *http.Request) {
				w.Write([]byte(test.Response))
			})
			testServer := httptest.NewServer(serveMux)
			testClient := testServer.Client()

			server, err := NewUncivServer(testServer.URL, testClient)
			assert.Nil(t, err)

			version, err := server.GetAuthVersion()
			assert.Nil(t, err)
			assert.Equal(t, test.Expected, version)
		})
	}
}
