package response

import (
	"api/src/domain/model"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/go-cmp/cmp"
)

type testResponse struct {
	Message string `json:"message,omitempty"`
	ID      string `json:"id,omitempty"`
}

func TestOK(t *testing.T) {
	type args struct {
		body any
	}
	type expected struct {
		statusCode  int
		contentType string
		body        map[string]string
	}

	tests := []struct {
		testName string
		args     args
		expected expected
	}{
		{
			testName: "success",
			args: args{
				body: testResponse{Message: "success"},
			},
			expected: expected{
				statusCode:  http.StatusOK,
				contentType: "application/json",
				body:        map[string]string{"message": "success"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			w := httptest.NewRecorder()
			OK(w, tt.args.body)

			resp := w.Result()
			if diff := cmp.Diff(tt.expected.statusCode, resp.StatusCode); diff != "" {
				t.Errorf("status code mismatch (-want +got):\n%s", diff)
			}

			if diff := cmp.Diff(tt.expected.contentType, resp.Header.Get("Content-Type")); diff != "" {
				t.Errorf("Content-Type mismatch (-want +got):\n%s", diff)
			}

			var result map[string]string
			if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
				t.Fatalf("failed to decode response body: %v", err)
			}

			if diff := cmp.Diff(tt.expected.body, result); diff != "" {
				t.Errorf("body mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestCreated(t *testing.T) {
	type args struct {
		body any
	}
	type expected struct {
		statusCode  int
		contentType string
		body        map[string]string
	}

	tests := []struct {
		testName string
		args     args
		expected expected
	}{
		{
			testName: "created",
			args: args{
				body: testResponse{ID: "123"},
			},
			expected: expected{
				statusCode:  http.StatusCreated,
				contentType: "application/json",
				body:        map[string]string{"id": "123"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			w := httptest.NewRecorder()
			Created(w, tt.args.body)

			resp := w.Result()
			if diff := cmp.Diff(tt.expected.statusCode, resp.StatusCode); diff != "" {
				t.Errorf("status code mismatch (-want +got):\n%s", diff)
			}

			if diff := cmp.Diff(tt.expected.contentType, resp.Header.Get("Content-Type")); diff != "" {
				t.Errorf("Content-Type mismatch (-want +got):\n%s", diff)
			}

			var result map[string]string
			if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
				t.Fatalf("failed to decode response body: %v", err)
			}

			if diff := cmp.Diff(tt.expected.body, result); diff != "" {
				t.Errorf("body mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestHandleAppError(t *testing.T) {
	type args struct {
		err model.AppError
	}
	type expected struct {
		statusCode  int
		contentType string
		body        map[string]string
	}

	tests := []struct {
		testName string
		args     args
		expected expected
	}{
		{
			testName: "bad request",
			args: args{
				err: model.NewBadRequestError(nil, "TestDomain"),
			},
			expected: expected{
				statusCode:  http.StatusBadRequest,
				contentType: "application/json",
				body: map[string]string{
					"type":    model.BadRequestErrorName,
					"domain":  "TestDomain",
					"message": "BadRequestError [TestDomain]",
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			w := httptest.NewRecorder()
			HandleAppError(w, tt.args.err)

			resp := w.Result()
			if diff := cmp.Diff(tt.expected.statusCode, resp.StatusCode); diff != "" {
				t.Errorf("status code mismatch (-want +got):\n%s", diff)
			}

			if diff := cmp.Diff(tt.expected.contentType, resp.Header.Get("Content-Type")); diff != "" {
				t.Errorf("Content-Type mismatch (-want +got):\n%s", diff)
			}

			var result map[string]string
			if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
				t.Fatalf("failed to decode response body: %v", err)
			}

			if diff := cmp.Diff(tt.expected.body, result); diff != "" {
				t.Errorf("body mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
