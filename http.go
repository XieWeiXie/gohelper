package gohelper

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

/*

 */
type HTTPAction struct {
	Url      string
	query    string
	val      url.Values
	format   string
	header   *http.Header
	request  *http.Request
	response *http.Response
}

func (H *HTTPAction) newHttpRequest(method string, body io.Reader) {
	req, _ := http.NewRequest(method, H.Url, body)
	H.request = req
}

func (H *HTTPAction) Add(key string, value string) {
	if H.val == nil {
		H.val = make(url.Values)
	}
	H.val, _ = url.ParseQuery(H.Url)
	H.val.Add(key, value)
}

func (H *HTTPAction) Format() {
	if H.val == nil {
		H.val = make(url.Values)
	}
	q := H.val.Encode()
	H.request.URL.RawQuery = q
}

func (H *HTTPAction) HeaderAdd(key string, value string) {
	if H.header == nil {
		H.header = &H.request.Header
	}
	H.header.Add(key, value)
}

func (H *HTTPAction) Do() {
	H.Run()
}
func (H *HTTPAction) String() string {
	var (
		out strings.Builder
	)
	out.WriteString(fmt.Sprintf("[%d %s]", H.response.StatusCode, http.StatusText(H.response.StatusCode)))
	return out.String()

}
func (H *HTTPAction) Run() {
	client := http.DefaultClient
	H.response, _ = client.Do(H.request)
}
func (H *HTTPAction) RoundTripper() {
	transport := http.DefaultTransport
	H.response, _ = transport.RoundTrip(H.request)
}
