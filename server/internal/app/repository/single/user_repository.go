package single

import (
	"context"
	"go_chi_template/generated/db/go_chi_template/public/model"

	"go_chi_template/generated/db/go_chi_template/public/table"

	. "github.com/go-jet/jet/v2/postgres"
	"github.com/go-jet/jet/v2/qrm"
)

func GetUserByEmail(ctx context.Context, db qrm.DB, email string) (*model.User, error) {
	tbl := table.User
	stmt := SELECT(
		tbl.AllColumns,
	).FROM(tbl).
		WHERE(LOWER(tbl.Email).EQ(LOWER(String(email)))).
		LIMIT(1)

	rows := []model.User{}
	err := stmt.QueryContext(ctx, db, &rows)

	if len(rows) != 1 {
		return nil, err
	}

	row := rows[0]
	return &row, err
}

func GetUsersByTenantId(ctx context.Context, db qrm.DB, tenantId int64) ([]model.User, error) {
	tbl := table.User
	stmt := SELECT(tbl.AllColumns).FROM(tbl).WHERE(tbl.TenantID.EQ(Int(tenantId)))

	rows := []model.User{}
	err := stmt.QueryContext(ctx, db, &rows)

	return rows, err
}
