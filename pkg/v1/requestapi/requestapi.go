package requestapi

import (
	"bytes"
	"crypto/tls"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/ajikamaludin/go-grpc_basic/pkg/v1/cert"
	"github.com/ajikamaludin/go-grpc_basic/pkg/v1/utils/constants"
)

// Info is the http req info
type ReqInfo struct {
	URL         string
	Method      string
	HeadersInfo map[string]interface{}
	Body        []byte
}

type ResInfo struct {
	StatusCode int
	Header     http.Header
	Body       []byte
}

func Invoke(reqinf *ReqInfo, timeout time.Duration, crt *cert.Cert) (*ResInfo, error) {
	var req *http.Request
	var err error

	switch reqinf.Method {
	case constants.MethodGET:
		req, err = http.NewRequest(constants.MethodGET, reqinf.URL, nil)
	case constants.MethodPOST:
		req, err = http.NewRequest(constants.MethodPOST, reqinf.URL, bytes.NewReader(reqinf.Body))
	}
	if err != nil {
		return nil, err
	}

	// set header
	for key, value := range reqinf.HeadersInfo {
		req.Header.Add(key, value.(string))
	}

	// execute
	cl := newHTTPCLientCrt(crt, timeout)

	res, err := cl.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if !(res.StatusCode == http.StatusOK || res.StatusCode == http.StatusCreated) {
		return nil, errors.New(fmt.Sprintf("%v for %v", res.StatusCode, reqinf.URL))
	}

	// read body
	buf, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return &ResInfo{
		StatusCode: res.StatusCode,
		Body:       buf,
	}, nil
}

func newHTTPCLientCrt(crt *cert.Cert, timeout time.Duration) *http.Client {
	// define the default http client
	defaultRoundTripper := http.DefaultTransport
	defaultTransportPtr, _ := defaultRoundTripper.(*http.Transport)
	tr := defaultTransportPtr.Clone()

	if crt == nil {
		return &http.Client{
			Timeout: timeout,
		}
	}

	// use cert
	tr.TLSClientConfig = &tls.Config{
		RootCAs: crt.Pool,
	}
	if crt.AllowSkip {
		tr.TLSClientConfig = &tls.Config{
			InsecureSkipVerify: true,
		}
	}

	// define the max connection idle
	tr.MaxIdleConns = 100
	tr.MaxIdleConnsPerHost = 20

	return &http.Client{
		Timeout:   timeout,
		Transport: tr,
	}

}
