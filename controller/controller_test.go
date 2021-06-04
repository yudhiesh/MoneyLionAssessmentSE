package controller

import (
	"bytes"
	"net/http"
	"net/url"
	"strconv"
	"testing"
)

func TestGetCanAcess(t *testing.T) {

	app := newTestApplication(t)
	ts := newTestServer(t, app.routes())
	defer ts.Close()
	tests := []struct {
		name     string
		urlPath  string
		wantCode int
		wantBody []byte
	}{
		{"Missing Entire URL Parameter", "/feature", http.StatusNotFound, []byte(`{"error":"Missing URL query parameters email/featureName"}`)},
		{"Missing Email URL Parameter", "/feature?email=test@gmail.com", http.StatusNotFound, []byte(`{"error":"Missing URL query parameters email/featureName"}`)},
		{"Missing featureName URL Parameter", "/feature?featureName=financial-tracking", http.StatusNotFound, []byte(`{"error":"Missing URL query parameters email/featureName"}`)},
		{"Email does not exist", "/feature?email=test@gmail.com&featureName=financial-tracking", http.StatusNotFound, []byte(`{"error":"No matching record found"}`)},
		{"Email and FeatureName exist", "/feature?email=test1@gmail.com&featureName=financial-tracking", http.StatusOK, []byte(`{"can_access":true}`)},
		{"Email and FeatureName exist", "/feature?email=test1@gmail.com&featureName=crypto", http.StatusOK, []byte(`{"can_access":true}`)},
		{"FeatureName does not exist", "/feature?email=test1@gmail.com&featureName=premium", http.StatusNotFound, []byte(`{"error":"No matching record found"}`)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			statusCode, _, body := ts.get(t, tt.urlPath)
			if statusCode != tt.wantCode {
				t.Errorf("want %d; got %d", tt.wantCode, statusCode)
			}
			if !bytes.Contains(body, tt.wantBody) {
				t.Errorf("want body to contain %q", tt.wantBody)
			}
		})
	}
}

func TestInsetFeature(t *testing.T) {

	app := newTestApplication(t)
	ts := newTestServer(t, app.routes())
	defer ts.Close()
	tests := []struct {
		name        string
		urlPath     string
		featureName string
		email       string
		enable      bool
		wantCode    int
	}{
		{"Missing body 1", "/feature", "", "", false, http.StatusNotModified},
		{"Missing body 2", "/feature", "", "", true, http.StatusNotModified},
		{"Incorrect value for parameters 1", "/feature", "premium", "test1@gmail.com", false, http.StatusNotModified},
		{"Correct case", "/feature", "automated-investing", "test3@gmail.com", true, http.StatusOK},
		{"Incorrect user email", "/feature", "premium", "test10@gmail.com", true, http.StatusNotModified},
		{"Correct case", "/feature", "financial-tracking", "test4@gmail.com", true, http.StatusOK},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			form := url.Values{}
			form.Add("featureName", tt.featureName)
			form.Add("email", tt.email)
			enable := strconv.FormatBool(tt.enable)
			form.Add("enable", enable)
			statusCode, _, _ := ts.postForm(t, tt.urlPath, form)
			if statusCode != tt.wantCode {
				t.Errorf("want %d; got %d", tt.wantCode, statusCode)
			}
		})
	}
}
