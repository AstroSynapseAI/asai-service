package server

import (
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/AstroSynapseAI/app-service/app"
	"github.com/AstroSynapseAI/app-service/controllers"
	"github.com/AstroSynapseAI/app-service/controllers/ws"
	"github.com/AstroSynapseAI/app-service/sdk/crud/database"
	"github.com/AstroSynapseAI/app-service/sdk/rest"
	"github.com/bwmarrin/discordgo"
)

type AsaiServer struct {
	discordClient *discordgo.Session
}

func NewAsaiServer() *AsaiServer {
	var err error
	server := &AsaiServer{}

	server.discordClient, err = discordgo.New("Bot " + os.Getenv("DISCORD_API_KEY"))
	if err != nil {
		fmt.Println("Failed to create Discord session:", err)
	}

	return server
}

func (server *AsaiServer) Run(db *database.Database) error {
	// Initialize the Discord client
	discordCtrl := controllers.NewDiscordController(db)
	server.discordClient.AddHandler(discordCtrl.MsgHandler)
	server.discordClient.AddHandler(discordCtrl.NewMemberHandler)
	server.discordClient.Identify.Intents = discordgo.IntentsGuildMessages

	err := server.discordClient.Open()
	if err != nil {
		fmt.Println("Failed to open Discord connection:", err)
		return err
	}

	router := rest.NewRouter()
	router.Mux.StrictSlash(true)

	// Serve API controllers
	routes := app.NewRoutes(router, db)
	router.Load(routes)

	// Serve WebSocket
	wsManager := ws.NewManager(db)
	router.Mux.HandleFunc("/ws/chat", wsManager.Handler)

	// Serve Websites
	webCtrl := controllers.NewWebController(router)
	webCtrl.Run()

	//
	// PORT is defined by heroku env
	port := os.Getenv("PORT")
	
	// If PORT is not defined, server is running locally
	if port == "" {
		err := router.Listen(":8082")
		if err != nil {
			fmt.Println("Failed to listen:", err)
			return err
		}
		return nil
	}

	// If PORT is defined, server is running on Heroku
	// Create a TCP listener with heroku port
	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		fmt.Println("Failed to listen:", err)
		return err
	}

	// Serve the HTTP server
	err = http.Serve(listener, router.Mux)
	if err != nil {
		fmt.Println("Failed to serve:", err)
		return err
	}

	// This need to be revised and probably wrapped in an plugins interface
	// is part of the discord client
	// Setup signal capturing for closing discord connection (I think)
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM)

	// Wait for SIGTERM signal
	<-stop

	// Cleanly close down the Discord session.
	server.discordClient.Close()

	return nil
}
