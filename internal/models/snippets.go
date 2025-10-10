package models

import (
	"database/sql"
	"errors"
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
	// Build the query
	stmt := `SELECT id, title, content, created, expires FROM snippets WHERE expires > UTC_TIMESTAMP() AND id = ?`

	// Execute the query
	row := m.DB.QueryRow(stmt, id)

	// Initialize a pointer to a new zeroed Snippet struct
	s := &Snippet{}

	// Scan the query results into the struct
	err := row.Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNoRecord
		} else {
			return nil, err
		}
	}

	return s, nil
}

// GetRecent This will return the 10 most recently created snippets.
func (m *SnippetModel) GetRecent(limit int) ([]*Snippet, error) {
	return nil, nil
}

func (m *SnippetModel) Latest() ([]*Snippet, error) {
	// SQL statement
	stmt := `SELECT id, title, content, created, expires FROM snippets
    WHERE expires > UTC_TIMESTAMP() ORDER BY id DESC LIMIT 10`

	// Execute the query
	rows, err := m.DB.Query(stmt)
	if err != nil {
		return nil, err
	}
	// Close the row object before the GetRecent() method returns.
	defer rows.Close()

	// Initialize an empty slice to hold the snippets
	snippets := []*Snippet{}

	// Loop through the rows and populate the slice
	for rows.Next() {
		s := &Snippet{}

		err := rows.Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)
		if err != nil {
			return nil, err
		}
		snippets = append(snippets, s)
	}

	// Check for any errors
	if err = rows.Err(); err != nil {
		return nil, err
	}

	// Return the slice of snippets
	return snippets, nil
}
