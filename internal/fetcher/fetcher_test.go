package fetcher_test

import (
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/iamolegga/md5http/internal/fetcher"
	"github.com/iamolegga/md5http/pkg/randbuf"
)

func TestFetcher_Fetch(t *testing.T) {
	payload, _ := randbuf.New(16)
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write(payload)
	})
	mux.HandleFunc("/empty", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})
	srv := httptest.NewServer(mux)
	defer srv.Close()

	tests := []struct {
		name     string
		url      string
		wantBody []byte
		wantErr  bool
	}{
		{"positive", srv.URL, payload, false},
		{"negative", "http://127.0.0.1:12345", nil, true},
		{"empty", srv.URL + "/empty", []byte{}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := fetcher.New()
			gotBody, err := f.Fetch(tt.url)
			if (err != nil) != tt.wantErr {
				t.Errorf("Fetch() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotBody, tt.wantBody) {
				t.Errorf("Fetch() gotBody = %v, want %v", gotBody, tt.wantBody)
			}
		})
	}
}
