package unciv

import (
	"fmt"
	"net/http"

	yaml "gopkg.in/yaml.v3"
)

// Auth versions
const (
	AUTH_VERSION_NO_AUTH = 0
	AUTH_VERSION_BASIC   = 1
)

type ServerAuthJson struct {
	AuthVersion int `yaml:"authVersion"`
}

func (server *uncivServer) GetAuthVersion() (int, error) {
	u := server.URL.JoinPath(SERVER_LIVE_ROUTE)
	req, err := http.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		return 0, err
	}
	resp, err := server.HttpClient.Do(req)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("%w: status code %d", ErrBadStatusCode, resp.StatusCode)
	}

	var auth ServerAuthJson
	dec := yaml.NewDecoder(resp.Body)
	if err := dec.Decode(&auth); err != nil {
		// failed to decode JSON, so assume server does not support auth
		// see: https://github.com/yairm210/Unciv/pull/8716#issuecomment-1439683692
		/* >If the server returns either {authVersion: 0} or any unparsable
		   string (like the current true) the server is assumed to not
		   support authentication
		*/
		return 0, nil
	}
	return auth.AuthVersion, nil
}
