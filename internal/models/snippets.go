package models

import (
	"database/sql"
	"errors"
	"time"
)

type Snippet struct {
	ID      int
	Title   string
	Content string
	Created time.Time
	Expires time.Time
}

// SQL DB Pool wrapper
type SnippetModel struct {
	DB *sql.DB
}

// Insert new Snippet in database
func (m *SnippetModel) Insert(title string, content string, expires int) (int, error) {
	statement := `INSERT INTO snippets(title, content, created, expires)
	VALUES(?, ?, UTC_TIMESTAMP(), DATE_ADD(UTC_TIMESTAMP(), INTERVAL ? DAY))`

	result, err := m.DB.Exec(statement, title, content, expires)
	if err != nil {
		return 0, err
	}

	// Get ID of our latest inserted record above
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), err
}

// Get Snippet by ID
func (m *SnippetModel) Get(id int) (*Snippet, error) {
	statement := `SELECT * FROM snippets
	WHERE expires > UTC_TIMESTAMP() AND id = ?`

	row := m.DB.QueryRow(statement, id)

	s := &Snippet{}

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

// Get latest 10 snippets
func (m *SnippetModel) Latest() ([]*Snippet, error) {
	statement := `SELECT id, title, content, created, expires FROM snippets
	WHERE expires > UTC_TIMESTAMP() ORDER BY id DESC LIMIT 10`

	rows, err := m.DB.Query(statement)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	// Empty slice of Snippets to hold the results
	snippets := []*Snippet{}

	for rows.Next() {
		// Pointer to a new zeroed Snippet struct
		s := &Snippet{}

		// Copy the values from result to zeroed struct. Number of arguments should be equal to number of columns returned.
		err := rows.Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)
		if err != nil {
			return nil, err
		}

		snippets = append(snippets, s)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return snippets, nil
}
