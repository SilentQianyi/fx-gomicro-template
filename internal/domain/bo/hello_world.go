package bo

import (
	boCommon "chat/internal/domain/bo/common"
	"github.com/samber/lo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type HelloWorldQueryByLanguage struct {
	Language string

	Page      *boCommon.PageQuery
	endTime   time.Time
	startTime time.Time
}

func NewHelloWorldQueryByLanguage(language string) *HelloWorldQueryByLanguage {
	return &HelloWorldQueryByLanguage{
		Language: language,
	}
}

func (c *HelloWorldQueryByLanguage) SetPage(pageNo, pageSize int64) *HelloWorldQueryByLanguage {
	c.Page = boCommon.NewPageQuery(pageNo, pageSize)
	return c
}

func (c *HelloWorldQueryByLanguage) SetStartTime(startTime time.Time) *HelloWorldQueryByLanguage {
	c.startTime = startTime
	return c
}

func (c *HelloWorldQueryByLanguage) GetQuery() bson.D {
	queryBson := bson.D{}
	if lo.IsNotEmpty(c.Language) {
		queryBson = append(queryBson, bson.E{Key: "language", Value: c.Language})
	}
	return queryBson
}

func (c *HelloWorldQueryByLanguage) GetOptions() *options.FindOptions {
	opts := options.Find()
	if lo.IsNotEmpty(c.startTime) {
		sortBson := bson.D{
			bson.E{Key: "startTime", Value: 1},
		}
		opts.SetSort(sortBson)
	}
	if lo.IsNotEmpty(c.Page) {
		opts.SetSkip(c.Page.Offset()).SetLimit(c.Page.Limit())
	}
	return opts
}
