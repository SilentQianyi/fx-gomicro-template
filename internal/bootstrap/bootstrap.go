package bootstrap

import (
	"chat/internal/bootstrap/asynq"
	"chat/internal/bootstrap/logger"
	"chat/internal/bootstrap/mongo"
	"chat/internal/bootstrap/server"
	"chat/internal/config"
	"chat/internal/db"
	"chat/internal/scheduler"
	"chat/internal/service"
	"context"
	"fmt"
	"github.com/gookit/color"
	"go-micro.dev/v4"
	"go.uber.org/fx"
	"os"
)

func Init(name, version, banner string) {
	color.Red.Println(banner)
	appId, _ := os.Hostname()
	_ = os.Setenv("SERVICE_ID", appId)
	_ = os.Setenv("SERVICE_NAME", name)
	_ = os.Setenv("SERVICE_VERSION", version)
	println("  Service Name:    " + name)
	println("  Service Version: " + version)
	println()
	println("🚀 starting service...")
}

func Run() {
	var srv micro.Service
	container := fx.New(
		config.Module,
		db.Module,
		service.Module,
		fx.Options(
			logger.Module,
			asynq.Module,
			mongo.Module,
			server.Module,
			scheduler.Module,
		),
		fx.Populate(&srv),
	)

	if err := container.Start(context.TODO()); err != nil {
		fmt.Printf("❗ service started failed: %s\r\n", err.Error())
		return
	}
	fmt.Println("srv run")
	if err := srv.Run(); err != nil {
		fmt.Printf("❗ service exited: %s\r\n", err.Error())
		return
	}
}
