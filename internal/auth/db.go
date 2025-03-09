package auth

// CRUD database

type AuthDB interface {
	GetAllLogins() ([]string, error)
}

type AuthDBMemory struct {
	logins []string
}

func NewAuthDBMemory() *AuthDBMemory {
	return &AuthDBMemory{logins: []string{"test1", "test2", "test3"}}
}

func (db *AuthDBMemory) GetAllLogins() ([]string, error) {
	return db.logins, nil
}