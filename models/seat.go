package models

import (
	"context"
	"errors"
	"strconv"
	"time"

	db "sukasaair/repository/mongodb"
	"sukasaair/repository/redis"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var dbInstance db.Database

func SetDatabase(database db.Database) {
	dbInstance = database
}

var redisClient redis.RedisClientInterface

func SetRedisClient(client redis.RedisClientInterface) {
	redisClient = client
}

func ReserveSeat(seatNumber int, phone, name string, age int) error {
	ctx := context.TODO()
	lockKey := "seat:" + strconv.Itoa(seatNumber)
	lockTTL := 5 * time.Second

	success, err := redisClient.SetNX(ctx, lockKey, "locked", lockTTL)
	if err != nil {
		return err
	}
	if !success {
		return errors.New("seat is being reserved, try again later")
	}
	defer redisClient.Del(ctx, lockKey)

	seat, err := dbInstance.FindOne(ctx, bson.M{"seatNumber": seatNumber})
	if err == mongo.ErrNoDocuments {
		return errors.New("seat not found")
	} else if err != nil {
		return err
	}

	if seat.Reserved {
		return errors.New("seat already reserved")
	}

	err = dbInstance.UpdateOne(ctx, bson.M{"seatNumber": seatNumber}, bson.M{"$set": bson.M{
		"reserved":  true,
		"passenger": name,
		"phone":     phone,
		"age":       age,
	}})
	return err
}

func ResetSeats() {
	dbInstance.UpdateMany(context.TODO(), bson.M{}, bson.M{"$set": bson.M{"reserved": false}})
}
