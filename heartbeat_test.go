package heartbeat_test

import (
	"heartbeat"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/goadesign/goa"
)

func TestHeartbeat_DefaultURL(t *testing.T) {
	service := goa.New("API")
	heartbeat.Heartbeat(service, "")
	server := httptest.NewServer(service.Mux)

	res, err := http.Get(server.URL + "/health")
	if err != nil {
		t.Fatalf("Server error %s", err)
	}
	if res.StatusCode != http.StatusOK {
		t.Fatalf("Expected a 200 OK response, got %d", res.StatusCode)
	}

	server.Close()
}

func TestHeartbeat_CustomURL(t *testing.T) {
	heartbeatURL := "/custom"

	service := goa.New("API")
	heartbeat.Heartbeat(service, heartbeatURL)
	server := httptest.NewServer(service.Mux)

	res, err := http.Get(server.URL + heartbeatURL)
	if err != nil {
		t.Fatalf("Server error %s", err)
	}
	if res.StatusCode != http.StatusOK {
		t.Fatalf("Expected a 200 OK response, got %d", res.StatusCode)
	}

	server.Close()
}

func TestHeartbeat_404URL(t *testing.T) {
	service := goa.New("API")
	heartbeat.Heartbeat(service, "")
	server := httptest.NewServer(service.Mux)

	res, err := http.Get(server.URL + "/other-url-gives-a-404")
	if err != nil {
		t.Fatalf("Server error %s", err)
	}
	if res.StatusCode != http.StatusNotFound {
		t.Fatalf("Expected a 404 OK response, got %d", res.StatusCode)
	}

	server.Close()
}
