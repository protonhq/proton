package registry

import (
	"github.com/sarulabs/di"

	"github.com/protonhq/proton/infra/config"
	"github.com/protonhq/proton/infra/db"
	"github.com/protonhq/proton/usecase"
)

// Container - container host all depency
type Container struct {
	ctn di.Container
}

// NewContainer - create application container
func NewContainer(conf *config.Configuration) (*Container, error) {
	builder, err := di.NewBuilder()
	if err != nil {
		return nil, err
	}

	// Create DB
	database, err := db.InitDB(conf.Database.ConnectionString())

	if err != nil {
		return nil, err
	}

	if err = builder.Add([]di.Def{
		{
			Name: "account-usecase",
			Build: func(ctn di.Container) (interface{}, error) {
				repo := db.NewAccountRepository(database)
				return usecase.NewAccountUsecase(repo), nil
			},
		},
	}...); err != nil {
		return nil, err
	}

	return &Container{
		ctn: builder.Build(),
	}, nil
}

// Resolve - resolve dependency
func (c *Container) Resolve(name string) interface{} {
	return c.ctn.Get(name)
}

// Clean - clean container
func (c *Container) Clean() error {
	return c.ctn.Clean()
}
