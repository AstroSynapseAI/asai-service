package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/AstroSynapseAI/engine-service/chains"
	"github.com/AstroSynapseAI/engine-service/config"
	"github.com/GoLangWebSDK/rest"
	"github.com/bwmarrin/discordgo"
	"github.com/gorilla/websocket"
	options "github.com/tmc/langchaingo/chains"
)

var asaiChain *chains.AsaiChain
var discordClient *discordgo.Session

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func init() {
	var err error
	config.LoadEnvironment()

	asaiChain, err = chains.NewAsaiChain()
	if err != nil {
		fmt.Println("Failed to create AsaiChain:", err)
		return
	}

	token := os.Getenv("DISCORD_API_KEY")
	discordClient, err = discordgo.New("Bot " + token)
	if err != nil {
		fmt.Println("Failed to create Discord session:", err)
		return
	}
}

func main() {
	// Initialize the Discord client
	discordClient.AddHandler(DiscordMsgHandler)
	discordClient.Identify.Intents = discordgo.IntentsGuildMessages

	err := discordClient.Open()
	if err != nil {
		fmt.Println("Failed to open Discord connection:", err)
		return
	}

	router := rest.NewRouter()
	ctrl := rest.NewController(router)

	router.Mux.StrictSlash(true)

	ctrl.Post("/api/chat/msg", PostHandler)

	router.Mux.HandleFunc("/api/chat/socket", StreamHandler)

	port := os.Getenv("PORT")
	if port == "" {
		router.Listen(":8082")
		return
	}

	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		fmt.Println("Failed to listen:", err)
		return
	}

	// Start the HTTP server using the router and the listener
	err = http.Serve(listener, router.Mux)
	if err != nil {
		fmt.Println("Failed to serve:", err)
		return
	}

	// Setup signal capturing
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM)

	// Wait for SIGTERM signal
	<-stop

	// Cleanly close down the Discord session.
	discordClient.Close()

}

func PostHandler(ctx *rest.Context) {
	// Parse the incoming http request
	var request struct {
		SessionId  string `json:"session_id"`
		UserPrompt string `json:"user_prompt"`
	}

	err := ctx.JsonDecode(&request)
	if err != nil {
		fmt.Println("Bad Request: %w", err)
		ctx.JsonResponse(400, err)
		return
	}

	asaiChain.SetSessionID(request.SessionId)

	response, err := asaiChain.Run(context.Background(), request.UserPrompt)
	if err != nil {
		fmt.Println(err)
		ctx.JsonResponse(500, err)
		return
	}

	var responseJson struct {
		Content string `json:"content"`
	}

	responseJson.Content = response

	ctx.JsonResponse(200, responseJson)
}

func DiscordMsgHandler(session *discordgo.Session, msg *discordgo.MessageCreate) {
	if msg.Author.ID == session.State.User.ID {
		return
	}

	sessionID := msg.Author.ID
	userPrompt := msg.Content

	if strings.Contains(msg.Content, "@"+session.State.User.ID) {
		asaiChain.SetSessionID(sessionID)
		response, err := asaiChain.Run(context.Background(), userPrompt)
		if err != nil {
			fmt.Println(err)
			return
		}

		session.ChannelMessageSend(msg.ChannelID, response)
	}
}

func StreamHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Failed to initate socket:", err)
		return
	}

	defer conn.Close()

	var request struct {
		SessionId  string `json:"session_id"`
		UserPrompt string `json:"user_prompt"`
	}

	for	{
		msgType, p, err := conn.ReadMessage()
		if err != nil {
			fmt.Println("Failed to read message:", err)
			return
		}

		err = json.Unmarshal(p, &request)
		if err != nil {
			fmt.Println("Failed to parse JSON:", err)
			conn.WriteMessage(websocket.TextMessage, []byte("Failed to parse JSON"))
			continue
		}

		asaiChain.SetSessionID(request.SessionId)

		_, err = asaiChain.Run(context.Background(), 
			request.UserPrompt, 
			options.WithStreamingFunc(func(ctx context.Context, chunk []byte) error { 
				conn.WriteMessage(msgType, chunk)
				return nil 
			}),
		)
		
		if err != nil {
			fmt.Println(err)
			conn.WriteMessage(websocket.TextMessage, []byte("Error during chain run"))
		}
	}
}
