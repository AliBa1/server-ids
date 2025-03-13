package mock

import (
	"server-ids/internal/models"
)

func GetMockDocuments() []models.Document {
	return []models.Document{
		*models.NewDocument("Top Secret Case Study #1", "admin12345", true),
	}
}
