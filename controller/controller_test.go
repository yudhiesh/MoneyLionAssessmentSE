package controller

import (
	"bytes"
	"net/http"
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
		{"Missing Entire URL Parameter", "/feature", http.StatusUnprocessableEntity, []byte(`{"status":422,"message":"Missing URL query parameters email/featureName"}`)},
		{"Missing Email URL Parameter", "/feature?email=test@gmail.com", http.StatusUnprocessableEntity, []byte(`{"status":422,"message":"Missing URL query parameters email/featureName"}`)},
		{"Missing featureName URL Parameter", "/feature?featureName=financial-tracking", http.StatusUnprocessableEntity, []byte(`{"status":422,"message":"Missing URL query parameters email/featureName"}`)},
		{"Missing featureName URL Parameter", "/feature?featureName=financial-tracking", http.StatusUnprocessableEntity, []byte(`{"status":422,"message":"Missing URL query parameters email/featureName"}`)},
		{"Email does not exist", "/feature?email=test@gmail.com&featureName=financial-tracking", http.StatusNotFound, []byte(`{"status":404,"message":"No matching record found"}`)},
		{"Email does not exist", "/feature?email=test@gmail.com&featureName=financial-tracking", http.StatusNotFound, []byte(`{"status":404,"message":"No matching record found"}`)},
		{"Email and FeatureName exist", "/feature?email=test1@gmail.com&featureName=financial-tracking", http.StatusOK, []byte(`{"response":{"status":200,"message":"Success"},"data":{"can_access":true}}`)},
		{"Email and FeatureName exist", "/feature?email=test1@gmail.com&featureName=crypto", http.StatusOK, []byte(`{"response":{"status":200,"message":"Success"},"data":{"can_access":true}}`)},
		{"FeatureName does not exist", "/feature?email=test1@gmail.com&featureName=premium", http.StatusNotFound, []byte(`{"status":404,"message":"No matching record found"}`)},
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

// `
// ERROR	2021/06/04 22:07:15 controller.go:50: No matching record found
// INFO	2021/06/04 22:08:23 HTTP/1.1 GET /feature
// ERROR	2021/06/04 22:08:23 controller.go:42: Missing URL query parameters email/featureName
// INFO	2021/06/04 22:08:23 HTTP/1.1 GET /feature?email=test@gmail.com
// ERROR	2021/06/04 22:08:23 controller.go:42: Missing URL query parameters email/featureName
// INFO	2021/06/04 22:08:23 HTTP/1.1 GET /feature?featureName=financial-tracking
// ERROR	2021/06/04 22:08:23 controller.go:42: Missing URL query parameters email/featureName
// INFO	2021/06/04 22:08:23 HTTP/1.1 GET /feature?featureName=financial-tracking
// ERROR	2021/06/04 22:08:23 controller.go:42: Missing URL query parameters email/featureName
// INFO	2021/06/04 22:08:23 HTTP/1.1 GET /feature?email=test@gmail.com&featureName=financial-tracking
// ERROR	2021/06/04 22:08:23 controller.go:50: No matching record found
// INFO	2021/06/04 22:08:23 HTP/1.1 GET /feature?email=test@gmail.com&featureName=financial-tracking
// ERROR	2021/06/04 22:08:23 controller.go:50: No matching record found
// INFO	2021/06/04 22:08:23 HTTP/1.1 GET /feature?email=test1@gmail.com&featureName=financial-tracking
// INFO	2021/06/04 22:08:23 Success
// INFO	2021/06/04 22:08:23 HTTP/1.1 GET /feature?email=test1@gmail.com&featureName=crypto
// INFO	2021/06/04 22:08:23 Success
// INFO	2021/06/04 22:08:23 HTTP/1.1 GET /feature?email=test1@gmail.com&featureName=premium
// ERROR	2021/06/04 22:08:23 controller.go:50: No matching record found

// `
