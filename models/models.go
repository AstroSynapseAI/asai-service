package models

// I keep my models in a single file when I need to work on table relationships,
// makes it easier to manage when you have overview of all models in a single file.

// type User struct {
// 	gorm.Model
// 	ID 			 int
// 	Email    string
// 	Password string
// 	Username string
// 	AvatarID int
// 	Avatar   *Avatar
// 	Avatars  []Avatar `gorm:"many2many:user_avatars;"`
// 	Role     *Role
// }

// type Agent struct {
//   gorm.Model
//   ID    				int
//   Name  				string
// 	Slug  				string
// 	LLMID   			int
// 	LLM   				*LLM
// 	DefaultPrimer string
// 	Primer 				string
// 	IsActive 			bool
//   Tools 				[]Tool `gorm:"many2many:agent_tools;"`
// }

// type Tool struct {
//   gorm.Model
//   ID      int
//   Name    string
// 	Slug  	string
// 	Token   string
// 	Avatars []Avatar `gorm:"many2many:avatar_tools;"`
//   Agents  []Agent  `gorm:"many2many:agent_tools;"`
// }

// type LLM struct {
// 	gorm.Model
// 	ID   		int
// 	Name 		string
// 	Slug  	string
// 	Avatars []Avatar
// 	Agents 	[]Agent
// }

// type Plugins struct {
// 	gorm.Model
// 	ID   	 int
// 	Name 	 string
// 	Slug   string
// 	Avatars []Avatar `gorm:"many2many:avatar_plugins;"`
// }

// type Role struct {
// 	gorm.Model
// 	ID 		int
// 	Name 	string
// 	Slug  string
// 	Users []User
// }
