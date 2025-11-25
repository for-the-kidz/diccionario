package diccionario

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/for-the-kidz/diccionario/wordlist"
	"github.com/gin-gonic/gin"
)

type fakeWordList struct {
	words []string
	err   error
}

func (f *fakeWordList) AddWord(word string) error {
	return nil
}

func (f *fakeWordList) GetWords() ([]string, error) {
	if f.err != nil {
		return nil, f.err
	}
	return f.words, nil
}

func newTestServer(w wordlist.WordList) *Server {
	gin.SetMode(gin.TestMode)
	r := gin.New()
	s := &Server{
		r: r,
		w: w,
	}

	r.GET("/exists/:word", s.WordExists)

	return s
}

func TestWordExists(t *testing.T) {
	tests := []struct {
		name       string
		wordParam  string
		words      []string
		getErr     error
		wantStatus int
		wantExists bool
		wantBody   string
	}{
		{
			name:       "word exists with exact match",
			wordParam:  "hola",
			words:      []string{"hola", "adios"},
			wantStatus: http.StatusOK,
			wantExists: true,
		},
		{
			name:       "word exists as prefix",
			wordParam:  "ad",
			words:      []string{"hola", "adios"},
			wantStatus: http.StatusOK,
			wantExists: true,
		},
		{
			name:       "word does not exist",
			wordParam:  "bonjour",
			words:      []string{"hola", "adios"},
			wantStatus: http.StatusOK,
			wantExists: false,
		},
		{
			name:       "empty word list",
			wordParam:  "hola",
			words:      []string{},
			wantStatus: http.StatusOK,
			wantExists: false,
		},
		{
			name:       "GetWords returns error",
			wordParam:  "hola",
			getErr:     errors.New("boom"),
			wantStatus: http.StatusBadRequest,
			wantBody:   "boom",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fw := &fakeWordList{
				words: tt.words,
				err:   tt.getErr,
			}

			s := newTestServer(fw)

			rec := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodGet, "/exists/"+tt.wordParam, nil)

			s.r.ServeHTTP(rec, req)

			if rec.Code != tt.wantStatus {
				t.Fatalf("status code = %d, want %d", rec.Code, tt.wantStatus)
			}

			if tt.wantStatus != http.StatusOK {
				if tt.wantBody != "" && rec.Body.String() != tt.wantBody {
					t.Fatalf("body = %q, want %q", rec.Body.String(), tt.wantBody)
				}
				return
			}

			var resp ExistsResponse
			if err := json.Unmarshal(rec.Body.Bytes(), &resp); err != nil {
				t.Fatalf("failed to unmarshal response: %v", err)
			}

			if resp.Exists != tt.wantExists {
				t.Fatalf("Exists = %v, want %v", resp.Exists, tt.wantExists)
			}
		})
	}
}
