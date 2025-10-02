package integration

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/unknwon/go-mockgen/internal/integration/testdata/mocks"
)

func TestStrictConstructor(t *testing.T) {
	// All invocations panic by default
	mock := mocks.NewStrictMockRetrier()

	assert.Panics(t, func() {
		_ = mock.Retry(context.Background(), func() error {
			return nil
		})
	})

	// Should not panic if overwritten
	mock.RetryFunc.SetDefaultReturn(nil)

	assert.NotPanics(t, func() {
		_ = mock.Retry(context.Background(), func() error {
			return nil
		})
	})
}
