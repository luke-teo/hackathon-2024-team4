package multi

import (
	"context"
	"go_chi_template/generated/db/go_chi_template/public/table"
	"go_chi_template/internal/app/dto"

	. "github.com/go-jet/jet/v2/postgres"
	"github.com/go-jet/jet/v2/qrm"
)

func GetUserWithTenantById(ctx context.Context, db qrm.DB, id int) (*dto.UserWithTenantDto, error) {
	tenantTbl := table.Tenant
	userTbl := table.User

	stmt := SELECT(
		userTbl.AllColumns,
		tenantTbl.AllColumns,
	).FROM(tenantTbl.INNER_JOIN(userTbl, userTbl.TenantID.EQ(tenantTbl.ID)))

	rows := []dto.UserWithTenantDto{}
	err := stmt.QueryContext(ctx, db, &rows)

	if len(rows) != 1 {
		return nil, err
	}

	return &rows[0], err
}
