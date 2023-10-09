package migrations

import (
	"context"
	"fmt"

	"github.com/uptrace/bun"
	"yahtzmen/model"
)

func init() {
	Migrations.MustRegister(
		func(ctx context.Context, db *bun.DB) error {
			fmt.Print(" [up migration] ")
			_, err := db.NewCreateTable().Model((*model.Product)(nil)).Exec(ctx)
			if err != nil {
				return err
			}
			return nil
		},
		func(ctx context.Context, db *bun.DB) error {
			fmt.Print(" [down migration] ")
			_, err := db.NewDropTable().Model((*model.Product)(nil)).IfExists().Exec(ctx)
			if err != nil {
				return err
			}
			return nil
		},
	)
}
