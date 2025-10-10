package models

import (
	"database/sql"
	"time"
)

// Snippet Defines a Snippet type to hold data for individual snippets.
type Snippet struct {
	ID      int
	Title   string
	Content string
	Created time.Time
	Expires time.Time
}

// SnippetModel Define a SnippetModel type which wraps a sql.DB connection pool.
type SnippetModel struct {
	DB *sql.DB
}

// Insert This will insert a new snippet into the database.
func (m *SnippetModel) Insert(title string, content string, expires int) (int, error) {
	stmt := `INSERT INTO snippets (title, content, created, expires ) VALUES (?, ?, UTC_TIMESTAMP(), DATE_ADD(UTC_TIMESTAMP(), INTERVAL ? DAY))`

	result, err := m.DB.Exec(stmt, title, content, expires)
	if err != nil {
		return 0, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

// Get This will return a specific snippet based on its ID.
func (m *SnippetModel) Get(id int) (*Snippet, error) {
	return nil, nil
}

// GetRecent This will return the 10 most recently created snippets.
func (m *SnippetModel) GetRecent(limit int) ([]*Snippet, error) {
	return nil, nil
}
