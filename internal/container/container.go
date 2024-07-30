package container

import (
	"github.com/nutsp/golang-clean-architecture/config"
	"github.com/nutsp/golang-clean-architecture/internal/app"
	"github.com/nutsp/golang-clean-architecture/internal/handlers"
	"github.com/nutsp/golang-clean-architecture/internal/infastructure/database"
	"github.com/nutsp/golang-clean-architecture/internal/middlewares"
	"github.com/nutsp/golang-clean-architecture/internal/repositories"
	"github.com/nutsp/golang-clean-architecture/internal/usecase"
	httpClient "github.com/nutsp/golang-clean-architecture/pkg/httpclient"
	"github.com/nutsp/golang-clean-architecture/pkg/observability"
	"go.uber.org/dig"
)

type Container struct {
	container *dig.Container
	Error     error
}

type Dependency struct {
	Constructor interface{}
	Interface   interface{}
	Token       string
}

func NewContainer() *Container {
	c := &Container{}
	c.Configure()

	return c
}

func (cn *Container) Run() *Container {
	err := cn.container.Invoke(app.NewApp)
	if err != nil {
		panic(err)
	}

	return cn
}

func (cn *Container) Configure() {
	cn.container = dig.New()

	deps := []Dependency{
		{
			Constructor: config.NewLoadConfig,
		},
		{
			Constructor: func(cfg *config.Config) *observability.ZapLogger {
				return observability.NewZapLogger(cfg.Logger)
			},
			Interface: new(observability.Logger),
			Token:     "Logger",
		},
		{
			Constructor: func(cfg *config.Config) *httpClient.Client {
				return httpClient.NewClient(cfg.HttpClient)
			},
			Interface: new(httpClient.IClient),
			Token:     "HttpClient",
		},
		{
			Constructor: database.NewDatabase,
			Interface:   new(database.IDatabase),
			Token:       "Database",
		},
		{
			Constructor: middlewares.NewMiddleware,
			Interface:   new(middlewares.IMiddleware),
			Token:       "Middleware",
		},
		{
			Constructor: repositories.NewMailerRepository,
			Interface:   new(repositories.IMailerRepository),
			Token:       "MailerRepository",
		},
		{
			Constructor: repositories.NewUserRepository,
			Interface:   new(repositories.IUserRepository),
			Token:       "UserRepository",
		},
		{
			Constructor: usecase.NewUserUsecase,
			Interface:   new(usecase.IUserUsecase),
			Token:       "UserUsecase",
		},
		{
			Constructor: handlers.NewUserHandler,
			Interface:   new(handlers.IUserHandler),
			Token:       "UserHandler",
		},
	}

	for _, dep := range deps {
		var err error
		if dep.Interface != nil {
			err = cn.container.Provide(dep.Constructor, dig.As(dep.Interface), dig.Name(dep.Token))
		} else {
			err = cn.container.Provide(dep.Constructor)
		}
		cn.Error = err
	}
}
