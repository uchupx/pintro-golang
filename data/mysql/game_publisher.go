package mysql

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/uchupx/pintro-golang/data"
	"github.com/uchupx/pintro-golang/data/model"
)

const findGamePublishersByQuery = `
	SELECT 
		* 
	FROM game_publisher
	LIMIT ?, ?;
`

type gamePublisherMySQLRepository struct {
	db *sqlx.DB
}

func (m gamePublisherMySQLRepository) FindByQuery(ctx context.Context, query data.GamePublisherQuery) (*data.Collection, error) {
	var games []model.GamePublisher
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
		var game model.GamePublisher

		err = rows.Scan(
			&game.Id,
			&game.GameId,
			&game.PublisherId,
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

func NewGamePublisherMysqlRepo(db *sqlx.DB) gamePublisherMySQLRepository {
	return gamePublisherMySQLRepository{
		db: db,
	}
}
