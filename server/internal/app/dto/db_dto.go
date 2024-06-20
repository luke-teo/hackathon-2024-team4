package dto

import (
	"go_chi_template/generated/db/go_chi_template/public/model"
)

type TenantWithUserCountDto struct {
	model.Tenant
	UserCount int `alias:"tenant.user_count"`
}

type UserWithTenantDto struct {
	Tenant model.Tenant
	User   model.User
}

type DepartmentWithAncestors struct {
	model.Department
	Ancestors []model.Department `alias:"ancestor.*"`
}
