package testing

import (
	"fmt"

	"github.com/goravel/framework/contracts/database/seeder"
	"github.com/goravel/framework/errors"
)

type TestCase struct {
}

func (receiver *TestCase) Seed(seeds ...seeder.Seeder) {
	command := "db:seed"
	if len(seeds) > 0 {
		command += " --seeder"
		for _, seed := range seeds {
			command += fmt.Sprintf(" %s", seed.Signature())
		}
	}

	if artisanFacade == nil {
		panic(errors.ArtisanFacadeNotSet.SetModule(errors.ModuleTesting))
	}

	artisanFacade.Call(command)
}

func (receiver *TestCase) RefreshDatabase(seeds ...seeder.Seeder) {
	if artisanFacade == nil {
		panic(errors.ArtisanFacadeNotSet.SetModule(errors.ModuleTesting))
	}

	artisanFacade.Call("migrate:refresh")
}
