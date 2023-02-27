package do

import (
	"chat/internal/domain/enum"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type HelloWorld struct {
	Id        primitive.ObjectID    `bson:"_id,omitempty"`
	Language  string                `bson:"language,omitempty"`
	Status    enum.HelloWorldStatus `bson:"status,omitempty"`
	Deleted   bool                  `bson:"deleted"`
	endTime   time.Time             `bson:"endTime,omitempty"`
	startTime time.Time             `bson:"startTime,omitempty"`
}
