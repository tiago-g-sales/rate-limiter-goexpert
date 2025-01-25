package pkg

import (

	"net/http"
	"net/http/httptest"
	"testing"
)


func TestRateLimiter(t *testing.T) {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	tests := []struct {
		name           string
		requests       int
		expectedStatus int
	}{
		{"Within limit", 5, http.StatusUnprocessableEntity},
		{"Exceed limit", 10, http.StatusUnprocessableEntity},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, err := http.NewRequest("GET", "/", nil)
			if err != nil {
				t.Fatal(err)
			}

			rr := httptest.NewRecorder()
			rlMiddleware := RateLimiter(handler)

			for i := 0; i < tt.requests; i++ {
				rlMiddleware.ServeHTTP(rr, req)
			}

			if status := rr.Code; status != tt.expectedStatus {
				t.Errorf("handler returned wrong status code: got %v want %v", status, tt.expectedStatus)
			}
		})
	}
}

