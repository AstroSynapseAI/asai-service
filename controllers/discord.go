package controllers

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/AstroSynapseAI/app-service/engine"
	"github.com/AstroSynapseAI/app-service/engine/chains"
	"github.com/AstroSynapseAI/app-service/sdk/crud/database"
	"github.com/bwmarrin/discordgo"
)

type DiscordController struct {
	db          *database.Database
	asaiChain   *chains.AsaiChain
	WelcomeChID string
	ClientType  string
}

func NewDiscordController(db *database.Database) *DiscordController {
	ctrl := &DiscordController{
		db:          db,
		WelcomeChID: os.Getenv("WELCOME_CHANNEL_ID"),
		ClientType:  "Discord",
	}

	asaiConfig := engine.NewConfig(db)
	asaiChain, err := chains.NewAsaiChain(asaiConfig)
	if err != nil {
		fmt.Println("Failed to initate socket:", err)
		return nil
	}

	ctrl.asaiChain = asaiChain
	ctrl.asaiChain.SetClientType(ctrl.ClientType)

	return ctrl
}

func (ctrl *DiscordController) MsgHandler(session *discordgo.Session, msg *discordgo.MessageCreate) {
	if msg.Author.ID == session.State.User.ID {
		return
	}

	sessionID := msg.Author.ID
	userPrompt := msg.Content

	if strings.Contains(msg.Content, "@"+session.State.User.ID) {
		ctrl.asaiChain.SetSessionID(sessionID)

		response, err := ctrl.asaiChain.Prompt(context.Background(), userPrompt)
		if err != nil {
			fmt.Println(err)
			return
		}

		_, _ = session.ChannelMessageSend(msg.ChannelID, response)
	}
}

func (ctrl *DiscordController) NewMemberHandler(session *discordgo.Session, addEvent *discordgo.GuildMemberAdd) {
	sessionID := addEvent.User.ID
	userName := addEvent.User.Username

	ctrl.asaiChain.SetSessionID(sessionID)

	userPrompt := fmt.Sprintf("New user, %s (%s), has joined the server.", userName, sessionID)

	response, err := ctrl.asaiChain.Prompt(context.Background(), userPrompt)
	if err != nil {
		fmt.Println(err)
		return
	}

	_, _ = session.ChannelMessageSend(ctrl.WelcomeChID, response)
}
