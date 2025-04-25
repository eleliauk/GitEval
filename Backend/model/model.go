package model

import "github.com/google/wire"

var ProviderSet = wire.NewSet(
	NewData,
	NewDB,
	NewGormUserDAO,
	NewGormDomainDAO,
	NewGormContactDAO,
)
