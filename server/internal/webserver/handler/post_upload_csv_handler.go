package handler

import (
	"context"
	"fmt"

	"first_move/generated/oapi"
)

func PostUploadCsv(
	ctx context.Context,
	request oapi.PostUploadCsvRequestObject,
) (oapi.PostUploadCsvResponseObject, error) {
	f, err := request.Body.ReadForm(1 * 1024 * 1024)
	if err != nil {
		return nil, err
	}

	for _, line := range f.Value {
		fmt.Println(line)
	}

	return nil, nil
}
