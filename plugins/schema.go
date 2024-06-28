package plugins

import (
	"github.com/GoLangWebSDK/crud/database"
)

type Plugins interface {
	LoadConfig(db *database.Database) error
	OpenConnection(db *database.Database) error
}

type Plugin interface {
	OpenConnection(db *database.Database)
}
