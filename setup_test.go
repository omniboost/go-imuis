package imuis_test

import (
	"log"
	"net/http"
	"os"
	"testing"

	"github.com/fiorix/wsdl2go/soap"
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

	h := http.DefaultClient
	h.Transport = imuis.DumpTransport{RoundTripper: http.DefaultTransport}
	cli := &soap.Client{
		URL:         "https://cloudswitch.imuisonline.com/ws1_xml.asmx",
		Namespace:   imuis.Namespace,
		Config:      h,
		ContentType: "text/xml; charset=utf-8; action=\"%s\"",
		Envelope:    "http://www.w3.org/2003/05/soap-envelope",
	}

	client = imuis.NewWs1_xmlSoap(cli)

	resp, err := client.Login(&imuis.Login{PartnerKey: &partnerKey, Omgevingscode: &omgevingsCode})
	if err != nil {
		log.Fatal(err)
	}
	sessionID = *resp.SessionId

	m.Run()
}
