package mysql

import (
	"context"
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/uchupx/pintro-golang/data/model"
)

const findGamePublishersByGameIdsQuery = `
	SELECT 
		* 
	FROM game_publisher WHERE game_id IN (%s);
`

const findGamePublishersByPublisherIdsQuery = `
	SELECT 
		* 
	FROM game_publisher WHERE publisher_id IN (%s);
`

type gamePublisherMySQLRepository struct {
	db *sqlx.DB
}

func (m gamePublisherMySQLRepository) FindByGameIds(ctx context.Context, ids []uint64) ([]model.GamePublisher, error) {
	var items []model.GamePublisher
	var args []interface{}

	if len(ids) == 0 {
		return nil, nil
	}

	placeholder := strings.TrimRight(strings.Repeat("?,", len(ids)), ",")
	query := fmt.Sprintf(findGamePublishersByGameIdsQuery, placeholder)

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
		var item model.GamePublisher

		err = rows.Scan(
			&item.Id,
			&item.GameId,
			&item.PublisherId,
		)

		if err != nil {
			return nil, err
		}

		items = append(items, item)
	}

	return items, nil
}

func (m gamePublisherMySQLRepository) FindByPublisherIds(ctx context.Context, ids []uint64) ([]model.GamePublisher, error) {
	var items []model.GamePublisher
	var args []interface{}

	if len(ids) == 0 {
		return nil, nil
	}

	placeholder := strings.TrimRight(strings.Repeat("?,", len(ids)), ",")
	query := fmt.Sprintf(findGamePublishersByPublisherIdsQuery, placeholder)

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
		var item model.GamePublisher

		err = rows.Scan(
			&item.Id,
			&item.GameId,
			&item.PublisherId,
		)

		if err != nil {
			return nil, err
		}

		items = append(items, item)
	}

	return items, nil
}

func NewGamePublisherMysqlRepo(db *sqlx.DB) gamePublisherMySQLRepository {
	return gamePublisherMySQLRepository{
		db: db,
	}
}
