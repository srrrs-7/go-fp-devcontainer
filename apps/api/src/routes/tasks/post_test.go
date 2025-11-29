package tasks

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

func TestPostHandler(t *testing.T) {
	type args struct {
		formData map[string]string
	}
	type expected struct {
		statusCode int
		hasError   bool
	}

	tests := []struct {
		testName string
		args     args
		expected expected
	}{
		{
			testName: "valid request",
			args: args{
				formData: map[string]string{
					"title":       "New Task",
					"description": "Task Description",
				},
			},
			expected: expected{
				statusCode: http.StatusCreated,
				hasError:   false,
			},
		},
		{
			testName: "title too short",
			args: args{
				formData: map[string]string{
					"title": "ab",
				},
			},
			expected: expected{
				statusCode: http.StatusBadRequest,
				hasError:   true,
			},
		},
		{
			testName: "missing title",
			args: args{
				formData: map[string]string{
					"description": "Only description",
				},
			},
			expected: expected{
				statusCode: http.StatusBadRequest,
				hasError:   true,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			formData := url.Values{}
			for k, v := range tt.args.formData {
				formData.Set(k, v)
			}
			req := httptest.NewRequest(http.MethodPost, "/tasks", strings.NewReader(formData.Encode()))
			req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

			w := httptest.NewRecorder()
			PostHandler(w, req)

			resp := w.Result()
			if resp.StatusCode != tt.expected.statusCode {
				t.Errorf("expected status %v, got %v", tt.expected.statusCode, resp.StatusCode)
			}

			if resp.Header.Get("Content-Type") != "application/json" {
				t.Errorf("expected Content-Type application/json, got %v", resp.Header.Get("Content-Type"))
			}

			var result map[string]interface{}
			if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
				t.Fatalf("failed to decode response body: %v", err)
			}

			if tt.expected.hasError {
				if _, ok := result["type"]; !ok {
					t.Errorf("expected error response to have 'type' field")
				}
			} else {
				if _, ok := result["id"]; !ok {
					t.Errorf("expected success response to have 'id' field")
				}
			}
		})
	}
}
