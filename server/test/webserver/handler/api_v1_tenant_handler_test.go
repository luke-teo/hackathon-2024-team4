package handler

import (
	"context"
	"encoding/json"
	"go_chi_template/generated/oapi"
	"go_chi_template/test"
	tf "go_chi_template/test/factory/tenant_factory"
	"testing"
)

func TestListTenant(t *testing.T) {
	// setup
	app := test.SetupTestApp(t)
	oapiClient := test.SetupTestOpenAPIClient(app)
	defer app.DB().Close()

	tf.SeedTenant(app, tf.WithName("Tenant A"), tf.WithShortCode("aaay2"))

	// send request
	authHeaders := test.SetupAuthHeaders(app)
	res, _ := oapiClient.GetApiV1Tenant(context.Background(), authHeaders)

	// assert against body
	if res.StatusCode != 200 {
		t.Fatal()
	}

	var respDto oapi.GetApiV1Tenant200JSONResponse
	json.NewDecoder(res.Body).Decode(&respDto)
	res.Body.Close()

	if len(respDto.Tenants) != 1 {
		t.Fatal()
	}

	if respDto.Tenants[0].ShortCode != "aaay2" {
		t.Fatal()
	}
}

func TestCreateTenant(t *testing.T) {
	// setup
	app := test.SetupTestApp(t)
	oapiClient := test.SetupTestOpenAPIClient(app)
	defer app.DB().Close()

	tf.SeedTenant(app, tf.WithName("Tenant A"), tf.WithShortCode("aaay2"))

	// send request
	reqDto := oapi.PostApiV1TenantJSONRequestBody{
		Name:      "Tenant B",
		ShortCode: "aaay2",
	}

	authHeaders := test.SetupAuthHeaders(app)
	res, _ := oapiClient.PostApiV1Tenant(context.Background(), reqDto, authHeaders)

	// assert against body
	if res.StatusCode != 400 {
		t.FailNow()
	}

	var respDto oapi.PostApiV1Tenant400JSONResponse
	json.NewDecoder(res.Body).Decode(&respDto)
	res.Body.Close()

	if *respDto.Data.ShortCode.UniqueShortCode == "" {
		t.Fatal()
	}
}
