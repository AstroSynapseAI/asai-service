package controllers

import (
	"os"

	"github.com/GoLangWebSDK/crud/database"
	"github.com/bwmarrin/discordgo"
)

type DiscordController struct {
	db          *database.Database
	WelcomeChID string
	ClientType  string
}

func NewDiscordController(db *database.Database) *DiscordController {
	ctrl := &DiscordController{
		db:          db,
		WelcomeChID: os.Getenv("WELCOME_CHANNEL_ID"),
		ClientType:  "Discord",
	}

	return ctrl
}

func (ctrl *DiscordController) MsgHandler(session *discordgo.Session, msg *discordgo.MessageCreate) {
	// asaiChain := chains.NewAsaiChain(ctrl.db)
	//
	// if msg.Author.ID == session.State.User.ID {
	// 	return
	// }
	//
	// avatarID := uint(1) //tmp hardcoded need to find a way to have a refernce for each plugin to know what avatar is connected to it
	// asaiChain.LoadAvatar(avatarID, msg.Author.ID, ctrl.ClientType)
	//
	// if strings.Contains(msg.Content, "@"+session.State.User.ID) {
	// 	response, err := asaiChain.Prompt(context.Background(), msg.Content)
	// 	if err != nil {
	// 		fmt.Println(err)
	// 		return
	// 	}
	//
	// 	_, _ = session.ChannelMessageSend(msg.ChannelID, response)
	// }
}

func (ctrl *DiscordController) NewMemberHandler(session *discordgo.Session, addEvent *discordgo.GuildMemberAdd) {
	// Stoping this for now, need to implement it properly

	// asaiChain, err := chains.NewAsaiChain(engine.NewConfig(ctrl.db))
	// if err != nil {
	// 	fmt.Println("Failed to initate asai chain:", err)
	// 	return
	// }

	// userPrompt := fmt.Sprintf("New user, %s (%s), has joined the server.", userName, sessionID)

	// response, err := ctrl.asaiChain.Prompt(context.Background(), userPrompt)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// _, _ = session.ChannelMessageSend(ctrl.WelcomeChID, response)
}
