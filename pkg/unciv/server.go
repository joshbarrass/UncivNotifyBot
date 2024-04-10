package unciv

import (
	"net/http"
	"net/url"
)

const DEFAULT_UNCIV_SERVER_URL = "https://uncivserver.xyz"

// routes
const (
	SERVER_LIVE_ROUTE  = "/islive"
	SERVER_FILES_ROUTE = "/files"
)

type UncivServer interface {
	GetAuthVersion() (int, error)
}

type uncivServer struct {
	URL        *url.URL
	HttpClient *http.Client
}

func NewUncivServer(u string, client *http.Client) (UncivServer, error) {
	if client == nil {
		client = http.DefaultClient
	}
	parsedURL, err := url.Parse(u)
	if err != nil {
		return nil, err
	}
	return &uncivServer{
		URL:        parsedURL,
		HttpClient: client,
	}, nil
}

func NewDefaultUncivServer() UncivServer {
	server, _ := NewUncivServer(DEFAULT_UNCIV_SERVER_URL, nil)
	return server
}
