package main

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/AstroSynapseAI/engine-service/config"
	"github.com/AstroSynapseAI/engine-service/cortex/chains"
	api "github.com/AstroSynapseAI/engine-service/servers/rest"
	"github.com/AstroSynapseAI/engine-service/servers/ws"
	"github.com/GoLangWebSDK/rest"
	"github.com/bwmarrin/discordgo"
)

var (
	asaiChain     *chains.AsaiChain
	discordClient *discordgo.Session
	wsManager     *ws.Manager
)

func init() {
	var err error

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

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

	wsManager = ws.NewManager(ctx)
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
	router.Mux.StrictSlash(true)

	// Web client server
	static := http.FileServer(http.Dir("./servers/static"))
	assets := http.FileServer(http.Dir("./servers/static/assets"))

	router.Mux.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", assets))
	router.Mux.Handle("/", static)

	// API server
	ctrl := rest.NewController(router)
	ctrl.Get("/api/chat/history/{session_id}", api.GetHistory)
	ctrl.Post("/api/chat/msg", PostHandler)
	ctrl.Get("/api/users/session", api.GetSession)

	// Websocket server
	router.Mux.HandleFunc("/api/chat/socket", wsManager.Handler)

	port := os.Getenv("PORT")
	if port == "" {
		err := router.Listen(":8082")
		if err != nil {
			fmt.Println("Failed to listen:", err)
			return
		}
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
		_ = ctx.JsonResponse(400, err)
		return
	}

	asaiChain.SetSessionID(request.SessionId)

	response, err := asaiChain.Run(context.Background(), request.UserPrompt)
	if err != nil {
		fmt.Println(err)
		_ = ctx.JsonResponse(500, err)
		return
	}

	var responseJson struct {
		Content string `json:"content"`
	}

	responseJson.Content = response

	_ = ctx.JsonResponse(200, responseJson)
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

		_, _ = session.ChannelMessageSend(msg.ChannelID, response)
	}
}
