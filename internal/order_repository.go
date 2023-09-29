package internal

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type OrderRepository struct {
	database *mongo.Database
}

func NewOrderRepository(database *mongo.Database) *OrderRepository {
	return &OrderRepository{database: database}
}

func (r *OrderRepository) Create(ctx context.Context, order *Order) error {
	order.ID = primitive.NewObjectID().Hex()
	for i := range order.Items {
		order.Items[i].ID = primitive.NewObjectID().Hex()
	}

	_, err := r.database.Collection("orders").InsertOne(ctx, order)
	return err
}

func (r *OrderRepository) FindAll(ctx context.Context) ([]Order, error) {
	cursor, err := r.database.Collection("orders").Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var orders []Order
	for cursor.Next(ctx) {
		var order Order
		if err := cursor.Decode(&order); err != nil {
			return nil, err
		}
		orders = append(orders, order)
	}
	return orders, nil
}

func (r *OrderRepository) FindOne(ctx context.Context, id string) (Order, error) {
	var order Order
	if err := r.database.Collection("orders").FindOne(ctx, bson.M{"_id": id}).Decode(&order); err != nil {
		return Order{}, err
	}
	return order, nil
}