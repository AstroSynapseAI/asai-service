package controllers

// import (
// 	"context"
// 	"database/sql/driver"
// 	"encoding/json"
// 	"errors"
// 	"fmt"

// 	"github.com/AstroSynapseAI/app-service/app"
// 	"github.com/AstroSynapseAI/app-service/engine/chains"
// 	"github.com/GoLangWebSDK/rest"
// 	"github.com/thanhpk/randstr"
// 	"gorm.io/driver/postgres"
// 	"gorm.io/gorm"
// )

// type ChatHistory struct {
// 	ID           int       `gorm:"primary_key"`
// 	SessionID    string    `gorm:"type:varchar(256)"`
// 	BufferString string    `gorm:"type:text"`
// 	ChatHistory  *Messages `json:"chat_history" gorm:"type:jsonb;column:chat_history"`
// }

// type Messages []Message

// type Message struct {
// 	Type    string `json:"type"`
// 	Content string `json:"text"`
// }

// func (m Messages) Value() (driver.Value, error) {
// 	return json.Marshal(m)
// }

// // Scan implements the sql.Scanner interface, this method allows us to
// // define how we convert the Message data from the database into our Message type.
// func (m *Messages) Scan(src interface{}) error {
// 	if bytes, ok := src.([]byte); ok {
// 		return json.Unmarshal(bytes, m)
// 	}
// 	return errors.New("could not scan type into Message")
// }

// func GetHistory(ctx *rest.Context) {
// 	fmt.Println("Fetching history...")
// 	// dsn := config.SetupPostgreDSN()
// 	dsn := app.CONFIG.DSN
// 	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}

// 	var sessionID = ctx.GetParam("session_id")
// 	var history *ChatHistory

// 	err = db.Where(ChatHistory{SessionID: sessionID}).Find(&history).Error
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}

// 	_ = ctx.JsonResponse(200, history.ChatHistory)
// }

// func GetSession(ctx *rest.Context) {
// 	sessionID := randstr.String(16)

// 	var reponseJson struct {
// 		SessionId string `json:"session_id"`
// 	}

// 	reponseJson.SessionId = sessionID
// 	_ = ctx.JsonResponse(200, reponseJson)

// }

// func PostHandler(ctx *rest.Context) {
// 	asaiChain, _ := chains.NewAsaiChain()
// 	// Parse the incoming http request
// 	var request struct {
// 		SessionId  string `json:"session_id"`
// 		UserPrompt string `json:"user_prompt"`
// 	}

// 	err := ctx.JsonDecode(&request)
// 	if err != nil {
// 		fmt.Println("Bad Request: %w", err)
// 		_ = ctx.JsonResponse(400, err)
// 		return
// 	}

// 	asaiChain.SetSessionID(request.SessionId)

// 	response, err := asaiChain.Prompt(context.Background(), request.UserPrompt)
// 	if err != nil {
// 		fmt.Println(err)
// 		_ = ctx.JsonResponse(500, err)
// 		return
// 	}

// 	var responseJson struct {
// 		Content string `json:"content"`
// 	}

// 	responseJson.Content = response

// 	_ = ctx.JsonResponse(200, responseJson)
// }
