package validateusercreate

import (
	"context"
	"encoding/json"
	"go_chi_template/config"
	"go_chi_template/generated/oapi"
	"go_chi_template/internal/app/enum"
	"go_chi_template/internal/app/repository/single"
	"strconv"

	"github.com/gookit/validate"
)

func Handle(
	ctx context.Context,
	app *config.App,
	payload *oapi.PostApiV1UserJSONRequestBody,
) (*oapi.PostApiV1User400JSONResponse, error) {
	v := validate.Struct(payload)

	v.AddValidator("uniqueEmail", func(val interface{}) bool {
		existingUser, _ := single.GetUserByEmail(ctx, app.DB(), payload.Email)
		return existingUser == nil
	})
	v.AddValidator("tenantExists", func(val interface{}) bool {
		tenantId, _ := strconv.ParseInt(payload.TenantId, 10, 64)
		existingTenant, _ := single.GetTenantById(ctx, app.DB(), tenantId)
		return existingTenant == nil
	})
	v.AddMessages(map[string]string{
		"uniqueEmail": "Another user with this email already exists",
	})

	v.AddRule("name", "required")
	v.AddRule("name", "minLen", 1)
	v.AddRule("name", "maxLen", 50)

	v.AddRule("email", "required")
	v.AddRule("email", "maxLen", 255)
	v.AddRule("email", "email")
	v.AddRule("email", "uniqueEmail")

	v.AddRule("tenantId", "tenantExists")

	if v.Validate() {
		return nil, nil
	} else {
		// convert errors from map into OpenAPI struct
		// non-present errors will be omitted or retained depending on OpenAPI spec
		errJson, err := json.Marshal(v.Errors.All())

		if err != nil {
			return nil, err
		}

		var errDto oapi.UserCreateValidationError
		err = json.Unmarshal(errJson, &errDto)

		if err != nil {
			return nil, err
		}

		errorEnum := enum.ValidationFailedErrorEnum()
		errResp := oapi.PostApiV1User400JSONResponse{
			ErrorCode:    errorEnum.Code,
			ErrorMessage: errorEnum.Message,
			Data:         errDto,
		}

		return &errResp, nil
	}

}
