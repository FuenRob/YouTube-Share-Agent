package main

import (
	"fmt"
	"log"
	"youtube-share-agent/config"
	"youtube-share-agent/internal/deepseek"
	"youtube-share-agent/internal/discord"
	"youtube-share-agent/internal/youtube"
)

func main() {
	// Cargar configuración
	cfg := config.LoadConfig()

	// Inicializar clientes
	ytClient, discordClient := youtube.NewYouTubeClient(cfg.YouTubeAPIKey), discord.NewDiscordClient(cfg.DiscordToken, cfg.ChannelID)

	// 1. Obtener último vídeo de YouTube
	video, err := ytClient.GetLatestVideo(cfg.YouTubeChannelID)
	if err != nil {
		log.Fatalf("🚨 Error obteniendo el último vídeo: %v", err)
	}

	log.Printf("✅ Vídeo obtenido: %s", video.Title)

	// 2. Generar mensaje con DeepSeek
	var message string
	if cfg.DeepSeekAPIKey != "" {
		dsClient := deepseek.NewClient(cfg.DeepSeekAPIKey)
		generatedMsg, err := dsClient.GenerateMotivationalMessage(video.Title, video.Description)

		if err != nil {
			log.Printf("⚠️ No se pudo generar mensaje con DeepSeek: %v", err)
			message = createDefaultMessage(video)
		} else {
			message = fmt.Sprintf("%s\n\n🔗 Enlace: %s\n\n @everyone", generatedMsg, video.URL)
			log.Println("✨ Mensaje generado con DeepSeek")
		}
	} else {
		message = createDefaultMessage(video)
		log.Println("ℹ️ Mensaje predeterminado (sin DeepSeek)")
	}

	// 3. Enviar a Discord
	if err := discordClient.SendMessage(message); err != nil {
		log.Fatalf("🚨 Error enviando a Discord: %v", err)
	}

	log.Println("🎉 Mensaje enviado exitosamente a Discord!")
	fmt.Println(message) // Mostrar el mensaje en consola
}

// createDefaultMessage genera un mensaje básico si falla DeepSeek
func createDefaultMessage(v *youtube.Video) string {
	return fmt.Sprintf(
		"🎥 **Nuevo vídeo disponible!**\n\n"+
			"**%s**\n\n"+
			"📅 Publicado: %s\n"+
			"🔗 Enlace: %s\n\n"+
			"¡No te lo pierdas! 👀\n\n @everyone",
		v.Title,
		v.PublishedAt.Format("02/01/2006 15:04"),
		v.URL,
	)
}
