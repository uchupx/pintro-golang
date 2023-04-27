package mysql

import (
	"context"
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/uchupx/pintro-golang/data"
	"github.com/uchupx/pintro-golang/data/model"
)

const findGamesByQuery = `
	SELECT 
		* 
	FROM game
	LIMIT ?, ?;
`

const findGamesByIdsQuery = `
	SELECT * FROM game where id IN (%s)
`

type gameMySQLRepository struct {
	db *sqlx.DB
}

func (m gameMySQLRepository) FindByQuery(ctx context.Context, query data.GameQuery) (*data.Collection, error) {
	var games []model.Game
	var args []interface{}

	stmt, err := m.db.PrepareContext(ctx, findGamesByQuery)
	if err != nil {
		return nil, err
	}

	limit, offset := ConvertPagination(query.PerPage, query.Page)

	args = append(args, offset)
	args = append(args, limit)

	rows, err := stmt.QueryContext(ctx, args...)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var game model.Game

		err = rows.Scan(
			&game.Id,
			&game.GenreId,
			&game.GameName,
		)

		if err != nil {
			return nil, err
		}

		games = append(games, game)
	}

	return &data.Collection{
		Data: games,
	}, nil
}

func (m gameMySQLRepository) FindByIds(ctx context.Context, ids []uint64) ([]model.Game, error) {
	var games []model.Game
	var args []interface{}

	if len(ids) == 0 {
		return nil, nil
	}

	placeholder := strings.TrimRight(strings.Repeat("?,", len(ids)), ",")
	query := fmt.Sprintf(findPublisherByIdsQuery, placeholder)

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
		var game model.Game

		err = rows.Scan(
			&game.Id,
			&game.GenreId,
			&game.GameName,
		)

		if err != nil {
			return nil, err
		}

		games = append(games, game)
	}

	return games, nil
}

func NewGameMysqlRepo(db *sqlx.DB) gameMySQLRepository {
	return gameMySQLRepository{
		db: db,
	}
}
