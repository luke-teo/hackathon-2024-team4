package single

import (
	"context"
	"go_chi_template/internal/app/dto"

	. "github.com/go-jet/jet/v2/postgres"
	"github.com/go-jet/jet/v2/qrm"
)

func GetDepartments(ctx context.Context, db qrm.DB) ([]dto.DepartmentWithAncestors, error) {
	stmt := RawStatement(`
      SELECT
        department.id AS "department.id",
        department.name AS "department.name",
        department.custom_id AS "department.custom_id",
        department.hierarchy AS "department.hierarchy",
        department.created_at AS "department.created_at",
        department.updated_at AS "department.updated_at",
        ancestor.id AS "ancestor.id",
        ancestor.name AS "ancestor.name",
        ancestor.custom_id AS "ancestor.custom_id",
        ancestor.hierarchy AS "ancestor.hierarchy",
        ancestor.created_at AS "ancestor.created_at",
        ancestor.updated_at AS "ancestor.updated_at"
      FROM     department
      JOIN     regexp_split_to_table(department.hierarchy::text, '[.]') WITH ordinality t(hierarchy, ord)
      ON       true
      JOIN     department AS ancestor
      ON       ancestor.custom_id = t.hierarchy::text
      ORDER BY department.id, t.ord
    `)
	rows := []dto.DepartmentWithAncestors{}
	err := stmt.QueryContext(ctx, db, &rows)

	return rows, err
}
