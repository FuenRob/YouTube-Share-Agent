package discord

import (
	"strings"
	"testing"
)

// TestNewDiscordClient verifica que al crear el cliente se asignen correctamente el token y el channelID.
func TestNewDiscordClient(t *testing.T) {
	token := "mi-token-falso"
	channelID := "mi-canal-falso"
	client := NewDiscordClient(token, channelID)

	if client.Token != token {
		t.Errorf("Se esperaba token %s, pero se obtuvo %s", token, client.Token)
	}
	if client.ChannelID != channelID {
		t.Errorf("Se esperaba channelID %s, pero se obtuvo %s", channelID, client.ChannelID)
	}
}

// TestSendMessage verifica el comportamiento del método SendMessage.
// Debido a que usamos datos falsos, se espera que la llamada a la API falle y se devuelva un error.
// Este test nos ayuda a comprobar que se captura y se formatea correctamente el error.
func TestSendMessage(t *testing.T) {
	token := "mi-token-falso"
	channelID := "mi-canal-falso"
	client := NewDiscordClient(token, channelID)

	err := client.SendMessage("Mensaje de prueba")
	if err == nil {
		t.Error("Se esperaba un error al enviar el mensaje con token y canal falsos, pero no se produjo ningún error")
	} else {
		// Comprobamos que el mensaje de error contenga alguna de las frases esperadas.
		if !strings.Contains(err.Error(), "error enviando mensaje") && !strings.Contains(err.Error(), "error creando sesión") {
			t.Errorf("Mensaje de error inesperado: %s", err.Error())
		}
	}
}
