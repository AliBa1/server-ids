package document

import (
	"database/sql"
	"errors"
	"server-ids/internal/models"
)

type DocsRepository struct {
	db *sql.DB
}

func NewDocRepository(db *sql.DB) *DocsRepository {
	return &DocsRepository{
		db: db,
	}
}

// GetUser from session key
// AddSession

func (r *DocsRepository) GetDocs() ([]models.Document, error) {
	query := `SELECT title, content, locked FROM documents;`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var docs []models.Document
	for rows.Next() {
		var doc models.Document
		err := rows.Scan(&doc.Title, &doc.Content, &doc.IsLocked)
		if err != nil {
			return docs, err
		}
		docs = append(docs, doc)
	}

	err = rows.Err()
	if err != nil {
		return docs, err
	}
	return docs, err
}

func (r *DocsRepository) GetDoc(title string) (*models.Document, error) {
	query := `SELECT title, content, locked FROM documents WHERE title = ?;`
	row := r.db.QueryRow(query, title)

	var doc models.Document
	err := row.Scan(&doc.Title, &doc.Content, &doc.IsLocked)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("document '" + title + "' not found")
		}
		return nil, err
	}

	return &doc, nil
}
