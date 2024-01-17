package repositories

import (
	"github.com/AstroSynapseAI/app-service/models"
	"github.com/AstroSynapseAI/app-service/sdk/crud/database"
	"github.com/AstroSynapseAI/app-service/sdk/crud/orms/gorm"
	db "gorm.io/gorm"
)

type LLMSRepository struct {
	Repo   *gorm.Repository[models.LLM]
	Active *gorm.Repository[models.ActiveLLM]
}

func NewLLMSRepository(db *database.Database) *LLMSRepository {
	return &LLMSRepository{
		Repo:   gorm.NewRepository[models.LLM](db, models.LLM{}),
		Active: gorm.NewRepository[models.ActiveLLM](db, models.ActiveLLM{}),
	}
}

func (llm *LLMSRepository) SaveActiveLLM(avatarData models.ActiveLLM) error {
	result := llm.Active.DB.Save(&avatarData)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (llm *LLMSRepository) ToggleActiveLLM(avatarID uint, LLMID uint, active bool) error {
	var activeLLM models.ActiveLLM

	result := llm.Active.DB.Where("avatar_id = ? AND LLM_id = ?", avatarID, LLMID).First(&activeLLM)
	if result.Error == db.ErrRecordNotFound {
		activeLLM.LLMID = LLMID
		activeLLM.AvatarID = avatarID
	}

	activeLLM.IsActive = active

	result = llm.Active.DB.Save(&activeLLM)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (llm *LLMSRepository) GetActiveLLMs(avatarID uint) []models.ActiveLLM {
	query := llm.Active.DB
	query = query.Preload("LLM")
	query = query.Where("avatar_id = ?", avatarID)

	var activeLLMs []models.ActiveLLM

	result := query.Find(&activeLLMs)
	if result.Error != nil {
		return []models.ActiveLLM{}
	}

	return activeLLMs
}

func (llm *LLMSRepository) GetActiveLLM(avatarID uint, LLMID uint) (models.ActiveLLM, error) {
	query := llm.Active.DB
	query = query.Preload("LLM")
	query = query.Where("avatar_id = ? AND LLM_id = ?", avatarID, LLMID)

	var activeLLM models.ActiveLLM

	result := query.First(&activeLLM)
	if result.Error != nil {
		return models.ActiveLLM{}, result.Error
	}

	return activeLLM, nil
}
