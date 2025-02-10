# YouTube-Discord Agent 🤖🎥

![Go](https://img.shields.io/badge/Go-1.20+-00ADD8?logo=go&logoColor=white)
![Discord](https://img.shields.io/badge/Discord-Bot-5865F2?logo=discord&logoColor=white)
![YouTube](https://img.shields.io/badge/YouTube-API-FF0000?logo=youtube&logoColor=white)
![DeepSeek](https://img.shields.io/badge/DeepSeek-LLM-00A67E?logo=openai&logoColor=white)

Un agente en Go que obtiene el último vídeo de tu canal de YouTube, genera un mensaje motivacional usando DeepSeek, y lo publica en un canal de Discord. ¡Automatiza la promoción de tus vídeos con estilo! 🚀

---

## Características ✨

- **Obtén el último vídeo de YouTube**: Usa la API de YouTube Data v3 para obtener el vídeo más reciente de tu canal.
- **Genera mensajes motivacionales**: Integra DeepSeek para crear mensajes personalizados que aumenten el engagement.
- **Publica en Discord**: Envía automáticamente el mensaje a un canal de Discord usando un bot.
- **Manejo de errores robusto**: Fallback a mensajes predeterminados si falla DeepSeek.
- **Fácil de configurar**: Solo necesitas tus claves de API y tokens.

---

## Requisitos 📋

- **Go 1.20+**: [Descargar e instalar Go](https://golang.org/dl/)
- **Cuenta de Google Cloud**: Para obtener la API Key de YouTube.
- **Bot de Discord**: Crea un bot en el [Portal de Desarrolladores de Discord](https://discord.com/developers/applications).
- **API Key de DeepSeek**: Opcional, para generar mensajes personalizados.

---

## Configuración ⚙️

1. **Clona el repositorio**:
   ```bash
   git clone https://github.com/tuusuario/youtube-discord-agent.git
   cd youtube-discord-agent
   ```
2. **Configura las credenciales**: Crea un archivo `.env` en la raíz del proyecto.
      - Agrega las credenciales de YouTube y Discord:
        ```env
        API_KEY_YOUTUBE=your_youtube_api_key
        ID_CHANNEL_YOUTUBE=your_youtube_channel_id
        TOKEN_DISCORD=your_discord_bot_token
        ID_CHANNEL_DISCORD=your_discord_channel_id
        API_KEY_DEEPSEEK=your_deepseek_api_key
        ```
      - Si no sé quiere usar DeepSeek, deja el campo vacío en tu env.


3. **Instala las dependencias**: Ejecuta el siguiente comando para instalar las dependencias necesarias.
   ```bash
    go mod tidy
   ```

4. **Ejecuta el proyecto**: Finalmente, ejecuta el agente con el siguiente comando.
   ```bash
   go run main.go
   ```

## Ejemplo de Uso 🚀

1. **Obtener el último vídeo de YouTube**:
   ```json
   {
     "title": "¡Nuevo vídeo! 🎥",
     "description": "¡Hola, amigos! Hoy les traigo un nuevo vídeo sobre...",
     "PublishedAt": "2022-01-01T00:00:00Z",
     "url": "https://www.youtube.com/watch?v=video_id"
   }
   ```

2. **Mensaje generado por DeepSeek**:
   ```json
   {
     "message": "¡Hola, amigos! Hoy les traigo un nuevo vídeo sobre... ¡No se lo pierdan! 🚀"
   }
   ```

3. **Mensaje publicado en Discord**:
   ```markdown
    🎥 **¡Nuevo vídeo!** 🎥
    ¡Hola, amigos! Hoy les traigo un nuevo vídeo sobre... ¡No se lo pierdan! 🚀
    Publicado el 01 de Enero de 2022.
    [Ver vídeo](https://www.youtube.com/watch?v=video_id)
    ```

## Contribuir 🤝
¡Las contribuciones son bienvenidas! Si tienes ideas para mejorar el proyecto, sigue estos pasos:
1. Haz un fork del repositorio. 
2. Crea una rama con tu feature (git checkout -b feature/nueva-funcionalidad). 
3. Haz commit de tus cambios (git commit -m 'Añade nueva funcionalidad'). 
4. Haz push a la rama (git push origin feature/nueva-funcionalidad). 
5. Abre una Pull Request.

## Licencia 📄
Este proyecto está bajo la Licencia MIT - mira el archivo [LICENSE](LICENSE) para detalles.