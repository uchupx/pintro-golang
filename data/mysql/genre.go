package mysql

import (
	"context"
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/uchupx/pintro-golang/data/model"
)

const findGenresByQuery = `
	SELECT 
		* 
	FROM genre
`

const findGenreByIdsQuery = `
	SELECT * FROM genre where id IN (%s)
`

type genreMySQLRepository struct {
	db *sqlx.DB
}

func (m genreMySQLRepository) FindAll(ctx context.Context) ([]model.Genre, error) {
	var genres []model.Genre

	stmt, err := m.db.PrepareContext(ctx, findGenresByQuery)
	if err != nil {
		return nil, err
	}

	rows, err := stmt.QueryContext(ctx)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var genre model.Genre

		err = rows.Scan(
			&genre.Id,
			&genre.GenreName,
		)

		if err != nil {
			return nil, err
		}

		genres = append(genres, genre)
	}

	return genres, nil
}

func (m genreMySQLRepository) FindByIds(ctx context.Context, ids []uint64) ([]model.Genre, error) {
	var items []model.Genre
	var args []interface{}

	if len(ids) == 0 {
		return nil, nil
	}

	placeholder := strings.TrimRight(strings.Repeat("?,", len(ids)), ",")
	query := fmt.Sprintf(findGenreByIdsQuery, placeholder)

	stmt, err := m.db.PrepareContext(ctx, query)
	if err != nil {
		return nil, err
	}

	for _, id := range ids {
		args = append(args, id)
	}

	rows, err := stmt.QueryContext(ctx, args...)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var item model.Genre

		err = rows.Scan(
			&item.Id,
			&item.GenreName,
		)

		if err != nil {
			return nil, err
		}

		items = append(items, item)
	}

	return items, nil
}

func NewGenreMysqlRepo(db *sqlx.DB) genreMySQLRepository {
	return genreMySQLRepository{
		db: db,
	}
}
