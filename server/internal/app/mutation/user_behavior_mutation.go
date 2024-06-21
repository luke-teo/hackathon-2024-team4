package mutation

import (
	"context"

	"github.com/go-jet/jet/qrm"

	"first_move/generated/db/first_move/public/model"
	"first_move/generated/db/first_move/public/table"
)

func InsertUserBehavior(ctx context.Context, db qrm.DB, userBehaviors []model.UserBehavior) error {
	tbl := table.UserBehavior

	stmt := tbl.INSERT(tbl.AllColumns).MODELS(userBehaviors)

	_, err := stmt.ExecContext(ctx, db)
	if err != nil {
		return err
	}

	return nil
}
