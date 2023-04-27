package mysql

import (
	"context"
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/uchupx/pintro-golang/data/model"
)

const findPublishersByQuery = `
	SELECT 
		* 
	FROM publisher
`

const findPublisherByIdsQuery = `
	SELECT * FROM publisher where id IN (%s)
`

type publisherMySQLRepository struct {
	db *sqlx.DB
}

func (m publisherMySQLRepository) FindAll(ctx context.Context) ([]model.Publisher, error) {
	var pubishers []model.Publisher

	stmt, err := m.db.PrepareContext(ctx, findPublishersByQuery)
	if err != nil {
		return nil, err
	}

	rows, err := stmt.QueryContext(ctx)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var publisher model.Publisher

		err = rows.Scan(
			&publisher.Id,
			&publisher.PublisherName,
		)

		if err != nil {
			return nil, err
		}

		pubishers = append(pubishers, publisher)
	}

	return pubishers, nil
}

func (m publisherMySQLRepository) FindByIds(ctx context.Context, ids []uint64) ([]model.Publisher, error) {
	var publishers []model.Publisher
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
		var publisher model.Publisher

		err = rows.Scan(
			&publisher.Id,
			&publisher.PublisherName,
		)

		if err != nil {
			return nil, err
		}

		publishers = append(publishers, publisher)
	}

	return publishers, nil
}

func NewPublisherMysqlRepo(db *sqlx.DB) publisherMySQLRepository {
	return publisherMySQLRepository{
		db: db,
	}
}
