package validatetenantcreate

import (
	"context"
	"encoding/json"
	"go_chi_template/config"
	"go_chi_template/generated/oapi"
	"go_chi_template/internal/app/enum"
	"go_chi_template/internal/app/repository/single"

	"github.com/gookit/validate"
)

func Handle(
	ctx context.Context,
	app *config.App,
	payload *oapi.PostApiV1TenantJSONRequestBody,
) (*oapi.PostApiV1Tenant400JSONResponse, error) {
	v := validate.Struct(payload)

	// add custom validators unique to this app service
	// require app so we can access the DB and perform DB validations
	v.AddValidator("uniqueName", func(val interface{}) bool {
		existingTenant, _ := single.GetTenantByName(ctx, app.DB(), payload.Name)
		return existingTenant == nil
	})
	v.AddValidator("uniqueShortCode", func(val interface{}) bool {
		existingTenant, _ := single.GetTenantByShortCode(ctx, app.DB(), payload.ShortCode)
		return existingTenant == nil
	})
	v.AddMessages(map[string]string{
		"uniqueName":      "Another tenant with this name already exists",
		"uniqueShortCode": "Another tenant with this short code already exists",
	})

	// add rules for validation
	// can use struct field names or JSON name, but prefer JSON names
	// struct validation is possible but avoided because our structs are generated from OpenAPI
	v.AddRule("name", "required")
	v.AddRule("name", "minLen", 1)
	v.AddRule("name", "maxLen", 255)
	v.AddRule("name", "uniqueName")

	v.AddRule("shortCode", "required")
	v.AddRule("shortCode", "minLen", 4)
	v.AddRule("shortCode", "maxLen", 5)
	v.AddRule("shortCode", "alphaNum")
	v.AddRule("shortCode", "uniqueShortCode")

	if v.Validate() {
		return nil, nil
	} else {
		return renderResp(v)
	}
}

func renderResp(v *validate.Validation) (*oapi.PostApiV1Tenant400JSONResponse, error) {
	// convert errors from map into OpenAPI struct
	// non-present errors will be omitted or retained depending on OpenAPI spec
	errJson, err := json.Marshal(v.Errors.All())

	if err != nil {
		return nil, err
	}

	var errDto oapi.TenantCreateValidationError
	err = json.Unmarshal(errJson, &errDto)

	if err != nil {
		return nil, err
	}

	errorEnum := enum.ValidationFailedErrorEnum()
	errResp := oapi.PostApiV1Tenant400JSONResponse{
		ErrorCode:    errorEnum.Code,
		ErrorMessage: errorEnum.Message,
		Data:         errDto,
	}

	return &errResp, nil
}
