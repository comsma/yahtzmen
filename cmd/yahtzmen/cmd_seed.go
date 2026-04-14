package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"strings"

	"github.com/comsma/yahtzmen/internal/config"
	"github.com/comsma/yahtzmen/internal/models"
	_ "github.com/go-sql-driver/mysql"
	"github.com/urfave/cli/v3"
)

var seedProducts = []models.CreateProductParams{
	{
		ID:          "seed_chair_walnut_01",
		Name:        "Walnut Accent Chair",
		Price:       "1200.00",
		Description: sql.NullString{String: "A hand-crafted accent chair in solid American walnut with mortise and tenon joinery.", Valid: true},
		Notes:       sql.NullString{String: "Lead time 6–8 weeks.", Valid: true},
		Dimensions:  sql.NullString{String: `{"w": 28, "d": 30, "h": 34, "unit": "in"}`, Valid: true},
		Features:    sql.NullString{String: "Solid walnut frame, hand-rubbed oil finish, removable cushion", Valid: true},
		Islive:      true,
	},
	{
		ID:          "seed_table_oak_01",
		Name:        "White Oak Dining Table",
		Price:       "3800.00",
		Description: sql.NullString{String: "A live-edge white oak dining table with a steel hairpin base. Seats six comfortably.", Valid: true},
		Notes:       sql.NullString{String: "Custom sizing available.", Valid: true},
		Dimensions:  sql.NullString{String: `{"w": 42, "d": 84, "h": 30, "unit": "in"}`, Valid: true},
		Features:    sql.NullString{String: "Live-edge slab, steel hairpin legs, matte lacquer finish", Valid: true},
		Islive:      true,
	},
	{
		ID:          "seed_shelf_maple_01",
		Name:        "Floating Maple Shelf",
		Price:       "420.00",
		Description: sql.NullString{String: "A minimalist wall-mounted shelf in hard maple. Hidden bracket system for a clean look.", Valid: true},
		Notes:       sql.NullString{String: "Mounting hardware included.", Valid: true},
		Dimensions:  sql.NullString{String: `{"w": 48, "d": 10, "h": 2, "unit": "in"}`, Valid: true},
		Features:    sql.NullString{String: "Hard maple, hidden steel brackets, natural oil finish", Valid: true},
		Islive:      true,
	},
	{
		ID:          "seed_bench_ash_01",
		Name:        "Ash Entryway Bench",
		Price:       "680.00",
		Description: sql.NullString{String: "A sturdy entryway bench in white ash with woven leather seat straps.", Valid: true},
		Notes:       sql.NullString{String: "Leather color options available on request.", Valid: true},
		Dimensions:  sql.NullString{String: `{"w": 48, "d": 14, "h": 18, "unit": "in"}`, Valid: true},
		Features:    sql.NullString{String: "White ash frame, hand-laced leather straps, beeswax finish", Valid: true},
		Islive:      true,
	},
	{
		ID:          "seed_desk_cherry_01",
		Name:        "Cherry Writing Desk",
		Price:       "2200.00",
		Description: sql.NullString{String: "A refined writing desk in American cherry with two dovetail drawers.", Valid: true},
		Notes:       sql.NullString{String: "Drawer pulls hand-forged in bronze.", Valid: true},
		Dimensions:  sql.NullString{String: `{"w": 54, "d": 26, "h": 30, "unit": "in"}`, Valid: true},
		Features:    sql.NullString{String: "Solid cherry, dovetail drawers, bronze pulls, hand-rubbed finish", Valid: true},
		Islive:      true,
	},
}

func seedCmd(cfg **config.Config) *cli.Command {
	return &cli.Command{
		Name:  "seed",
		Usage: "Insert sample products into the database (safe to run multiple times)",
		Action: func(ctx context.Context, cmd *cli.Command) error {
			return runSeed(*cfg)
		},
	}
}

func runSeed(cfg *config.Config) error {
	db, err := sql.Open("mysql", cfg.DSN)
	if err != nil {
		return fmt.Errorf("db open: %w", err)
	}
	defer db.Close()

	q := models.New(db)

	inserted := 0
	skipped := 0
	for _, p := range seedProducts {
		err := q.CreateProduct(context.Background(), p)
		if err != nil {
			// duplicate key = already seeded, skip
			if isDuplicateKey(err) {
				skipped++
				continue
			}
			return fmt.Errorf("insert %q: %w", p.Name, err)
		}
		inserted++
	}

	log.Printf("seed complete: %d inserted, %d already existed", inserted, skipped)
	return nil
}

func isDuplicateKey(err error) bool {
	return err != nil && strings.Contains(err.Error(), "1062")
}
