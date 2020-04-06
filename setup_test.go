package imuis_test

import (
	"log"
	"net/http"
	"os"
	"testing"

	"github.com/omniboost/go-imuis"
)

var (
	client        imuis.Ws1_xmlSoap
	partnerKey    string
	omgevingsCode string
	debug         bool
	sessionID     string
)

func TestMain(m *testing.M) {
	partnerKey = os.Getenv("IMUIS_PARTNERKEY")
	omgevingsCode = os.Getenv("IMUIS_OMGEVINGSCODE")
	if os.Getenv("IMUIS_DEBUG") != "" {
		debug = true
	}

	// h.Transport = imuis.DumpTransport{RoundTripper: http.DefaultTransport}

	h := http.DefaultClient
	h.Transport = imuis.DumpTransport{RoundTripper: http.DefaultTransport}
	client = imuis.NewClient(h)
	resp, err := client.Login(&imuis.Login{PartnerKey: &partnerKey, Omgevingscode: &omgevingsCode})
	if err != nil {
		log.Fatal(err)
	}
	sessionID = *resp.SessionId

	m.Run()
}
