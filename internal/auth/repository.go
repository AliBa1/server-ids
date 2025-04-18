package auth

import (
	"database/sql"
	"server-ids/internal/models"

	"github.com/google/uuid"
)

type AuthRepository struct {
	db *sql.DB
}

func NewAuthRepository(db *sql.DB) *AuthRepository {
	return &AuthRepository{
		db: db,
	}
}

func (r *AuthRepository) AddSession(key uuid.UUID, user models.User) error {
	query := `
		INSERT INTO sessions (key, username)
		VALUES (?, ?);
	`

	_, err := r.db.Exec(query, key.String(), user.Username)
	return err
}

// func (r *AuthRepository) FindSession(key uuid.UUID, user models.User) error {
// 	query := `
// 		SELECT sessions (key, username)
// 		VALUES (?, ?);
// 	`

// 	_, err := r.db.Exec(query, key.String(), user.Username)
// 	return err
// }
