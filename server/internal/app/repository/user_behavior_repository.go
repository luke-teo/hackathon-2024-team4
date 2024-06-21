package repository

import (
	"context"

	. "github.com/go-jet/jet/v2/postgres"
	"github.com/go-jet/jet/v2/qrm"

	"first_move/generated/db/first_move/public/model"
	"first_move/generated/db/first_move/public/table"
)

func GetUserBehaviorByUserId(
	ctx context.Context,
	db qrm.DB,
	userId string,
) ([]model.UserBehavior, error) {
	tbl := table.UserBehavior

	stmt := SELECT(
		tbl.AllColumns,
	).FROM(tbl).
		WHERE(tbl.UserID.EQ(String(userId))).
		ORDER_BY(tbl.Date.ASC())

	rows := []model.UserBehavior{}
	err := stmt.QueryContext(ctx, db, &rows)
	if err != nil {
		return nil, err
	}

	return rows, nil
}
