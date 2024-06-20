package multi

import (
	"context"
	"go_chi_template/generated/db/go_chi_template/public/table"
	"go_chi_template/internal/app/dto"

	. "github.com/go-jet/jet/v2/postgres"
	"github.com/go-jet/jet/v2/qrm"
)

func GetTenantsWithUserCount(ctx context.Context, db qrm.DB) ([]dto.TenantWithUserCountDto, error) {
	tenantTbl := table.Tenant
	userTbl := table.User

	stmt := SELECT(
		tenantTbl.AllColumns,
		COALESCE(COUNT(userTbl.ID), Int(0)).AS("tenant.user_count"),
	).FROM(tenantTbl.LEFT_JOIN(userTbl, userTbl.TenantID.EQ(tenantTbl.ID))).
		GROUP_BY(
			// specify all columns since it is required by PostgreSQL group by
			// TODO: figure out how to get this from table.Tenant.AllColumns
			tenantTbl.ID,
			tenantTbl.Name,
			tenantTbl.ShortCode,
			tenantTbl.CreatedAt,
			tenantTbl.UpdatedAt,
		).ORDER_BY(tenantTbl.ShortCode.ASC())

	rows := []dto.TenantWithUserCountDto{}
	err := stmt.QueryContext(ctx, db, &rows)

	return rows, err
}
