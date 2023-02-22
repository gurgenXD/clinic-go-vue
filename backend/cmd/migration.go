package cmd

import (
	"errors"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/spf13/cobra"
)

const (
	migrationsDir = "migrations"

	stepsFlag      = "steps"
	stepsFlagShort = "s"
	stepsDefault   = 0
	stepsUsage     = "Number of migrated versions."
)

var (
	steps uint
)

var (
	migration *migrate.Migrate

	migrationCmd = &cobra.Command{
		Use:   "migration",
		Short: "Database migrations",
		RunE: func(cmd *cobra.Command, args []string) error {
			var err error
			url := "postgresql://postgres:postgres@localhost:5432/postgres?sslmode=disable&search_path=clinic"
			if migration, err = migrate.New("file://"+migrationsDir, url); err != nil {
				return fmt.Errorf("Failed to create connection: %w", err)
			}

			return nil
		},
	}

	migrationUpCmd = &cobra.Command{
		Use:   "up",
		Short: "Migration upgrade",
		RunE: func(cmd *cobra.Command, args []string) error {
			err := cmd.Parent().RunE(cmd.Parent(), args)
			if err != nil {
				return fmt.Errorf("Failed to migrate: %w", err)
			}

			if steps == 0 {
				if err = migration.Up(); err != nil {
					if errors.Is(err, migrate.ErrNoChange) {
						fmt.Println("Migration has no changes.")
						return nil
					}

					return fmt.Errorf("Falied to upgrade: %w", err)
				}
			} else {
				if err = migration.Steps(int(steps)); err != nil {
					return fmt.Errorf("Falied to upgrade for %d step(s): %w", steps, err)
				}
			}

			return nil
		},
	}

	migrationDownCmd = &cobra.Command{
		Use:   "down",
		Short: "Migration downgrade",
		RunE: func(cmd *cobra.Command, args []string) error {
			err := cmd.Parent().RunE(cmd.Parent(), args)
			if err != nil {
				return fmt.Errorf("Failed to migrate: %w", err)
			}

			if err = migration.Steps(-1); err != nil {
				return fmt.Errorf("Falied to downgrade for 1 step: %w", err)
			}

			return nil
		},
	}
)

func init() {
	migrationCmd.AddCommand(migrationUpCmd, migrationDownCmd)

	migrationCmd.PersistentFlags().UintVarP(
		&steps, stepsFlag, stepsFlagShort, stepsDefault, stepsUsage,
	)

}
