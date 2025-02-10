package deepseek

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Client struct {
	APIKey string
	Model  string // Ej: "deepseek-chat"
}

func NewClient(apiKey string) *Client {
	return &Client{
		APIKey: apiKey,
		Model:  "deepseek-chat", // Verifica el modelo actual en la documentación
	}
}

type MessageRequest struct {
	Model    string `json:"model"`
	Messages []struct {
		Role    string `json:"role"`
		Content string `json:"content"`
	} `json:"messages"`
	Temperature float32 `json:"temperature"` // Controla la creatividad (0-2)
}

type MessageResponse struct {
	Choices []struct {
		Message struct {
			Content string `json:"content"`
		} `json:"message"`
	} `json:"choices"`
}

// GenerateMotivationalMessage genera un mensaje atractivo basado en el título del vídeo
func (c *Client) GenerateMotivationalMessage(videoTitle string, videoDescription ...string) (string, error) {
	prompt := fmt.Sprintf(`
        Como experto en marketing digital y community management, genera un mensaje motivacional 
        para anunciar un nuevo vídeo de YouTube en Discord. El mensaje debe:
        
        1. Ser entusiasta pero profesional
        2. Incluir emojis relevantes (máximo 3)
        3. Destacar el valor principal del vídeo
        4. Terminar con un llamado a la acción
        5. Máximo 2 párrafos cortos
        
        Título del vídeo: "%s"
        Descripción: "%s"
        
        Formato requerido:
        [Emoji relacionado] [Texto motivacional creativo]
        [Llamado a la acción con emoji]
        [Enlace] (el enlace lo añadiremos después)
    `, videoTitle, firstOrEmpty(videoDescription))

	requestBody := MessageRequest{
		Model: c.Model,
		Messages: []struct {
			Role    string `json:"role"`
			Content string `json:"content"`
		}{
			{"user", prompt},
		},
		Temperature: 0.7, // Balance entre creatividad y enfoque
	}

	body, _ := json.Marshal(requestBody)
	req, _ := http.NewRequest("POST", "https://api.deepseek.com/v1/chat/completions", bytes.NewBuffer(body))
	req.Header.Set("Authorization", "Bearer "+c.APIKey)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("error en la petición a DeepSeek: %v", err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println("⚠️ Error al cerrar el cuerpo de la respuesta")
		}
	}(resp.Body)

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("API DeepSeek devolvió código %d", resp.StatusCode)
	}

	var response MessageResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return "", fmt.Errorf("error decodificando respuesta: %v", err)
	}

	if len(response.Choices) == 0 {
		return "", fmt.Errorf("no se generó contenido")
	}

	return response.Choices[0].Message.Content, nil
}

// Helper para manejar parámetro opcional de descripción
func firstOrEmpty(s []string) string {
	if len(s) > 0 {
		return s[0]
	}
	return ""
}
