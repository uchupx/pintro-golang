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

const deleteGameByIdQuery = `
	DELETE FROM game WHERE id = ?;
`

const updateGameByIdQuery = `
	UPDATE game SET game_name=?, genre_id=? WHERE id=?;
`
const insertGameQuery = `
	INSERT INTO game(game_name,genre_id) VALUES(?,?)
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
	query := fmt.Sprintf(findGamesByIdsQuery, placeholder)

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

func (m gameMySQLRepository) Delete(ctx context.Context, game model.Game) (*int64, error) {
	stmt, err := m.db.PrepareContext(ctx, deleteGameByIdQuery)
	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	row, err := stmt.ExecContext(ctx, game.Id)
	if err != nil {
		return nil, err
	}

	rowsAffected, err := row.RowsAffected()
	if err != nil {
		return nil, err
	}

	return &rowsAffected, nil
}

func (m gameMySQLRepository) Update(ctx context.Context, game model.Game) (*int64, error) {
	stmt, err := m.db.PrepareContext(ctx, updateGameByIdQuery)
	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	row, err := stmt.ExecContext(ctx, game.GameName, game.GenreId, game.Id)
	if err != nil {
		return nil, err
	}

	rowsAffected, err := row.RowsAffected()
	if err != nil {
		return nil, err
	}

	return &rowsAffected, nil
}

func (m gameMySQLRepository) Insert(ctx context.Context, game model.Game) (*int64, error) {
	stmt, err := m.db.PrepareContext(ctx, insertGameQuery)
	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	row, err := stmt.ExecContext(ctx, game.GameName, game.GenreId)
	if err != nil {
		return nil, err
	}

	rowsAffected, err := row.RowsAffected()
	if err != nil {
		return nil, err
	}

	return &rowsAffected, nil
}

func NewGameMysqlRepo(db *sqlx.DB) gameMySQLRepository {
	return gameMySQLRepository{
		db: db,
	}
}
