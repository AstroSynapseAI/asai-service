package plugins

import (
	"fmt"

	"github.com/AstroSynapseAI/app-service/engine/plugins/discord"
	"github.com/AstroSynapseAI/app-service/repositories"
	"github.com/AstroSynapseAI/app-service/sdk/crud/database"
)

// I need to get all the plugins platform supports and go through them.

// Then we need to get all the active plugins for each of the plugins and load thir api tokens.

// Finally we need to add plugin sepcific hadlers for all the stuff plugin supports, and open the connection for each plugin.

type PluginLoader struct {
	Plugins []Plugin
}

var _ Plugins = (*PluginLoader)(nil)

func NewLoader() *PluginLoader {
	return &PluginLoader{}
}

func (loader *PluginLoader) LoadConfig(db *database.Database) error {
	pluginsRepo := repositories.NewPluginsRepository(db)
	pluginsRecords, err := pluginsRepo.Repo.ReadAll()
	if err != nil {
		fmt.Println(err)
		return err
	}

	for _, pluginRecord := range pluginsRecords {
		if pluginRecord.Slug == "discord" {
			discordPlugin := discord.NewDiscordPlugin(pluginRecord)
			loader.Plugins = append(loader.Plugins, discordPlugin)
		}
	}

	return nil
}

func (loader *PluginLoader) OpenConnection(db *database.Database) error {

	for _, plugin := range loader.Plugins {
		plugin.OpenConnection(db)
	}

	return nil
}
