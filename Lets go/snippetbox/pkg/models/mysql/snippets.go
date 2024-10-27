package mysql

import (
	"database/sql"
	"errors"
	"snippetbox/pkg/models"
)

type SnippetModel struct {
	DB *sql.DB
}

func (m *SnippetModel) Insert(title, content, expires string) (int, error) {
	query_string := `insert into snippets (title, content, created, expires)
	values(?, ?, utc_timestamp(), date_add(utc_timestamp(), interval ? day))`
	
	result, err := m.DB.Exec(query_string, title, content, expires)
	
	if err != nil {
		return 0, err
	}

	lastId, err := result.LastInsertId()
	
	if err != nil {
		return 0, nil
	}

	return int(lastId), nil
}

func (m *SnippetModel) Get(id int) (*models.Snippet, error) {
	query_string := `select id, title, content, created, expires from snippets
	where expires > utc_timestamp() and id = ?`

	sql_row := m.DB.QueryRow(query_string, id)

	s := &models.Snippet{}

	err := sql_row.Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, models.ErrNoRecord
		} else {
			return nil, err
		}
	}

	return s, nil
}

func (m *SnippetModel) Latest() ([]*models.Snippet, error) {
	query_string := `SELECT id, title, content, created, expires FROM snippets
	WHERE expires > UTC_TIMESTAMP() ORDER BY created DESC, id DESC LIMIT 10`

	res_rows, err := m.DB.Query(query_string)
	if err != nil {
		return nil, err
	}
	defer res_rows.Close()

	res_slice := make([]*models.Snippet, 0, 10)
	for res_rows.Next() {
		s := &models.Snippet{}
		err = res_rows.Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)
		if err != nil {
			return nil, err
		}
		res_slice = append(res_slice, s)
	}

	if err := res_rows.Err(); err != nil {
		return nil, err
	}

	return res_slice, nil
}