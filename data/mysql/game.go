package mysql

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/uchupx/pintro-golang/data"
	"github.com/uchupx/pintro-golang/data/model"
)

const findGamesByQuery = `
	SELECT * FROM game;
`

type gameMySQLRepository struct {
	db *sqlx.DB
}

func (m gameMySQLRepository) FindByQuery(ctx context.Context, query data.GameQuery) (*data.Collection, error) {
	var games []model.Game

	stmt, err := m.db.PrepareContext(ctx, findGamesByQuery)
	if err != nil {
		return nil, err
	}

	rows, err := stmt.QueryContext(ctx)
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

func NewGameMysqlRepo(db *sqlx.DB) gameMySQLRepository {
	return gameMySQLRepository{
		db: db,
	}
}
