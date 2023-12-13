package server

import (
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
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
	router := rest.NewRouter()
	router.Mux.StrictSlash(true)

	// API server
	routes := app.NewRoutes(router, db)
	router.Load(routes)

	// Websocket server
	wsManager := ws.NewManager(db)
	router.Mux.HandleFunc("/ws/chat", wsManager.Handler)

	// Web client server
	staticDir := "./web/static"
	static := http.FileServer(http.Dir(staticDir))
	assets := http.FileServer(http.Dir(staticDir + "/assets"))

	router.Mux.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", assets))
	router.Mux.Handle("/", static)

	// Fallback to index.html for Vue Router
	router.Mux.PathPrefix("/").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		path := filepath.Join(staticDir + r.URL.Path)
		_, err := os.Stat(path)

		if os.IsNotExist(err) {
			http.ServeFile(w, r, staticDir+"/index.html")
			return
		}

		// If request is not for a directory, serve with the static file server as normal
		static.ServeHTTP(w, r)
		return
	})

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

	// Setup signal capturing
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM)

	// Wait for SIGTERM signal
	<-stop

	// Cleanly close down the Discord session.
	server.discordClient.Close()

	//
	// Start the HTTP server using the router and the listener
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

	// Serve the HTTP server
	err = http.Serve(listener, router.Mux)
	if err != nil {
		fmt.Println("Failed to serve:", err)
		return err
	}

	return nil
}
