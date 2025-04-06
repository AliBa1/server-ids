package mock

import (
	"server-ids/internal/models"
)

func GetMockDocuments() []models.Document {
	return []models.Document{
		*models.NewDocument("Onboarding Document", "Welcome to the company.", true),
		*models.NewDocument("First Doc Ever", "Everyone can see this document. Your welcome!", false),
		*models.NewDocument("Top Secret Case Study #1", "All contents of this document should be kept a secret. Only admins should have this document. Do NOT shore with anyone!", true),
	}
}
