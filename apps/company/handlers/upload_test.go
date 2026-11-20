package handlers

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestUploadSGBusinessRejectsInvalidJSON(t *testing.T) {
	for name, body := range map[string]string{
		"missing data":           `{}`,
		"null data":              `{"data":null}`,
		"duplicate data":         `{"data":{},"data":{}}`,
		"duplicate nested field": `{"data":{"name":"a","name":"b"}}`,
		"invalid UTF-8":          "{\"data\":\"\xff\"}",
		"trailing JSON":          `{"data":{}} {}`,
	} {
		t.Run(name, func(t *testing.T) {
			request := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(body))
			response := httptest.NewRecorder()
			// No repository is configured: invalid input must be rejected before persistence.
			registerDeps{}.UploadSGBusiness(response, request)
			if response.Code != http.StatusBadRequest {
				t.Fatalf("status = %d, want %d", response.Code, http.StatusBadRequest)
			}
		})
	}
}
