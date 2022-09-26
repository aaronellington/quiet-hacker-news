package forge

import (
	"net/http"
)

// MockableRoundTripper is foobar
type MockableRoundTripper struct {
	RoundTripFunc func(*http.Request) (*http.Response, error)
}

// RoundTrip is foobar
func (mockable MockableRoundTripper) RoundTrip(request *http.Request) (*http.Response, error) {
	return mockable.RoundTripFunc(request)
}
