package main

import (
	"fmt"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/cobra"
	"github.com/uptrace/bun/migrate"
)

func NewDbCommands(migrator *migrate.Migrator) *cobra.Command {
	var dbCmd = &cobra.Command{
		Use: "db",
		Run: func(cmd *cobra.Command, args []string) {

		},
	}
	var initCmd = &cobra.Command{
		Use: "init",
		Run: func(cmd *cobra.Command, args []string) {
			err := migrator.Init(cmd.Context())
			if err != nil {
				panic(err)
			}
		},
	}
	var migrateCmd = &cobra.Command{
		Use: "migrate",
		Run: func(cmd *cobra.Command, args []string) {

			group, err := migrator.Migrate(cmd.Context())
			if err != nil {
				panic(err)
				return
			}

			if group.ID == 0 {
				fmt.Printf("there are no new migrations to run\n")
			}

			fmt.Printf(
				"migrated to %s\n",
				group,
			)
			return

		},
	}
	var rollBackCmd = &cobra.Command{
		Use: "rollback",
		Run: func(cmd *cobra.Command, args []string) {
			if err := migrator.Lock(cmd.Context()); err != nil {
				panic(err)
			}
			defer migrator.Unlock(cmd.Context()) //nolint:errcheck

			group, err := migrator.Rollback(cmd.Context())
			if err != nil {
				panic(err)
			}
			if group.IsZero() {
				fmt.Printf("there are no groups to roll back\n")
				panic(err)
			}
			fmt.Printf(
				"rolled back %s\n",
				group,
			)
		},
	}
	var createGoCmd = &cobra.Command{
		Use:   "create_go",
		Short: "creates a go based migration",
		Run: func(cmd *cobra.Command, args []string) {

			name := strings.Join(
				args,
				"_",
			)

			mf, err := migrator.CreateGoMigration(
				cmd.Context(),
				name,
			)
			if err != nil {

			}
			fmt.Printf(
				"created migration %s (%s)\n",
				mf.Name,
				mf.Path,
			)
		},
	}
	dbCmd.AddCommand(initCmd)
	dbCmd.AddCommand(migrateCmd)
	dbCmd.AddCommand(rollBackCmd)
	dbCmd.AddCommand(createGoCmd)

	return dbCmd
}
