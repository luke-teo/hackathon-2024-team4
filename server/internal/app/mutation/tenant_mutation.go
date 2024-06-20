package mutation

import (
	"context"
	"go_chi_template/generated/db/go_chi_template/public/model"
	"go_chi_template/generated/db/go_chi_template/public/table"

	"github.com/go-jet/jet/v2/qrm"
)

func InsertTenant(ctx context.Context, db qrm.DB, tenant *model.Tenant) (*model.Tenant, error) {
	insertStmt := table.Tenant.INSERT(table.Tenant.MutableColumns).
		MODEL(tenant).
		RETURNING(table.Tenant.AllColumns)

	dest := []model.Tenant{}

	err := insertStmt.QueryContext(ctx, db, &dest)

	if err != nil || len(dest) != 1 {
		return nil, err
	}

	insertedTenant := dest[0]

	return &insertedTenant, nil
}
