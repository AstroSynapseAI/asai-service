package server

import (
	"fmt"
	"net/http"
	"os"

	"github.com/AstroSynapseAI/asai-service/app"
	"github.com/AstroSynapseAI/asai-service/controllers"
	"github.com/AstroSynapseAI/asai-service/plugins"
	"github.com/AstroSynapseAI/asai-service/sdk/rest"
	"github.com/GoLangWebSDK/crud/database"
	"github.com/bwmarrin/discordgo"
)

type AsaiServer struct {
	Plugins       *plugins.PluginLoader
	discordClient *discordgo.Session
}

func NewAsaiServer() *AsaiServer {
	server := &AsaiServer{
		Plugins: plugins.NewLoader(),
	}

	return server
}

func (server *AsaiServer) Run(db *database.Database) error {
	// server.Plugins.LoadConfig(db)

	router := rest.NewRouter()
	router.Mux.StrictSlash(true)

	// Serve API controllers
	router.Load(app.NewRoutes(db))

	// Serve WebSocket
	// wsManager := ws.NewManager(db)
	// router.Mux.HandleFunc("/ws/chat", wsManager.Handler)

	// Serve Websites
	webCtrl := controllers.NewWebController(router)
	webCtrl.Run()

	port := os.Getenv("PORT")

	if port == "" {
		// set default port 8080, assume local server
		port = "8080"
	}

	err := http.ListenAndServe(":"+port, router.Mux)
	if err != nil {
		fmt.Println("Failed to serve:", err)
		return err
	}

	// Open all plugins connections
	// server.Plugins.OpenConnection(db)

	return nil
}
