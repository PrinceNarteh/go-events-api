package mongorm

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Model struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	CreatedAt time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at" json:"updated_at"`
}

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
