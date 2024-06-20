package task_test

import (
	"go_chi_template/config/provider"
	"go_chi_template/internal/app/task"
	"go_chi_template/test"
	"testing"
)

func TestMain(t *testing.T) {
	// setup
	app := test.SetupTestApp(t)

	err := task.DispatchTenantCleanup(app, "Test")

	if err != nil {
		t.Fatal()
	}

	mockClient := app.Queue().Client.(*provider.MockAsynqClient)

	if len(mockClient.Tasks) != 1 {
		t.Fatal()
	}
}
