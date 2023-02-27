package grpc

import (
	"chat/api/common/v1"
	"chat/api/helloWorld/v1"
	"chat/internal/db"
	"chat/internal/service/biz"
	"context"
	"go-micro.dev/v4/server"
)

type helloWorldHandlerImpl struct {
	repo              db.HelloWorldRepository
	helloWorldService biz.HelloWorldService
}

func registerHelloWorldHandler(
	srv server.Server,
	repo db.HelloWorldRepository,
	helloWorldService biz.HelloWorldService,
) error {
	handler := &helloWorldHandlerImpl{
		repo:              repo,
		helloWorldService: helloWorldService,
	}
	return helloWorld.RegisterHelloWorldHandler(srv, handler)
}

func (c *helloWorldHandlerImpl) GetHelloWorldByLanguage(ctx context.Context, req *helloWorld.GetHelloWorldByLanguageRequest, resp *helloWorld.HelloWorldListModel) error {
	language := req.GetLanguage()
	list, err := c.helloWorldService.GetHelloWorldByLanguage(ctx, language)
	if err != nil {
		return err
	}
	resp = &helloWorld.HelloWorldListModel{
		Data: make([]*helloWorld.HelloWorldModel, 0),
	}
	for _, world := range list {
		resp.Data = append(resp.Data, &helloWorld.HelloWorldModel{
			Id:       world.Id.Hex(),
			Common:   &common.HelloWorldCommon{Language: world.Language},
			Language: world.Language,
			Status:   helloWorld.HelloWorldStatus(world.Status),
			Deleted:  world.Deleted,
		})
	}
	return nil
}
