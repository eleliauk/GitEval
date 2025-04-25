//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/GitEval/GitEval-Backend/api/route"
	"github.com/GitEval/GitEval-Backend/client"
	"github.com/GitEval/GitEval-Backend/conf"
	"github.com/GitEval/GitEval-Backend/controller"
	"github.com/GitEval/GitEval-Backend/middleware"
	"github.com/GitEval/GitEval-Backend/model"
	"github.com/GitEval/GitEval-Backend/model/cache"
	"github.com/GitEval/GitEval-Backend/pkg"
	"github.com/GitEval/GitEval-Backend/pkg/github"
	"github.com/GitEval/GitEval-Backend/service"
	"github.com/google/wire"
)

func WireApp(confPath string) (route.App, func()) {
	panic(wire.Build(
		cache.ProviderSet,
		conf.ProviderSet,
		route.ProviderSet,
		model.ProviderSet,
		service.ProviderSet,
		controller.ProviderSet,
		client.ProviderSet,
		pkg.ProviderSet,
		middleware.ProviderSet,
		wire.Bind(new(middleware.ParTokener), new(*middleware.JWTClient)),
		wire.Bind(new(route.AuthControllerProxy), new(*controller.AuthController)),
		wire.Bind(new(route.UserControllerProxy), new(*controller.UserController)),
		wire.Bind(new(controller.UserServiceProxy), new(*service.UserService)),
		wire.Bind(new(controller.GenerateJWTer), new(*middleware.JWTClient)),
		wire.Bind(new(controller.AuthServiceProxy), new(*service.AuthService)),
		wire.Bind(new(service.GitHubAPIProxy), new(*github.GitHubAPI)),
		wire.Bind(new(service.UserServiceProxy), new(*service.UserService)),
		wire.Bind(new(service.UserDAOProxy), new(*model.GormUserDAO)),
		wire.Bind(new(service.ContactDAOProxy), new(*model.GormContactDAO)),
		wire.Bind(new(service.DomainDAOProxy), new(*model.GormDomainDAO)),
		wire.Bind(new(service.GithubProxy), new(*github.GitHubAPI)),
		wire.Bind(new(service.Transaction), new(*model.Data)),
	))
}
