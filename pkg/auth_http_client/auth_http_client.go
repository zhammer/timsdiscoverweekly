package auth_http_client

import "net/http"

type transport struct {
	bearerToken string
}

func (t *transport) RoundTrip(r *http.Request) (*http.Response, error) {
	r.Header.Add("Authorization", "Bearer "+t.bearerToken)
	return http.DefaultTransport.RoundTrip(r)
}

func New(bearerToken string) *http.Client {
	return &http.Client{
		Transport: &transport{bearerToken},
	}
}
