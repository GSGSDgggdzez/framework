package console

import (
	"github.com/goravel/framework/contracts/config"
	"github.com/goravel/framework/contracts/console"
	"github.com/goravel/framework/contracts/console/command"
	contractsseeder "github.com/goravel/framework/contracts/database/seeder"
	"github.com/goravel/framework/errors"
	"github.com/goravel/framework/support/color"
)

type SeedCommand struct {
	config config.Config
	seeder contractsseeder.Facade
}

func NewSeedCommand(config config.Config, seeder contractsseeder.Facade) *SeedCommand {
	return &SeedCommand{
		config: config,
		seeder: seeder,
	}
}

// Signature The name and signature of the console command.
func (receiver *SeedCommand) Signature() string {
	return "db:seed"
}

// Description The console command description.
func (receiver *SeedCommand) Description() string {
	return "Seed the database with records"
}

// Extend The console command extend.
func (receiver *SeedCommand) Extend() command.Extend {
	return command.Extend{
		Category: "db",
		Flags: []command.Flag{
			&command.BoolFlag{
				Name:    "force",
				Aliases: []string{"f"},
				Usage:   "force the operation to run when in production",
			},
			&command.StringSliceFlag{
				Name:    "seeder",
				Aliases: []string{"s"},
				Usage:   "specify the seeder(s) to run",
			},
		},
	}
}

// Handle executes the console command.
func (receiver *SeedCommand) Handle(ctx console.Context) error {
	force := ctx.OptionBool("force")
	if err := receiver.ConfirmToProceed(force); err != nil {
		color.Red().Println(err)
		return nil
	}

	names := ctx.OptionSlice("seeder")
	seeders, err := receiver.GetSeeders(names)
	if err != nil {
		color.Red().Println(err)
		return nil
	}
	if len(seeders) == 0 {
		color.Red().Println("no seeders found")
		return nil
	}

	if err := receiver.seeder.Call(seeders); err != nil {
		color.Red().Printf("error running seeder: %v\n", err)
	}
	color.Green().Println("Database seeding completed successfully.")

	return nil
}

// ConfirmToProceed determines if the command should proceed based on user confirmation.
func (receiver *SeedCommand) ConfirmToProceed(force bool) error {
	if force || (receiver.config.Env("APP_ENV") != "production") {
		return nil
	}

	return errors.DBForceIsRequiredInProduction
}

// GetSeeders returns a seeder instances
func (receiver *SeedCommand) GetSeeders(names []string) ([]contractsseeder.Seeder, error) {
	if len(names) == 0 {
		return receiver.seeder.GetSeeders(), nil
	}
	var seeders []contractsseeder.Seeder
	for _, name := range names {
		seeder := receiver.seeder.GetSeeder(name)
		if seeder == nil {
			return nil, errors.DBSeederNotFound.Args(name)
		}
		seeders = append(seeders, seeder)
	}
	return seeders, nil
}
