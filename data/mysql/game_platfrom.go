package mysql

import (
	"context"
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/uchupx/pintro-golang/data/model"
)

const findGamePlatformsByPlatformIdsQuery = `
	SELECT 
		* 
	FROM game_platform WHERE platform_id IN (%s);
`

const findGamePlatformsByPublisherIdsQuery = `
	SELECT 
		* 
	FROM game_platform WHERE game_publisher_id IN (%s);
`

type gamePlatformMySQLRepository struct {
	db *sqlx.DB
}

func (m gamePlatformMySQLRepository) FindByPlatformIds(ctx context.Context, ids []uint64) ([]model.GamePlatform, error) {
	var items []model.GamePlatform
	var args []interface{}

	if len(ids) == 0 {
		return nil, nil
	}

	placeholder := strings.TrimRight(strings.Repeat("?,", len(ids)), ",")
	query := fmt.Sprintf(findGamePlatformsByPlatformIdsQuery, placeholder)

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
		var item model.GamePlatform

		err = rows.Scan(
			&item.Id,
			&item.GamePublisherId,
			&item.PlatformId,
			&item.ReleaseYear,
		)

		if err != nil {
			return nil, err
		}

		items = append(items, item)
	}

	return items, nil
}

func (m gamePlatformMySQLRepository) FindByPublisherIds(ctx context.Context, ids []uint64) ([]model.GamePlatform, error) {
	var items []model.GamePlatform
	var args []interface{}

	if len(ids) == 0 {
		return nil, nil
	}

	placeholder := strings.TrimRight(strings.Repeat("?,", len(ids)), ",")
	query := fmt.Sprintf(findGamePlatformsByPublisherIdsQuery, placeholder)

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
		var item model.GamePlatform

		err = rows.Scan(
			&item.Id,
			&item.GamePublisherId,
			&item.PlatformId,
			&item.ReleaseYear,
		)

		if err != nil {
			return nil, err
		}

		items = append(items, item)
	}

	return items, nil
}

func NewGamePlatformMysqlRepo(db *sqlx.DB) gamePlatformMySQLRepository {
	return gamePlatformMySQLRepository{
		db: db,
	}
}
