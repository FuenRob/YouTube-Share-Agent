package youtube

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

// Client maneja la comunicación con la API de YouTube
type Client struct {
	APIKey string
}

// NewYouTubeClient crea una nueva instancia del cliente de YouTube
func NewYouTubeClient(apiKey string) *Client {
	return &Client{APIKey: apiKey}
}

// Video representa los datos básicos de un vídeo de YouTube
type Video struct {
	Title       string    `json:"title"`
	URL         string    `json:"url"`
	PublishedAt time.Time `json:"published_at"`
	Description string    `json:"description,omitempty"`
}

// GetLatestVideo obtiene el último vídeo publicado en un canal
func (c *Client) GetLatestVideo(channelID string) (*Video, error) {
	// Construye la URL de la API de YouTube
	url := fmt.Sprintf(
		"https://www.googleapis.com/youtube/v3/search?key=%s&channelId=%s&part=snippet&order=date&maxResults=1&type=video",
		c.APIKey,
		channelID,
	)

	// Realiza la petición HTTP
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error en la petición HTTP: %v", err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println("⚠️ Error al cerrar el cuerpo de la respuesta")
		}
	}(resp.Body)

	// Verifica el código de estado
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API devolvió código %d", resp.StatusCode)
	}

	// Decodifica la respuesta JSON
	var result struct {
		Items []struct {
			ID struct {
				VideoID string `json:"videoId"`
			} `json:"id"`
			Snippet struct {
				Title       string `json:"title"`
				PublishedAt string `json:"publishedAt"`
				Description string `json:"description"`
			} `json:"snippet"`
		} `json:"items"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("error decodificando JSON: %v", err)
	}

	// Verifica si hay resultados
	if len(result.Items) == 0 {
		return nil, fmt.Errorf("no se encontraron vídeos en el canal")
	}

	// Parsea la fecha de publicación
	publishedAt, err := time.Parse(time.RFC3339, result.Items[0].Snippet.PublishedAt)
	if err != nil {
		return nil, fmt.Errorf("error parseando fecha: %v", err)
	}

	// Construye el objeto Video
	video := &Video{
		Title:       result.Items[0].Snippet.Title,
		URL:         fmt.Sprintf("https://youtu.be/%s", result.Items[0].ID.VideoID),
		PublishedAt: publishedAt,
		Description: result.Items[0].Snippet.Description,
	}

	return video, nil
}
