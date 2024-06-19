package discord

import (
	"github.com/bwmarrin/discordgo"
)

type AvatarHandler struct {
	avatarID   uint
	clientType string
}

func NewAvatarHandler(avatarID uint) *AvatarHandler {
	return &AvatarHandler{
		avatarID:   avatarID,
		clientType: "Discord",
	}
}

func (handler *AvatarHandler) Messages(session *discordgo.Session, msg *discordgo.MessageCreate) {
	if msg.Author.ID == session.State.User.ID {
		return
	}

	// handler.asaiChain.LoadAvatar(handler.avatarID, msg.Author.ID, handler.clientType)
	//
	// if strings.Contains(msg.Content, "@"+session.State.User.ID) {
	// 	response, err := handler.asaiChain.Prompt(context.Background(), msg.Content)
	// 	if err != nil {
	// 		fmt.Println(err)
	// 		return
	// 	}
	//
	// 	_, _ = session.ChannelMessageSend(msg.ChannelID, response)
	// }
}
