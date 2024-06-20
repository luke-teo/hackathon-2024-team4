package domainservice

import (
	"go_chi_template/generated/db/go_chi_template/public/model"
	"time"
)

func NewTenant(name string, shortCode string) *model.Tenant {
	newTenant := model.Tenant{
		Name:      name,
		ShortCode: shortCode,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	return &newTenant
}
