package selectconcept

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func makeDelayedServer(duration time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(duration)
		w.WriteHeader(http.StatusOK)
	}))
}

func TestRacer(t *testing.T) {
	t.Run("compare speed of servers", func(t *testing.T) {
		slowServer := makeDelayedServer(20 * time.Millisecond)
		fastServer := makeDelayedServer(0 * time.Millisecond)
		defer slowServer.Close()
		defer fastServer.Close()
		slowURL := slowServer.URL
		fastURL := fastServer.URL

		want := fastURL

		got, _ := Racer(slowURL, fastURL)

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
	t.Run("returns an error if server doesn't responds in 10 seconds", func(t *testing.T) {
		server := makeDelayedServer(time.Millisecond * 11)
		defer server.Close()

		_, err := ConfigurableRacer(server.URL, server.URL, time.Millisecond*10	)
		if err == nil {
			t.Error("Expected an error but got nil")
		}
	})
}
