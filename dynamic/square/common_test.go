package square

import "net/http"

type testRoundTripper struct {
	roundTripFunc func(r *http.Request) (*http.Response, error)
}

func (t *testRoundTripper) RoundTrip(r *http.Request) (*http.Response, error) {
	return t.roundTripFunc(r)
}
