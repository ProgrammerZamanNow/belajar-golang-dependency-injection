//go:build wireinject
// +build wireinject

package simple

import "github.com/google/wire"

func InitializedService() (*SimpleService, error) {
	wire.Build(
		NewSimpleRepository, NewSimpleService,
	)
	return nil, nil
}
