package models

import (
	"strconv"
	"testing"
	"time"

	"sukasaair/dto"
	mocks "sukasaair/repository/mock"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func TestReserveSeat(t *testing.T) {
	mockDB := new(mocks.Database)
	mockRedis := new(mocks.RedisClientInterface)
	dbInstance = mockDB
	SetRedisClient(mockRedis)

	seatNumber := 101
	lockKey := "seat:" + strconv.Itoa(seatNumber)
	lockTTL := 5 * time.Second

	t.Run("Successful seat reservation", func(t *testing.T) {
		mockRedis.On("SetNX", mock.Anything, lockKey, "locked", lockTTL).Return(true, nil)
		mockRedis.On("Del", mock.Anything, lockKey).Return(nil)

		mockDB.On("FindOne", mock.Anything, bson.M{"seatNumber": seatNumber}).Return(&dto.Seat{Reserved: false}, nil)
		mockDB.On("UpdateOne", mock.Anything, bson.M{"seatNumber": seatNumber}, mock.Anything).Return(nil)

		err := ReserveSeat(seatNumber, "9876543210", "John Doe", 30)
		assert.NoError(t, err)
	})

	t.Run("Seat already reserved", func(t *testing.T) {
		mockRedis.ExpectedCalls = nil // Reset mock expectations
		mockRedis.On("SetNX", mock.Anything, lockKey, "locked", lockTTL).Return(true, nil)
		mockRedis.On("Del", mock.Anything, lockKey).Return(nil)

		mockDB.ExpectedCalls = nil
		mockDB.On("FindOne", mock.Anything, bson.M{"seatNumber": seatNumber}).Return(&dto.Seat{Reserved: true}, nil)

		err := ReserveSeat(seatNumber, "9876543210", "John Doe", 30)
		assert.Error(t, err)
		assert.Equal(t, "seat already reserved", err.Error())
	})

	t.Run("Seat not found", func(t *testing.T) {
		mockRedis.ExpectedCalls = nil
		mockRedis.On("SetNX", mock.Anything, lockKey, "locked", lockTTL).Return(true, nil)
		mockRedis.On("Del", mock.Anything, lockKey).Return(nil)

		mockDB.ExpectedCalls = nil
		mockDB.On("FindOne", mock.Anything, bson.M{"seatNumber": seatNumber}).Return(nil, mongo.ErrNoDocuments)

		err := ReserveSeat(seatNumber, "9876543210", "John Doe", 30)
		assert.Error(t, err)
		assert.Equal(t, "seat not found", err.Error())
	})

	t.Run("Redis lock failure", func(t *testing.T) {
		mockRedis.ExpectedCalls = nil
		mockRedis.On("SetNX", mock.Anything, lockKey, "locked", lockTTL).Return(false, nil)

		err := ReserveSeat(seatNumber, "9876543210", "John Doe", 30)
		assert.Error(t, err)
		assert.Equal(t, "seat is being reserved, try again later", err.Error())
	})
}

func TestResetSeats(t *testing.T) {
	mockDB := new(mocks.Database)
	dbInstance = mockDB

	mockDB.On("UpdateMany", mock.Anything, bson.M{}, bson.M{"$set": bson.M{"reserved": false}}).Return(nil)

	ResetSeats()

	mockDB.AssertExpectations(t)
}
