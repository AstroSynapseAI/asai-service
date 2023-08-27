package main

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/AstroSynapseAI/engine-service/chains"
	"github.com/AstroSynapseAI/engine-service/config"
	"github.com/GoLangWebSDK/rest"
	"github.com/bwmarrin/discordgo"
)

var asaiChain *chains.AsaiChain
var discordClient *discordgo.Session

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

	asaiChain.SetSessionID(sessionID)
	response, err := asaiChain.Run(context.Background(), userPrompt)
	if err != nil {
		fmt.Println(err)
		return
	}

	session.ChannelMessageSend(msg.ChannelID, response)
}
