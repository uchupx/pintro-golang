package mysql

import (
	"context"
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/uchupx/pintro-golang/data/model"
)

const findPlatformsByQuery = `
	SELECT 
		* 
	FROM platform
`

const findPlatformByIdsQuery = `
	SELECT * FROM platform where id IN (%s)
`

type platformMySQLRepository struct {
	db *sqlx.DB
}

func (m platformMySQLRepository) FindAll(ctx context.Context) ([]model.Platform, error) {
	var items []model.Platform

	stmt, err := m.db.PrepareContext(ctx, findPlatformsByQuery)
	if err != nil {
		return nil, err
	}

	rows, err := stmt.QueryContext(ctx)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var item model.Platform

		err = rows.Scan(
			&item.Id,
			&item.PlatformName,
		)

		if err != nil {
			return nil, err
		}

		items = append(items, item)
	}

	return items, nil
}

func (m platformMySQLRepository) FindByIds(ctx context.Context, ids []uint64) ([]model.Platform, error) {
	var items []model.Platform
	var args []interface{}

	if len(ids) == 0 {
		return nil, nil
	}

	placeholder := strings.TrimRight(strings.Repeat("?,", len(ids)), ",")
	query := fmt.Sprintf(findPlatformByIdsQuery, placeholder)

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
		var item model.Platform

		err = rows.Scan(
			&item.Id,
			&item.PlatformName,
		)

		if err != nil {
			return nil, err
		}

		items = append(items, item)
	}

	return items, nil
}

func NewPlatformMysqlRepo(db *sqlx.DB) platformMySQLRepository {
	return platformMySQLRepository{
		db: db,
	}
}
