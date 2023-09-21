package main

import (
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_getRequests(t *testing.T) {
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		number := r.URL.Query().Get("number")
		switch number {
		case "1":
			w.WriteHeader(http.StatusOK)
		case "3", "4":
			t.Errorf("%s: should not be called", r.RequestURI)
			w.WriteHeader(http.StatusOK)
		}
	}))
	urls := []string{
		s.URL + "?number=1",
		"invalid url",
		s.URL + "?number=3",
		s.URL + "?number=4",
	}
	defer s.Close()

	res := getRequests(http.DefaultClient, urls)
	require.Equal(t, 1, len(res))
	require.Equal(t, 200, res[0].Status)
	require.Equal(t, 1, res[0].Number)
	require.Equal(t, urls[0], res[0].URL)
}
