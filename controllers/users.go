package controllers

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/AstroSynapseAI/app-service/models"
	"github.com/AstroSynapseAI/app-service/repositories"
	"github.com/AstroSynapseAI/app-service/sdk/crud/database"
	"github.com/AstroSynapseAI/app-service/sdk/rest"
	mail "github.com/xhit/go-simple-mail/v2"

	"github.com/thanhpk/randstr"
	"gopkg.in/yaml.v2"

	"github.com/AstroSynapseAI/app-service/sdk/crud/orms/gorm"
)

type UsersController struct {
	rest.Controller
	User    *repositories.UsersRepository
	Account *gorm.Repository[models.Account]
}

func NewUsersController(db *database.Database) *UsersController {
	return &UsersController{
		User:    repositories.NewUsersRepository(db),
		Account: gorm.NewRepository[models.Account](db, models.Account{}),
	}
}

func (ctrl *UsersController) Run() {
	ctrl.Post("/login", ctrl.Login)
	ctrl.Post("/register", ctrl.Register)
	ctrl.Post("/register/invite", ctrl.RegisterInvite)
	ctrl.Post("/confirm_email", ctrl.ConfirmEmail)
	ctrl.Post("/invite", ctrl.CreateInvite)
	ctrl.Post("/password_recovery", ctrl.CreatePasswordRecovery)
	ctrl.Post("/{id}/accounts/save", ctrl.SaveAccount)
	ctrl.Post("/{id}/save/profile", ctrl.SaveProfile)

	ctrl.Put("/{id}/change/password", ctrl.ChangePassword)
	ctrl.Put("/{id}/change/email", ctrl.ChangeEmail)

	ctrl.Get("/{id}/accounts", ctrl.GetAccounts)
	ctrl.Get("/{id}/accounts/{account_id}", ctrl.GetAccount)
	ctrl.Get("/{id}/avatars", ctrl.GetAvatar)
	ctrl.Get("/invited/{token}", ctrl.GetInvitedUser)
	ctrl.Get("/token", ctrl.GetToken)
	ctrl.Get("/password_recovery/{token}", ctrl.ValidatePasswordRecoveryToken)

}

// Default CRUD routes
func (ctrl *UsersController) ReadAll(ctx *rest.Context) {
	fmt.Println("UsersController.ReadAll")
	users, err := ctrl.User.GetAll()
	if err != nil {
		ctx.SetStatus(http.StatusInternalServerError)
		return
	}
	ctx.JsonResponse(http.StatusOK, users)
}

func (ctrl *UsersController) Read(ctx *rest.Context) {
	fmt.Println("Fetching user")
	userID := ctx.GetID("id")
	user, err := ctrl.User.Repo.Read(userID)
	if err != nil {
		ctx.SetStatus(http.StatusInternalServerError)
		return
	}
	ctx.JsonResponse(http.StatusOK, user)
}

func (ctrl *UsersController) GetToken(ctx *rest.Context) {
	fmt.Println("UsersController.GetToken")
	sessionToken := randstr.String(16)

	var reponseJson struct {
		Token string `json:"token"`
	}

	reponseJson.Token = sessionToken

	ctx.JsonResponse(http.StatusOK, reponseJson)
}

func (ctrl *UsersController) GetInvitedUser(ctx *rest.Context) {
	fmt.Println("Fetching user")
	inviteToken := ctx.GetParam("token")
	if inviteToken == "" {
		ctx.SetStatus(http.StatusBadRequest)
		return
	}

	user, err := ctrl.User.GetByInviteToken(inviteToken)
	if err != nil {
		ctx.SetStatus(http.StatusInternalServerError)
		return
	}

	ctx.JsonResponse(http.StatusOK, user)
}

func (ctrl *UsersController) ValidatePasswordRecoveryToken(ctx *rest.Context) {
	fmt.Println("UsersController.ValidatePasswordRecoveryToken")

	recoveryToken := ctx.GetParam("token")
	if recoveryToken == "" {
		ctx.JsonResponse(http.StatusBadRequest, struct{ Error string }{Error: "Invalid request body"})
		return
	}

	user, err := ctrl.User.GetByResetToken(recoveryToken)
	if err != nil {
		ctx.JsonResponse(http.StatusInternalServerError, struct{ Error string }{Error: err.Error()})
		return
	}

	if time.Since(user.PasswordResetTokenExpiry) >= 24*time.Hour {
		ctx.JsonResponse(http.StatusBadRequest, struct{ Error string }{Error: "Token expired"})
		return
	}

	ctx.JsonResponse(http.StatusOK, user)
}

// Custom routes

// password recovery
func (ctrl *UsersController) CreatePasswordRecovery(ctx *rest.Context) {
	fmt.Println("UsersController.CreatePasswordRecovery")
	var input struct {
		Email string `json:"email"`
	}

	err := ctx.JsonDecode(&input)
	if err != nil {
		ctx.JsonResponse(http.StatusBadRequest, struct{ Error string }{Error: "Invalid request body"})
		return
	}

	record, err := ctrl.User.CreateAndSendRecoveryToken(input.Email)
	if err != nil {
		ctx.JsonResponse(http.StatusBadRequest, struct{ Error string }{Error: err.Error()})
		return
	}

	var envSetup string

	if os.Getenv("ENVIRONMENT") == "LOCAL DEV" {
		envSetup = "http://localhost:5173/password_reset/"
	}

	if os.Getenv("ENVIRONMENT") == "HEROKU DEV" {
		envSetup = "https://dev.asai.astrosynapse.ai/password_reset/"
	}

	if os.Getenv("ENVIRONMENT") == "AWS DEV" {
		envSetup = "https://asai.astrosynapse.ai/password_reset/"
	}

	emailMessage := `
		<p>Please click on the link below to reset your password:</p>
		<a href='` + envSetup + `'>Reset your password</a>
		`
	confirmationEmail := mail.NewMSG()
	confirmationEmail.SetFrom("dispatch@astrosynapse.com")
	confirmationEmail.AddTo(input.Email)
	confirmationEmail.SetSubject("Confirm your AstroSynapse email")
	confirmationEmail.SetBody(mail.TextPlain, emailMessage)

	if confirmationEmail.Error != nil {
		fmt.Println(confirmationEmail.Error)
		ctx.JsonResponse(http.StatusInternalServerError, struct{ Error string }{Error: "Internal error"})
	}

	host, err := getSMTPHost()
	if err != nil {
		ctx.JsonResponse(http.StatusInternalServerError, struct{ Error string }{Error: "Internal error"})
	}

	password, err := getSMTPPassword()
	if err != nil {
		ctx.JsonResponse(http.StatusInternalServerError, struct{ Error string }{Error: "Internal error"})
	}

	username, err := getSMTPUsername()
	if err != nil {
		ctx.JsonResponse(http.StatusInternalServerError, struct{ Error string }{Error: "Internal error"})
	}

	smtp := mail.NewSMTPClient()
	smtp.Host = host
	// smtp.KeepAlive = false
	// smtp.ConnectTimeout = 10 * time.Second
	// smtp.SendTimeout = 10 * time.Second
	smtp.Password = password
	smtp.Username = username
	smtp.Encryption = mail.EncryptionSSLTLS
	smtp.Port = 465

	smtpClient, err := smtp.Connect()
	if err != nil {
		ctx.JsonResponse(http.StatusInternalServerError, struct{ Error string }{Error: "Internal error"})
	}

	err = confirmationEmail.Send(smtpClient)
	if err != nil {
		ctx.JsonResponse(http.StatusInternalServerError, struct{ Error string }{Error: "Internal error"})
	}

	ctx.JsonResponse(http.StatusOK, record)
}

func getSMTPHost() (string, error) {

	var Config struct {
		SMTPHost string `yaml:"smtp_host"`
	}

	keys, err := os.ReadFile("./app/keys.yaml")
	if err != nil {
		fmt.Println("Error reading keys.yaml:", err)
	}

	err = yaml.Unmarshal(keys, &Config)
	if err != nil {
		fmt.Println("Error unmarshalling keys.yaml:", err)
	}

	return Config.SMTPHost, nil
}

func getSMTPPassword() (string, error) {

	var Config struct {
		SMTPPassword string `yaml:"smtp_password"`
	}

	keys, err := os.ReadFile("./app/keys.yaml")
	if err != nil {
		fmt.Println("Error reading keys.yaml:", err)
	}

	err = yaml.Unmarshal(keys, &Config)
	if err != nil {
		fmt.Println("Error unmarshalling keys.yaml:", err)
	}

	return Config.SMTPPassword, nil
}

func getSMTPUsername() (string, error) {

	var Config struct {
		SMTPUsername string `yaml:"smtp_username"`
	}

	keys, err := os.ReadFile("./app/keys.yaml")
	if err != nil {
		fmt.Println("Error reading keys.yaml:", err)
	}

	err = yaml.Unmarshal(keys, &Config)
	if err != nil {
		fmt.Println("Error unmarshalling keys.yaml:", err)
	}

	return Config.SMTPUsername, nil
}

// create user invite
func (ctrl *UsersController) CreateInvite(ctx *rest.Context) {
	fmt.Println("UsersController.CreateInvite")
	var input struct {
		Username string `json:"username"`
	}

	err := ctx.JsonDecode(&input)
	if err != nil {
		ctx.SetStatus(http.StatusBadRequest)
		return
	}

	record, err := ctrl.User.CreateInvite(input.Username)
	if err != nil {
		ctx.SetStatus(http.StatusInternalServerError)
		return
	}

	ctx.JsonResponse(http.StatusOK, record)
}

func (ctrl *UsersController) ConfirmEmail(ctx *rest.Context) {
	fmt.Println("UsersController.ConfirmEmail")

	var reqData struct {
		Email string `json:"email"`
		Token string `json:"token"`
	}

	err := ctx.JsonDecode(&reqData)
	if err != nil {
		ctx.JsonResponse(http.StatusBadRequest, struct{ Error string }{Error: "Username is required"})
		return
	}

	account, err := ctrl.User.GetAccountByEmail(reqData.Token)
	if err != nil {
		ctx.SetStatus(http.StatusInternalServerError)
		return
	}

	account.Email = reqData.Email

	updatedAccount, err := ctrl.User.SaveAccount(account)
	if err != nil {
		ctx.JsonResponse(http.StatusInternalServerError, struct{ Error string }{Error: "Internal error"})
		return
	}

	ctx.JsonResponse(http.StatusOK, updatedAccount)

}

// register invited user
func (ctrl *UsersController) RegisterInvite(ctx *rest.Context) {
	fmt.Println("UsersController.RegisterInvite")
	var reqData struct {
		Username    string `json:"username"`
		Password    string `json:"password"`
		InviteToken string `json:"invite_token"`
	}

	err := ctx.JsonDecode(&reqData)
	if err != nil {
		ctx.JsonResponse(http.StatusBadRequest, struct{ Error string }{Error: "Username is required"})
		return
	}
	if reqData.Username == "" {
		ctx.JsonResponse(http.StatusBadRequest, struct{ Error string }{Error: "Username is required"})
		return
	}
	if reqData.Password == "" {
		ctx.JsonResponse(http.StatusBadRequest, struct{ Error string }{Error: "Password is required"})
		return
	}
	if reqData.InviteToken == "" {
		ctx.JsonResponse(http.StatusBadRequest, struct{ Error string }{Error: "Invite token is required"})
		return
	}
	if len(reqData.Password) < 8 {
		ctx.JsonResponse(http.StatusBadRequest, struct{ Error string }{Error: "Password must be at least 8 characters long"})
		return
	}

	user, err := ctrl.User.ConfirmInvite(
		reqData.Username,
		reqData.Password,
		reqData.InviteToken,
	)
	if err != nil {
		ctx.JsonResponse(http.StatusBadRequest, struct{ Error string }{Error: err.Error()})
		return
	}

	ctx.JsonResponse(http.StatusOK, user)
}

func (ctrl *UsersController) Register(ctx *rest.Context) {
	fmt.Println("UsersController.Register")
	var reqData struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	err := ctx.JsonDecode(&reqData)
	if err != nil {
		ctx.SetStatus(http.StatusBadRequest)
		return
	}

	user, err := ctrl.User.Register(reqData.Username, reqData.Password)
	if err != nil {
		ctx.SetStatus(http.StatusUnauthorized)
		return
	}

	ctx.JsonResponse(http.StatusOK, user)
}

func (ctrl *UsersController) Login(ctx *rest.Context) {
	fmt.Println("UsersController.Login")
	var reqData struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	err := ctx.JsonDecode(&reqData)
	if err != nil {
		ctx.JsonResponse(http.StatusBadRequest, struct{ Error string }{Error: err.Error()})
		return
	}

	if reqData.Username == "" {
		ctx.JsonResponse(http.StatusBadRequest, struct{ Error string }{Error: "Username is required"})
		return
	}

	if reqData.Password == "" {
		ctx.JsonResponse(http.StatusBadRequest, struct{ Error string }{Error: "Password is required"})
		return
	}

	user, err := ctrl.User.Login(reqData.Username, reqData.Password)
	if err != nil {
		ctx.JsonResponse(http.StatusUnauthorized, struct{ Error string }{Error: err.Error()})
		return
	}

	ctx.JsonResponse(http.StatusOK, user)
}

func (ctrl *UsersController) GetAccounts(ctx *rest.Context) {
	fmt.Println("UsersController.GetAccounts")

	userID := ctx.GetID()

	accounts, err := ctrl.User.GetAccounts(userID)
	if err != nil {
		ctx.SetStatus(http.StatusInternalServerError)
		return
	}

	ctx.JsonResponse(http.StatusOK, accounts)
}

func (ctrl *UsersController) GetAccountByEmail(ctx *rest.Context) {
	fmt.Println("UsersController.GetAccountByEmail")
	var reqData struct {
		Email string `json:"email"`
	}

	err := ctx.JsonDecode(&reqData)
	if err != nil {
		ctx.SetStatus(http.StatusBadRequest)
		return
	}

	account, err := ctrl.User.GetAccountByEmail(reqData.Email)
	if err != nil {
		ctx.SetStatus(http.StatusInternalServerError)
		return
	}

	ctx.JsonResponse(http.StatusOK, account)
}

func (ctrl *UsersController) GetAccount(ctx *rest.Context) {
	accountID := ctx.GetID("account_id")
	account, err := ctrl.User.GetAccount(ctx.GetID(), accountID)
	if err != nil {
		ctx.SetStatus(http.StatusInternalServerError)
		return
	}

	ctx.JsonResponse(http.StatusOK, account)
}

func (ctrl *UsersController) GetAvatar(ctx *rest.Context) {
	fmt.Println("UsersController.GetAvatar")
	avatar, err := ctrl.User.GetUserAvatar(ctx.GetID())
	if err != nil {
		ctx.SetStatus(http.StatusNotFound)
		return
	}

	ctx.JsonResponse(http.StatusOK, avatar)
}

func (ctrl *UsersController) Update(ctx *rest.Context) {
	fmt.Println("UsersController.Update")
	userID := ctx.GetID()

	var record models.User
	record.ID = userID

	err := ctx.JsonDecode(&record)
	if err != nil {
		ctx.SetStatus(http.StatusBadRequest)
		return
	}

	user, err := ctrl.User.Update(record)
	if err != nil {
		ctx.SetStatus(http.StatusInternalServerError)
		return
	}

	ctx.JsonResponse(http.StatusOK, user)
}

func (ctrl *UsersController) SaveAccount(ctx *rest.Context) {
	fmt.Println("UsersController.SaveAccount")
	userID := ctx.GetID()

	var record models.Account
	err := ctx.JsonDecode(&record)
	if err != nil {
		ctx.SetStatus(http.StatusBadRequest)
		return
	}

	record.UserID = userID
	account, err := ctrl.User.SaveAccount(record)
	if err != nil {
		ctx.SetStatus(http.StatusInternalServerError)
		return
	}

	ctx.JsonResponse(http.StatusOK, account)
}

func (ctrl *UsersController) SaveProfile(ctx *rest.Context) {
	fmt.Println("UsersController.SaveProfile")
	userID := ctx.GetID()

	var reqData struct {
		AccountID uint   `json:"account_id,omitempty"`
		Username  string `json:"username"`
		FirstName string `json:"first_name,omitempty"`
		LastName  string `json:"last_name,omitempty"`
	}

	err := ctx.JsonDecode(&reqData)
	if err != nil {
		ctx.JsonResponse(http.StatusBadRequest, struct{ Error string }{Error: "Invalid request body"})
		return
	}

	if reqData.Username == "" {
		ctx.JsonResponse(http.StatusBadRequest, struct{ Error string }{Error: "Username is required"})
		return
	}

	user, err := ctrl.User.FetchUser(userID)
	if err != nil {
		ctx.JsonResponse(http.StatusInternalServerError, struct{ Error string }{Error: "Failed to fetch user"})
		return
	}

	var account models.Account
	if len(user.Accounts) == 0 {
		account = models.Account{
			UserID:    userID,
			FirstName: reqData.FirstName,
			LastName:  reqData.LastName,
			Username:  reqData.Username,
		}
	} else {
		account = user.Accounts[0]
		account.ID = reqData.AccountID
		account.FirstName = reqData.FirstName
		account.LastName = reqData.LastName
		account.Username = reqData.Username
	}

	if account.UserID == 0 || account.FirstName == "" || account.LastName == "" || account.Username == "" {
		ctx.JsonResponse(http.StatusBadRequest, struct{ Error string }{Error: "Account data is invalid"})
		return
	}

	_, err = ctrl.User.SaveAccount(account)
	if err != nil {
		ctx.JsonResponse(http.StatusInternalServerError, struct{ Error string }{Error: "Failed to save account"})
		return
	}

	user, err = ctrl.User.UpdateUsername(userID, reqData.Username)
	if err != nil {
		ctx.JsonResponse(http.StatusInternalServerError, struct{ Error string }{Error: "Failed to update username"})
		return
	}

	ctx.JsonResponse(http.StatusOK, user)

}

func (ctrl *UsersController) ChangePassword(ctx *rest.Context) {
	fmt.Println("UsersController.ChangePassword")
	userID := ctx.GetID()

	var reqData struct {
		Password string `json:"password"`
	}

	err := ctx.JsonDecode(&reqData)
	if err != nil {
		ctx.JsonResponse(http.StatusBadRequest, struct{ Error string }{Error: "Invalid request body"})
		return
	}

	if len(reqData.Password) < 8 || reqData.Password == "" {
		ctx.JsonResponse(http.StatusBadRequest, struct{ Error string }{Error: "Password is not long enough"})
		return
	}

	user, err := ctrl.User.UpdatePassword(userID, reqData.Password)
	if err != nil {
		ctx.JsonResponse(http.StatusInternalServerError, struct{ Error string }{Error: "Failed to update password"})
		return
	}

	usr, err := ctrl.User.RemovePasswordResetToken(user.ID)
	if err != nil {
		ctx.JsonResponse(http.StatusInternalServerError, struct{ Error string }{Error: "Internal error"})
		return
	}

	ctx.JsonResponse(http.StatusOK, usr)
}

func (ctrl *UsersController) ChangeEmail(ctx *rest.Context) {
	fmt.Println("UsersController.ChangeEmail")

	var reqData struct {
		Email     string `json:"email,omitempty"`
		AccountID uint   `json:"account_id,omitempty"`
	}

	err := ctx.JsonDecode(&reqData)
	if err != nil {
		ctx.JsonResponse(http.StatusBadRequest, struct{ Error string }{Error: "Invalid request body"})
		return
	}

	user, err := ctrl.User.FetchUser(ctx.GetID())
	if err != nil {
		ctx.JsonResponse(http.StatusInternalServerError, struct{ Error string }{Error: "Failed to fetch user"})
		return
	}

	account := user.Accounts[0]
	account.Email = reqData.Email

	_, err = ctrl.User.SaveAccount(account)
	if err != nil {
		ctx.JsonResponse(http.StatusInternalServerError, struct{ Error string }{Error: "Failed to save account"})
		return
	}

	ctx.JsonResponse(http.StatusOK, nil)

	// var reqData struct {
	// 	Email string `json:"email"`
	// }
	//
	// account, err := ctrl.User.GetAccountByUserID(userID)
	// if err != nil {
	// 	ctx.JsonResponse(http.StatusBadRequest, struct{ Error string }{Error: "Account not found"})
	// 	return
	// }
	//
	// decodeErr := ctx.JsonDecode(&reqData)
	// if decodeErr != nil {
	// 	ctx.JsonResponse(http.StatusBadRequest, struct{ Error string }{Error: "Invalid request body"})
	// 	return
	// }
	//
	// token, err := ctrl.User.CreateAndSendEmailConfirmation(account.ID, reqData.Email)
	// if err != nil {
	// 	ctx.JsonResponse(http.StatusBadRequest, struct{ Error string }{Error: err.Error()})
	// 	return
	// }
	//
	// confirmEmailURL := ""
	//
	// if os.Getenv("ENVIRONMENT") == "LOCAL DEV" {
	// 	confirmEmailURL = "http://localhost:5173/email_confirmation/?token=" + token + "&email=" + reqData.Email
	// }
	//
	// if os.Getenv("ENVIRONMENT") == "HEROKU DEV" {
	// 	confirmEmailURL = "https://dev.asai.astrosynapse.ai/email_confirmation/?token=" + token + "&email=" + reqData.Email
	// }
	//
	// if os.Getenv("ENVIRONMENT") == "AWS DEV" {
	// 	confirmEmailURL = "https://asai.astrosynapse.ai/email_confirmation/?token=" + token + "&email=" + reqData.Email
	// }
	//
	// emailMessage := `
	// 	<p>Please click on the link below to confirm your email address:</p>
	// 	<a href='` + confirmEmailURL + `'>Confirm Email</a>
	// 	`
	// confirmationEmail := mail.NewMSG()
	// confirmationEmail.SetFrom("dispatch@astrosynapse.com")
	// confirmationEmail.AddTo(reqData.Email)
	// confirmationEmail.SetSubject("Confirm your email")
	// confirmationEmail.SetBody(mail.TextPlain, emailMessage)
	//
	// if confirmationEmail.Error != nil {
	// 	fmt.Println(confirmationEmail.Error)
	// 	ctx.JsonResponse(http.StatusInternalServerError, struct{ Error string }{Error: "Internal error"})
	// }
	//
	// host, err := getSMTPHost()
	// if err != nil {
	// 	ctx.JsonResponse(http.StatusInternalServerError, struct{ Error string }{Error: "Internal error"})
	// }
	//
	// password, err := getSMTPPassword()
	// if err != nil {
	// 	ctx.JsonResponse(http.StatusInternalServerError, struct{ Error string }{Error: "Internal error"})
	// }
	//
	// username, err := getSMTPUsername()
	// if err != nil {
	// 	ctx.JsonResponse(http.StatusInternalServerError, struct{ Error string }{Error: "Internal error"})
	// }
	//
	// smtp := mail.NewSMTPClient()
	// smtp.Host = host
	// // smtp.KeepAlive = false
	// // smtp.ConnectTimeout = 10 * time.Second
	// // smtp.SendTimeout = 10 * time.Second
	// smtp.Password = password
	// smtp.Username = username
	// smtp.Encryption = mail.EncryptionSSLTLS
	// smtp.Port = 465
	//
	// smtpClient, err := smtp.Connect()
	// if err != nil {
	// 	ctx.JsonResponse(http.StatusInternalServerError, struct{ Error string }{Error: "Internal error"})
	// }
	//
	// err = confirmationEmail.Send(smtpClient)
	// if err != nil {
	// 	ctx.JsonResponse(http.StatusInternalServerError, struct{ Error string }{Error: "Internal error"})
	// }
	//
	// ctx.JsonResponse(http.StatusOK, models.User{})
}
