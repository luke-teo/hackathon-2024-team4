package single

import (
	"context"
	"go_chi_template/generated/db/go_chi_template/public/model"
	tbl "go_chi_template/generated/db/go_chi_template/public/table"

	. "github.com/go-jet/jet/v2/postgres"
	"github.com/go-jet/jet/v2/qrm"
)

func GetTenantByName(ctx context.Context, db qrm.DB, name string) (*model.Tenant, error) {
	tbl := tbl.Tenant
	stmt := SELECT(tbl.AllColumns).FROM(tbl).WHERE(tbl.Name.EQ(String(name))).LIMIT(1)

	rows := []model.Tenant{}
	err := stmt.QueryContext(ctx, db, &rows)

	if len(rows) != 1 {
		return nil, err
	}

	row := rows[0]
	return &row, nil
}

func GetTenantByShortCode(ctx context.Context, db qrm.DB, shortCode string) (*model.Tenant, error) {
	tbl := tbl.Tenant
	stmt := SELECT(
		tbl.AllColumns,
	).FROM(tbl).
		WHERE(LOWER(tbl.ShortCode).EQ(LOWER(String(shortCode)))).
		LIMIT(1)

	rows := []model.Tenant{}
	err := stmt.QueryContext(ctx, db, &rows)

	if len(rows) != 1 {
		return nil, err
	}

	row := rows[0]
	return &row, nil
}

func GetTenantById(ctx context.Context, db qrm.DB, id int64) (*model.Tenant, error) {
	tbl := tbl.Tenant
	stmt := SELECT(tbl.AllColumns).FROM(tbl).WHERE(tbl.ID.EQ(Int(id))).LIMIT(1)

	rows := []model.Tenant{}
	err := stmt.QueryContext(ctx, db, &rows)

	if len(rows) != 1 {
		return nil, err
	}

	row := rows[0]
	return &row, err
}
