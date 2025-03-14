package mongodb

import (
	"context"
	"fmt"
	"log"
	"os"

	"sukasaair/constants"
	"sukasaair/dto"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Database interface {
	InsertOne(ctx context.Context, document interface{}) error
	InsertMany(ctx context.Context, documents []interface{}) error
	DeleteMany(ctx context.Context, filter interface{}) error
	FindOne(ctx context.Context, filter interface{}) (*dto.Seat, error)
	UpdateOne(ctx context.Context, filter interface{}, update interface{}) error
	UpdateMany(ctx context.Context, filter interface{}, update interface{}) error
	Client() *mongo.Client
	Collection(name string) *mongo.Collection
	InitSeats()
}

type MongoDB struct {
	client *mongo.Client
	db     *mongo.Database
}

func (m *MongoDB) Client() *mongo.Client {
	return m.client
}

func NewMongoDB() (Database, error) {
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: .env file not loaded, using environment variables instead")
	}

	mongoURI := os.Getenv("MONGO_URI")
	if mongoURI == "" {
		return nil, fmt.Errorf("MONGO_URI is not set")
	}

	clientOptions := options.Client().ApplyURI(mongoURI)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to MongoDB: %v", err)
	}

	if err := client.Ping(context.TODO(), nil); err != nil {
		return nil, fmt.Errorf("MongoDB is not reachable: %v", err)
	}

	fmt.Println("✅ Connected to MongoDB!")
	return &MongoDB{
		client: client,
		db:     client.Database("sukasaair"),
	}, nil
}

func (m *MongoDB) Collection(name string) *mongo.Collection {
	return m.db.Collection(name)
}

func (m *MongoDB) InsertOne(ctx context.Context, document interface{}) error {
	_, err := m.Collection("seats").InsertOne(ctx, document)
	return err
}

func (m *MongoDB) InsertMany(ctx context.Context, documents []interface{}) error {
	_, err := m.Collection("seats").InsertMany(ctx, documents)
	return err
}

func (m *MongoDB) DeleteMany(ctx context.Context, filter interface{}) error {
	_, err := m.Collection("seats").DeleteMany(ctx, filter)
	return err
}

func (m *MongoDB) FindOne(ctx context.Context, filter interface{}) (*dto.Seat, error) {
	var seat dto.Seat
	err := m.Collection("seats").FindOne(ctx, filter).Decode(&seat)
	if err != nil {
		return nil, err
	}
	return &seat, nil
}

func (m *MongoDB) UpdateOne(ctx context.Context, filter interface{}, update interface{}) error {
	_, err := m.Collection("seats").UpdateOne(ctx, filter, update)
	return err
}

func (m *MongoDB) UpdateMany(ctx context.Context, filter interface{}, update interface{}) error {
	_, err := m.Collection("seats").UpdateMany(ctx, filter, update)
	return err
}

func (m *MongoDB) InitSeats() {
	m.DeleteMany(context.TODO(), nil)

	var seats []interface{}
	for i := 1; i <= constants.MAX_SEATS_TO_BE_RESERVED; i++ {
		seats = append(seats, dto.Seat{SeatNumber: i, Reserved: false})
	}

	m.InsertMany(context.TODO(), seats)
}
