package cmd

import (
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/spf13/cobra"
)

const (
	migrationsDir = "migrations"
)

var (
	migration *migrate.Migrate

	migrationCmd = &cobra.Command{
		Use:   "migration",
		Short: "Database migrations",
	}

	migrationMakeCmd = &cobra.Command{
		Use:   "make",
		Short: "Migration make",
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
			fmt.Print("Migration upgrade\n")
			return nil
		},
	}

	migrationDownCmd = &cobra.Command{
		Use:   "down",
		Short: "Migration downgrade",
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Print("Migration downgrade\n")
			return nil
		},
	}
)

func init() {
	migrationCmd.AddCommand(migrationMakeCmd, migrationUpCmd, migrationDownCmd)
}
