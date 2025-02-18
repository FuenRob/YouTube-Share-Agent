package deepseek

import (
	"io"
	"net/http"
	"strings"
	"testing"
)

// roundTripFunc nos permite definir una función que actúa como un http.RoundTripper.
type roundTripFunc func(req *http.Request) (*http.Response, error)

// RoundTrip implementa la interfaz http.RoundTripper.
func (f roundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req)
}

// TestGenerateMotivationalMessageSuccess simula una respuesta exitosa de la API DeepSeek.
func TestGenerateMotivationalMessageSuccess(t *testing.T) {
	// Guardamos el transporte original para restaurarlo tras el test.
	originalTransport := http.DefaultTransport
	defer func() { http.DefaultTransport = originalTransport }()

	// Configuramos un transporte personalizado que simula la respuesta exitosa.
	http.DefaultTransport = roundTripFunc(func(req *http.Request) (*http.Response, error) {
		// Verificamos que se use el método POST y la URL correcta.
		if req.Method != "POST" {
			t.Errorf("Se esperaba método POST, se obtuvo %s", req.Method)
		}
		if req.URL.String() != "https://api.deepseek.com/v1/chat/completions" {
			t.Errorf("Se esperaba URL https://api.deepseek.com/v1/chat/completions, se obtuvo %s", req.URL.String())
		}
		// Verificamos los encabezados.
		if auth := req.Header.Get("Authorization"); auth != "Bearer fake-api-key" {
			t.Errorf("Se esperaba Authorization 'Bearer fake-api-key', se obtuvo %s", auth)
		}
		if ct := req.Header.Get("Content-Type"); ct != "application/json" {
			t.Errorf("Se esperaba Content-Type 'application/json', se obtuvo %s", ct)
		}

		// Retornamos una respuesta JSON simulada.
		responseJSON := `{
			"choices": [
				{
					"message": {
						"content": "¡Este vídeo te inspirará a alcanzar nuevas metas!"
					}
				}
			]
		}`
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(strings.NewReader(responseJSON)),
			Header:     make(http.Header),
		}, nil
	})

	client := NewClient("fake-api-key")
	message, err := client.GenerateMotivationalMessage("Título de prueba", "Descripción de prueba")
	if err != nil {
		t.Fatalf("Se esperaba éxito, pero se obtuvo error: %v", err)
	}

	expected := "¡Este vídeo te inspirará a alcanzar nuevas metas!"
	if message != expected {
		t.Errorf("Se esperaba mensaje %q, se obtuvo %q", expected, message)
	}
}

// TestGenerateMotivationalMessageHTTPError simula un error HTTP (por ejemplo, código 500).
func TestGenerateMotivationalMessageHTTPError(t *testing.T) {
	originalTransport := http.DefaultTransport
	defer func() { http.DefaultTransport = originalTransport }()

	http.DefaultTransport = roundTripFunc(func(req *http.Request) (*http.Response, error) {
		return &http.Response{
			StatusCode: http.StatusInternalServerError,
			Body:       io.NopCloser(strings.NewReader("Internal Server Error")),
			Header:     make(http.Header),
		}, nil
	})

	client := NewClient("fake-api-key")
	_, err := client.GenerateMotivationalMessage("Título de prueba")
	if err == nil {
		t.Fatal("Se esperaba error debido al código HTTP, pero no se obtuvo ninguno")
	}
	if !strings.Contains(err.Error(), "API DeepSeek devolvió código") {
		t.Errorf("Mensaje de error inesperado: %v", err)
	}
}

// TestGenerateMotivationalMessageInvalidJSON simula una respuesta con JSON inválido.
func TestGenerateMotivationalMessageInvalidJSON(t *testing.T) {
	originalTransport := http.DefaultTransport
	defer func() { http.DefaultTransport = originalTransport }()

	http.DefaultTransport = roundTripFunc(func(req *http.Request) (*http.Response, error) {
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(strings.NewReader("invalid json")),
			Header:     make(http.Header),
		}, nil
	})

	client := NewClient("fake-api-key")
	_, err := client.GenerateMotivationalMessage("Título de prueba")
	if err == nil {
		t.Fatal("Se esperaba error al decodificar JSON, pero no se obtuvo ninguno")
	}
	if !strings.Contains(err.Error(), "error decodificando respuesta") {
		t.Errorf("Mensaje de error inesperado: %v", err)
	}
}
