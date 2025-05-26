package DarkThroneApi

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestIsZeroValue(t *testing.T) {
	t.Run("zero int", func(t *testing.T) {
		if !isZeroValue(0) {
			t.Error("expected true for zero int")
		}
	})
	t.Run("nonzero int", func(t *testing.T) {
		if isZeroValue(1) {
			t.Error("expected false for nonzero int")
		}
	})
	t.Run("zero struct", func(t *testing.T) {
		type foo struct{ A int }
		if !isZeroValue(foo{}) {
			t.Error("expected true for zero struct")
		}
	})
}

func TestApiRequest_GetUrl(t *testing.T) {
	req := ApiRequest[string, string]{
		Endpoint: "foo/bar",
		Config:   &ApiRequestConfig{BaseURL: "http://localhost"},
	}
	url := req.GetUrl()
	if url != "http://localhost/foo/bar" {
		t.Errorf("unexpected url: %s", url)
	}
}

func TestApiRequest_DoRequest_Success(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"foo":"bar"}`))
	}))
	defer ts.Close()

	req := ApiRequest[struct{}, *map[string]string]{
		Method:   "GET",
		Endpoint: "",
		Config:   &ApiRequestConfig{BaseURL: ts.URL},
	}
	resp, err := req.DoRequest()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if (*resp)["foo"] != "bar" {
		t.Errorf("unexpected response: %+v", resp)
	}
}

func TestApiRequest_DoRequest_Error(t *testing.T) {
	req := ApiRequest[struct{}, struct{}]{
		Method:   "GET",
		Endpoint: "",
		Config:   &ApiRequestConfig{BaseURL: ""},
	}
	_, err := req.DoRequest()
	if err == nil {
		t.Error("expected error for missing BaseURL")
	}
}

func TestApiRequest_DoRequest_Non200(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
	}))
	defer ts.Close()

	req := ApiRequest[struct{}, struct{}]{
		Method:   "GET",
		Endpoint: "",
		Config:   &ApiRequestConfig{BaseURL: ts.URL},
	}
	_, err := req.DoRequest()
	if err == nil {
		t.Error("expected error for non-200 status")
	}
}
