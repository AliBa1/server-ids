package mock

import "server-ids/internal/models"

func GetMockUsers() []models.User {
	return []models.User{
		*models.NewUser("funguy123", "admin12345", "admin"),
		*models.NewUser("bossman", "emp12345", "employee"),
		*models.NewUser("grumpy", "guest12345", "guest"),
		*models.NewUser("jpearson", "guest12345", "guest"),
		*models.NewUser("fredrick5", "guest12345", "guest"),
		*models.NewUser("ballhoggary", "emp12345", "employee"),
		*models.NewUser("erick", "admin12345", "admin"),
		*models.NewUser("barrylarry", "emp12345", "employee"),
		*models.NewUser("twotthree", "guest12345", "guest"),
		*models.NewUser("yap", "guest12345", "guest"),
		*models.NewUser("boardman45", "guest12345", "guest"),
		*models.NewUser("1819twenty", "emp12345", "employee"),
		*models.NewUser("opi", "guest12345", "guest"),
		*models.NewUser("patrick", "guest12345", "guest"),
		*models.NewUser("fred111", "guest12345", "guest"),
		*models.NewUser("secure21", "guest12345", "guest"),
	}
}