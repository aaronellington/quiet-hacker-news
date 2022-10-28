package forge

import (
	"net/http"
)

type MockableRoundTripper struct {
	RoundTripFunc func(*http.Request) (*http.Response, error)
}

func (mockable MockableRoundTripper) RoundTrip(request *http.Request) (*http.Response, error) {
	return mockable.RoundTripFunc(request)
}
