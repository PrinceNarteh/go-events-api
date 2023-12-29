package mongorm

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Model struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	CreatedAt time.Time          `bson:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at"`
}

// func (m *Model) Create(ctx context.Context, db *mongo.Database, collectionName string, model interface{}) error {
//  collection := db.Collection(collectionName)

//  m.CreatedAt = time.Now()
//  m.UpdatedAt = time.Now()

//  res, err := collection.InsertOne(ctx, model)
//  if err != nil {
//   return err
//  }

//  m.ID = res.InsertedID.(primitive.ObjectID)
//  return nil
// }

func (m *Model) Create(ctx context.Context, collectionName string, model interface{}) error {
	col := InitDB().GetCollection(collectionName)

	m.CreatedAt = time.Now()
	m.UpdatedAt = time.Now()

	if res, err := col.InsertOne(ctx, model); err != nil {
		return err
	} else {
		m.ID = res.InsertedID.(primitive.ObjectID)
	}

	return nil
}
