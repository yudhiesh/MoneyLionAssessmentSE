package controller

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/http/httptest"
	"testing"
)

// Custom testServer which anonymously embeds a httptest.Server instance
type testServer struct {
	*httptest.Server
}

// Returns a new Application struct with infoLogs and ErrorLogs filled
func newTestApplication(t *testing.T) *Application {
	return &Application{
		ErrorLog: log.New(ioutil.Discard, "", 0),
		InfoLog:  log.New(ioutil.Discard, "", 0),
	}
}
func newTestServer(t *testing.T, h http.Handler) *testServer {
	// Creates a new test server passing in the value returned by our
	// app.routes() method as the handler for the server.
	// use httptest.NewServer() if testing http request
	ts := httptest.NewTLSServer(h)
	// Store any cookies sent in a HTTPS response, so that we can include them
	// in any subsequent requests back to the test server.
	jar, err := cookiejar.New(nil)
	if err != nil {
		t.Fatal(err)
	}
	ts.URL = "http://localhost:3000"
	ts.Client().Jar = jar
	// Don't automatically follow redirects, instead return the first HTTPS
	// response sent by our server so that we can test the reponse for that
	// specific request.
	ts.Client().CheckRedirect = func(req *http.Request, via []*http.Request) error {
		return http.ErrUseLastResponse
	}
	return &testServer{ts}
}

// Get method to the custom testServer type. This makes a GET request to the
// given urlPath and returns the StatusCode, Header, and the body
func (ts *testServer) get(t *testing.T, urlPath string) (int, http.Header, []byte) {
	rs, err := ts.Client().Get(ts.URL + urlPath)
	if err != nil {
		t.Fatal(err)
	}
	defer rs.Body.Close()
	body, err := ioutil.ReadAll(rs.Body)
	if err != nil {
		t.Fatal(err)
	}
	return rs.StatusCode, rs.Header, body
}

// Create a post method for sending a POST request to the test server
// form is a url.Values object which can contain any data that you want to send
// in the request body
func (ts *testServer) post(t *testing.T, urlPath string, body []byte) (int, http.Header, []byte) {
	rs, err := ts.Client().Post(ts.URL+urlPath, "application/json", bytes.NewReader(body))
	if err != nil {
		t.Fatal(err)
	}
	// Read the response body.
	defer rs.Body.Close()
	body, err = ioutil.ReadAll(rs.Body)
	if err != nil {
		t.Fatal(err)
	}

	// Return the response status, headers and body.
	return rs.StatusCode, rs.Header, body
}
