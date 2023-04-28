package mysql

import (
	"context"
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/uchupx/pintro-golang/data/model"
)

const findRegionsByQuery = `
	SELECT 
		* 
	FROM region
`

const findRegionByIdsQuery = `
	select region.*, region_sales.game_platform_id, region_sales.num_sales from region join region_sales on region_sales.region_id=region.id where region_sales.game_platform_id IN (%s);
`

type regionMySQLRepository struct {
	db *sqlx.DB
}

func (m regionMySQLRepository) FindAll(ctx context.Context) ([]model.Region, error) {
	var regions []model.Region

	stmt, err := m.db.PrepareContext(ctx, findRegionsByQuery)
	if err != nil {
		return nil, err
	}

	rows, err := stmt.QueryContext(ctx)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var region model.Region

		err = rows.Scan(
			&region.Id,
			&region.RegionName,
		)

		if err != nil {
			return nil, err
		}

		regions = append(regions, region)
	}

	return regions, nil
}

func (m regionMySQLRepository) FindByIds(ctx context.Context, ids []uint64) ([]model.RegionJoinRegionSales, error) {
	var items []model.RegionJoinRegionSales
	var args []interface{}

	if len(ids) == 0 {
		return nil, nil
	}

	placeholder := strings.TrimRight(strings.Repeat("?,", len(ids)), ",")
	query := fmt.Sprintf(findRegionByIdsQuery, placeholder)

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
		var item model.RegionJoinRegionSales

		err = rows.Scan(
			&item.Id,
			&item.RegionName,
			&item.GamePlatformId,
			&item.NumSales,
		)

		if err != nil {
			return nil, err
		}

		items = append(items, item)
	}

	return items, nil
}

func NewRegionMysqlRepo(db *sqlx.DB) regionMySQLRepository {
	return regionMySQLRepository{
		db: db,
	}
}
