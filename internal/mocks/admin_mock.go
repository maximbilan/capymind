//coverage:ignore file

package mocks

import "context"

type AdminStorageMock struct{}

func (storage AdminStorageMock) GetTotalUserCount(ctx *context.Context) (int64, error) {
	return 100, nil
}

func (storage AdminStorageMock) GetActiveUserCount(ctx *context.Context) (int64, error) {
	return 75, nil
}

func (storage AdminStorageMock) GetTotalNoteCount(ctx *context.Context) (int64, error) {
	return 999, nil
}
