package controller

import (
	"bytes"
	"net/http"
	"testing"
)

func TestGetCanAcess(t *testing.T) {

	app := newTestApplication(t)
	ts := newTestServer(t, app.Routes())
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

func TestInsertFeature(t *testing.T) {

	app := newTestApplication(t)
	ts := newTestServer(t, app.Routes())
	defer ts.Close()
	tests := []struct {
		name     string
		urlPath  string
		body     []byte
		wantCode int
	}{
		{"Missing entire body", "/feature", []byte(``), http.StatusNotModified},
		{"Missing body 1", "/feature", []byte(`{"featureName": "", "email": "", "can_access": ""}`), http.StatusNotModified},
		{"Missing body 2", "/feature", []byte(`{"featureName": "", "can_access": ""}`), http.StatusNotModified},
		{"Missing body 3", "/feature", []byte(`{"email": "", "can_access": ""}`), http.StatusNotModified},
		{"Incorrect value for parameters 1", "/feature", []byte(`{"featureName": "premium", "email": "test3@gmail.com", "can_access": true}`), http.StatusNotModified},
		{"Incorrect user email", "/feature", []byte(`{"featureName": "premium", "email": "test10@gmail.com", "can_access": true}`), http.StatusNotModified},
		{"Incorrect can_access", "/feature", []byte(`{"featureName": "financial-tracking", "email": "test2@gmail.com"}, "can_access": false`), http.StatusNotModified},
		{"Correct case 1", "/feature", []byte(`{"featureName": "automated-investing", "email": "test4@gmail.com", "can_access": true}`), http.StatusOK},
		{"Correct case 2", "/feature", []byte(`{"featureName": "crypto", "email": "test1@gmail.com", "can_access": false}`), http.StatusOK},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			statusCode, _, _ := ts.postForm(t, tt.urlPath, tt.body)
			if statusCode != tt.wantCode {
				t.Errorf("want %d; got %d", tt.wantCode, statusCode)
			}
		})
	}
}
