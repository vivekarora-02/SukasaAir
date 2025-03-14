package mongodb

import (
	"context"
	"errors"
	"os"
	"sukasaair/dto"
	mocks "sukasaair/repository/mock"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestNewMongoDB(t *testing.T) {
	os.Setenv("MONGO_URI", "mongodb://localhost:27017")

	db, err := NewMongoDB()

	assert.NoError(t, err)
	assert.NotNil(t, db)
}

// Test InsertOne using mockery
func TestInsertOne(t *testing.T) {
	mockDB := new(mocks.Database)
	seat := dto.Seat{SeatNumber: 101, Reserved: false}

	mockDB.On("InsertOne", mock.Anything, seat).Return(nil)

	err := mockDB.InsertOne(context.TODO(), seat)
	assert.NoError(t, err)

	mockDB.AssertExpectations(t)
}

// Test InsertMany using mockery
func TestInsertMany(t *testing.T) {
	mockDB := new(mocks.Database)
	seats := []interface{}{
		dto.Seat{SeatNumber: 1, Reserved: false},
		dto.Seat{SeatNumber: 2, Reserved: false},
	}

	mockDB.On("InsertMany", mock.Anything, seats).Return(nil)

	err := mockDB.InsertMany(context.TODO(), seats)
	assert.NoError(t, err)

	mockDB.AssertExpectations(t)
}

// Test DeleteMany using mockery
func TestDeleteMany(t *testing.T) {
	mockDB := new(mocks.Database)

	mockDB.On("DeleteMany", mock.Anything, mock.Anything).Return(nil)

	err := mockDB.DeleteMany(context.TODO(), nil)
	assert.NoError(t, err)

	mockDB.AssertExpectations(t)
}

// Test FindOne when seat exists
func TestFindOne_SeatExists(t *testing.T) {
	mockDB := new(mocks.Database)
	mockSeat := &dto.Seat{SeatNumber: 101, Reserved: false}

	mockDB.On("FindOne", mock.Anything, mock.Anything).Return(mockSeat, nil)

	seat, err := mockDB.FindOne(context.TODO(), map[string]interface{}{"seatNumber": 101})

	assert.NoError(t, err)
	assert.NotNil(t, seat)
	assert.Equal(t, 101, seat.SeatNumber)

	mockDB.AssertExpectations(t)
}

// Test FindOne when seat does not exist
func TestFindOne_SeatNotFound(t *testing.T) {
	mockDB := new(mocks.Database)

	mockDB.On("FindOne", mock.Anything, mock.Anything).Return((*dto.Seat)(nil), errors.New("not found"))

	seat, err := mockDB.FindOne(context.TODO(), map[string]interface{}{"seatNumber": 404})

	assert.Error(t, err)
	assert.Nil(t, seat)

	mockDB.AssertExpectations(t)
}

// Test UpdateOne using mockery
func TestUpdateOne(t *testing.T) {
	mockDB := new(mocks.Database)

	mockDB.On("UpdateOne", mock.Anything, mock.Anything, mock.Anything).Return(nil)

	err := mockDB.UpdateOne(context.TODO(), map[string]interface{}{"seatNumber": 101}, map[string]interface{}{"reserved": true})
	assert.NoError(t, err)

	mockDB.AssertExpectations(t)
}

// Test UpdateMany using mockery
func TestUpdateMany(t *testing.T) {
	mockDB := new(mocks.Database)

	mockDB.On("UpdateMany", mock.Anything, mock.Anything, mock.Anything).Return(nil)

	err := mockDB.UpdateMany(context.TODO(), nil, map[string]interface{}{"reserved": false})
	assert.NoError(t, err)

	mockDB.AssertExpectations(t)
}

func TestInitSeats(t *testing.T) {
	mockDB := new(mocks.Database)
	mockDB.On("InitSeats").Return(nil)
	mockDB.InitSeats()
	mockDB.AssertExpectations(t)
}
