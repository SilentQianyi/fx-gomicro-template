package db

import (
	"chat/internal/domain/bo"
	"chat/internal/domain/do"
	"context"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	helloWorldCollection = "hello_world"
)

type HelloWorldRepository interface {
	GetHelloWorldByLanguage(ctx context.Context, query *bo.HelloWorldQueryByLanguage) ([]*do.HelloWorld, error)
}

type helloWorldRepositoryImpl struct {
	helloWorld *mongo.Collection
}

func NewHelloWorldRepository(db *mongo.Database) HelloWorldRepository {
	return &helloWorldRepositoryImpl{
		helloWorld: db.Collection(helloWorldCollection),
	}
}

func (r *helloWorldRepositoryImpl) GetHelloWorldByLanguage(ctx context.Context, query *bo.HelloWorldQueryByLanguage) ([]*do.HelloWorld, error) {
	queryBson := query.GetQuery()
	opts := query.GetOptions()
	cursor, err := r.helloWorld.Find(ctx, queryBson, opts)
	if err != nil {
		return nil, err
	}

	var data []*do.HelloWorld
	if err = cursor.All(ctx, &data); err != nil {
		return nil, err
	}
	return data, nil
}
