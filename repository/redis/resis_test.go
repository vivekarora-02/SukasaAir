package redis

import (
	"context"
	"testing"
	"time"

	mocks "sukasaair/repository/mock"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestRedisClient(t *testing.T) {
	mockRedis := new(mocks.RedisClientInterface)
	ctx := context.Background()

	mockRedis.On("SetNX", ctx, "test_key", "test_value", mock.Anything).Return(true, nil)

	mockRedis.On("SetNX", ctx, "test_key", "new_value", mock.Anything).Return(false, nil)

	mockRedis.On("Del", ctx, "test_key").Return(nil)

	success, err := mockRedis.SetNX(ctx, "test_key", "test_value", time.Minute)
	assert.NoError(t, err)
	assert.True(t, success)

	success, err = mockRedis.SetNX(ctx, "test_key", "new_value", time.Minute)
	assert.NoError(t, err)
	assert.False(t, success)

	err = mockRedis.Del(ctx, "test_key")
	assert.NoError(t, err)

	mockRedis.AssertExpectations(t)
}
