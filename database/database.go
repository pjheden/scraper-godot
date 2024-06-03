package database

import (
	"database/sql"
	"fmt"

	"github.com/pjheden/scraper-godot/assets"

	sq "github.com/Masterminds/squirrel"
	// postgres driver.
	_ "github.com/lib/pq"
)

type Database struct {
	connStr string
}

func New(username, password, ip string) *Database {
	return &Database{
		connStr: fmt.Sprintf("postgresql://%s:%s@%s/todos?sslmode=disable", username, password, ip),
	}
}

func (d *Database) Open() (*sql.DB, error) {
	return sql.Open("postgres", d.connStr)
}

/*
Assets retrieves all assets defined in table assets and returns them in a slice of *Assets
*/
func (d *Database) Assets() ([]*assets.Asset, error) {
	query := sq.Select(
		"id",
		"title",
		"description",
		"creator",
		"version",
		"repository_url",
		"stars",
		"first_commit",
		"latest_commit",
	).
		From("assets")

	conn, err := d.Open()
	if err != nil {
		return nil, fmt.Errorf("connecting to db: %v", err)
	}

	defer conn.Close()

	rows, err := query.PlaceholderFormat(sq.Dollar).RunWith(conn).Query()
	if err != nil {
		return nil, fmt.Errorf("getting rows: %v", err)
	}

	allAssets := []*assets.Asset{}
	for rows.Next() {
		t := &assets.Asset{}
		err := rows.Scan(
			&(t.ID),
			&(t.Title),
			&(t.Description),
			&(t.Creator),
			&(t.Version),
			&(t.RepositoryURL),
			&(t.Stars),
			&(t.FirstCommit),
			&(t.LatestCommit),
		)
		if err != nil {
			return nil, fmt.Errorf("scanning issues: %v", err)
		}
		allAssets = append(allAssets, t)
	}

	return allAssets, nil
}
