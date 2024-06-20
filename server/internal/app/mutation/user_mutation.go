package mutation

import (
	"context"
	"go_chi_template/generated/db/go_chi_template/public/model"
	"go_chi_template/generated/db/go_chi_template/public/table"

	"github.com/go-jet/jet/v2/qrm"
)

func InsertUser(ctx context.Context, db qrm.DB, user *model.User) (*model.User, error) {
	insertStmt := table.User.INSERT(table.User.MutableColumns).
		MODEL(user).
		RETURNING(table.User.AllColumns)

	dest := []model.User{}

	err := insertStmt.QueryContext(ctx, db, &dest)

	if err != nil || len(dest) != 1 {
		return nil, err
	}

	insertedUser := dest[0]

	return &insertedUser, nil
}
