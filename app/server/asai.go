package server

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/AstroSynapseAI/app-service/controllers"
	"github.com/AstroSynapseAI/app-service/controllers/ws"
	"github.com/AstroSynapseAI/app-service/engine/chains"
	"github.com/GoLangWebSDK/rest"
	"github.com/bwmarrin/discordgo"
)

type AsaiServer struct {
	discordClient *discordgo.Session
	wsManager     *ws.Manager
	asaiChain     *chains.AsaiChain
}

func NewAsaiServer() *AsaiServer {
	var err error
	server := &AsaiServer{}

	server.asaiChain, err = chains.NewAsaiChain()
	if err != nil {
		fmt.Println("Failed to create AsaiChain:", err)
	}

	token := os.Getenv("DISCORD_API_KEY")
	server.discordClient, err = discordgo.New("Bot " + token)
	if err != nil {
		fmt.Println("Failed to create Discord session:", err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	server.wsManager = ws.NewManager(ctx)

	return server
}

func (server *AsaiServer) Run() error {
	// Initialize the Discord client
	server.discordClient.AddHandler(controllers.DiscordMsgHandler)
	server.discordClient.Identify.Intents = discordgo.IntentsGuildMessages

	err := server.discordClient.Open()
	if err != nil {
		fmt.Println("Failed to open Discord connection:", err)
		return err
	}

	router := rest.NewRouter()
	router.Mux.StrictSlash(true)

	// Web client server
	static := http.FileServer(http.Dir("./web/static"))
	assets := http.FileServer(http.Dir("./web/static/assets"))

	router.Mux.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", assets))
	router.Mux.Handle("/", static)

	// API server
	ctrl := rest.NewController(router)
	ctrl.Get("/api/chat/history/{session_id}", controllers.GetHistory)
	ctrl.Post("/api/chat/msg", controllers.PostHandler)
	ctrl.Get("/api/users/session", controllers.GetSession)

	// Websocket server
	router.Mux.HandleFunc("/api/chat/socket", server.wsManager.Handler)

	port := os.Getenv("PORT")
	if port == "" {
		err := router.Listen(":8082")
		if err != nil {
			fmt.Println("Failed to listen:", err)
			return err
		}
		return nil
	}

	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		fmt.Println("Failed to listen:", err)
		return err
	}

	// Start the HTTP server using the router and the listener
	err = http.Serve(listener, router.Mux)
	if err != nil {
		fmt.Println("Failed to serve:", err)
		return err
	}

	// Setup signal capturing
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM)

	// Wait for SIGTERM signal
	<-stop

	// Cleanly close down the Discord session.
	server.discordClient.Close()

	return nil

}
