package app

import (
	"context"
	"log"
	"net"

	"github.com/dbulyk/platform_common/pkg/closer"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"chat_server/internal/config"
	desc "chat_server/pkg/chat_server_v1"
)

// App является структурой для описания модели сервера
type App struct {
	sp         *serviceProvider
	grpcServer *grpc.Server
}

// NewApp инициализирует объект сервера
func NewApp(ctx context.Context) (*App, error) {
	a := &App{}

	err := a.initDeps(ctx)
	if err != nil {
		return nil, err
	}
	return a, nil
}

// Run запускает сервер по настроенным конфигам
func (a *App) Run() error {
	defer func() {
		closer.CloseAll()
		closer.Wait()
	}()

	return a.start()
}

func (a *App) initDeps(ctx context.Context) error {
	inits := []func(context.Context) error{
		a.initConfig,
		a.initServiceProvider,
		a.initGRPCServer,
	}

	for _, init := range inits {
		if err := init(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (a *App) initConfig(_ context.Context) error {
	err := config.Load(".env")
	if err != nil {
		return err
	}
	return nil
}

func (a *App) initServiceProvider(_ context.Context) error {
	a.sp = newServiceProvider()
	return nil
}

func (a *App) initGRPCServer(ctx context.Context) error {
	a.grpcServer = grpc.NewServer()
	reflection.Register(a.grpcServer)
	desc.RegisterChatServerV1Server(a.grpcServer, a.sp.ChatImplementation(ctx))
	return nil
}

func (a *App) start() error {
	log.Printf("Сервер запускается на %s", a.sp.GRPCConfig().Address())

	list, err := net.Listen("tcp", a.sp.GRPCConfig().Address())
	if err != nil {
		return err
	}

	err = a.grpcServer.Serve(list)
	if err != nil {
		return err
	}

	return nil
}
