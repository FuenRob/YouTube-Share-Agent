package discord

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
)

// Client maneja la conexión con Discord
type Client struct {
	Token     string
	ChannelID string
}

// NewDiscordClient crea una nueva instancia del cliente de Discord
func NewDiscordClient(token, channelID string) *Client {
	return &Client{
		Token:     token,
		ChannelID: channelID,
	}
}

// SendMessage envía un mensaje al canal configurado
func (c *Client) SendMessage(content string) error {
	// Crea una nueva sesión de Discord
	dg, err := discordgo.New("Bot " + c.Token)
	if err != nil {
		return fmt.Errorf("error creando sesión de Discord: %v", err)
	}

	// Envía el mensaje al canal especificado
	_, err = dg.ChannelMessageSend(c.ChannelID, content)
	if err != nil {
		return fmt.Errorf("error enviando mensaje: %v", err)
	}

	return nil
}
