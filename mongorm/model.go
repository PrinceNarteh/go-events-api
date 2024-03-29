package mongorm

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Model struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	CreatedAt time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at" json:"updated_at"`
}

func (m *Model) Find(ctx context.Context, collectionName string, filter interface{}) (docs []bson.M, err error) {
	collection := GetCollection(collectionName)

	cur, err := collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)

	var items []bson.M
	for cur.Next(ctx) {
		var item bson.M
		if err := cur.Decode(&item); err != nil {
			return nil, err
		}
		items = append(items, item)
	}
	return items, nil
}

func (m *Model) Create(ctx context.Context, collectionName string, model interface{}) error {
	col := GetCollection(collectionName)

	m.CreatedAt = time.Now()
	m.UpdatedAt = time.Now()

	if res, err := col.InsertOne(ctx, model); err != nil {
		return err
	} else {
		m.ID = res.InsertedID.(primitive.ObjectID)
	}

	return nil
}
