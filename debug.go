package imuis

import (
	"fmt"
	"net/http"
	"net/http/httputil"
)

type DumpTransport struct {
	RoundTripper http.RoundTripper
}

func (d DumpTransport) RoundTrip(h *http.Request) (*http.Response, error) {
	dump, _ := httputil.DumpRequestOut(h, true)
	fmt.Println(string(dump))

	resp, err := d.RoundTripper.RoundTrip(h)
	dump, _ = httputil.DumpResponse(resp, true)
	fmt.Println(string(dump))

	return resp, err
}
