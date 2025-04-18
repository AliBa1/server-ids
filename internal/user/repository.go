package user

import (
	"database/sql"
	"errors"
	"server-ids/internal/models"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) GetUsers() ([]models.User, error) {
	// query := `SELECT username, role, last_login_date FROM users;`
	query := `SELECT username, role FROM users;`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User
		// err := rows.Scan(&user.Username, &user.Role, &user.LastLoginDate)
		err := rows.Scan(&user.Username, &user.Role)
		if err != nil {
			return users, err
		}
		users = append(users, user)
	}

	err = rows.Err()
	if err != nil {
		return users, err
	}
	return users, err
}

func (r *UserRepository) GetUser(username string) (*models.User, error) {
	// query := `SELECT username, role, last_login_date FROM users WHERE username = ?;`
	query := `SELECT username, password, role FROM users WHERE username = ?;`
	row := r.db.QueryRow(query, username)

	var user models.User
	// err := row.Scan(&user.Username, &user.Role, &user.LastLoginDate)
	err := row.Scan(&user.Username, &user.Password, &user.Role)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("user '" + username + "' does not exist")
		}
		return nil, err
	}

	return &user, nil
}

func (r *UserRepository) CreateUser(user models.User) error {
	// query := `
	// 	INSERT INTO users (username, password, role, last_login_date)
	// 	VALUES (?, ?, ?, ?);
	// `
	query := `
		INSERT INTO users (username, password, role)
		VALUES (?, ?, ?);
	`

	// _, err := r.db.Exec(query, user.Username, user.Password, user.Role, user.LastLoginDate)
	_, err := r.db.Exec(query, user.Username, user.Password, user.Role)
	return err
}

func (r *UserRepository) UpdateUser(user models.User) error {
	query := `
		UPDATE users
		SET password = ?, role = ?, last_login_date = ?
		WHERE username = ?;
	`

	_, err := r.db.Exec(query, user.Password, user.Role, user.LastLoginDate, user.Username)
	return err
}
