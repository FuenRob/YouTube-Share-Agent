package youtube

import (
	"io"
	"net/http"
	"strings"
	"testing"
	"time"
)

// roundTripFunc es un tipo auxiliar que nos permite definir una función
// que se comporta como un http.RoundTripper.
type roundTripFunc func(req *http.Request) (*http.Response, error)

// RoundTrip implementa la interfaz http.RoundTripper.
func (f roundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req)
}

// TestGetLatestVideoSuccess simula una respuesta exitosa de la API de YouTube.
func TestGetLatestVideoSuccess(t *testing.T) {
	// Guardamos el transporte original para restaurarlo después.
	originalTransport := http.DefaultTransport
	defer func() { http.DefaultTransport = originalTransport }()

	// Creamos un transporte que simula una respuesta exitosa.
	http.DefaultTransport = roundTripFunc(func(req *http.Request) (*http.Response, error) {
		// Definimos una respuesta JSON simulada.
		jsonResponse := `{
			"items": [
				{
					"id": {"videoId": "abc123"},
					"snippet": {
						"title": "Test Video",
						"publishedAt": "2023-02-18T12:34:56Z",
						"description": "This is a test video."
					}
				}
			]
		}`

		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(strings.NewReader(jsonResponse)),
			Header:     make(http.Header),
		}, nil
	})

	client := NewYouTubeClient("fake-api-key")
	video, err := client.GetLatestVideo("fake-channel-id")
	if err != nil {
		t.Fatalf("Se esperaba respuesta exitosa, pero se obtuvo error: %v", err)
	}

	if video.Title != "Test Video" {
		t.Errorf("Se esperaba título 'Test Video', pero se obtuvo: %s", video.Title)
	}
	if video.URL != "https://youtu.be/abc123" {
		t.Errorf("Se esperaba URL 'https://youtu.be/abc123', pero se obtuvo: %s", video.URL)
	}

	expectedTime, _ := time.Parse(time.RFC3339, "2023-02-18T12:34:56Z")
	if !video.PublishedAt.Equal(expectedTime) {
		t.Errorf("Se esperaba fecha de publicación %v, pero se obtuvo: %v", expectedTime, video.PublishedAt)
	}
	if video.Description != "This is a test video." {
		t.Errorf("Se esperaba descripción 'This is a test video.', pero se obtuvo: %s", video.Description)
	}
}

// TestGetLatestVideoHTTPError simula un error HTTP (por ejemplo, código 500).
func TestGetLatestVideoHTTPError(t *testing.T) {
	originalTransport := http.DefaultTransport
	defer func() { http.DefaultTransport = originalTransport }()

	// El transporte simula una respuesta con código 500.
	http.DefaultTransport = roundTripFunc(func(req *http.Request) (*http.Response, error) {
		return &http.Response{
			StatusCode: http.StatusInternalServerError,
			Body:       io.NopCloser(strings.NewReader("Internal Server Error")),
			Header:     make(http.Header),
		}, nil
	})

	client := NewYouTubeClient("fake-api-key")
	_, err := client.GetLatestVideo("fake-channel-id")
	if err == nil {
		t.Fatal("Se esperaba error debido a código HTTP 500, pero no se obtuvo ninguno")
	}
	if !strings.Contains(err.Error(), "API devolvió código") {
		t.Errorf("El mensaje de error no es el esperado: %v", err)
	}
}

// TestGetLatestVideoInvalidJSON simula una respuesta con JSON inválido.
func TestGetLatestVideoInvalidJSON(t *testing.T) {
	originalTransport := http.DefaultTransport
	defer func() { http.DefaultTransport = originalTransport }()

	http.DefaultTransport = roundTripFunc(func(req *http.Request) (*http.Response, error) {
		// Retornamos un cuerpo que no es un JSON válido.
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(strings.NewReader("not a valid json")),
			Header:     make(http.Header),
		}, nil
	})

	client := NewYouTubeClient("fake-api-key")
	_, err := client.GetLatestVideo("fake-channel-id")
	if err == nil {
		t.Fatal("Se esperaba error al decodificar JSON, pero no se obtuvo ninguno")
	}
	if !strings.Contains(err.Error(), "error decodificando JSON") {
		t.Errorf("El mensaje de error no es el esperado: %v", err)
	}
}
