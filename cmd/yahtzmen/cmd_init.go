package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/comsma/yahtzmen/internal/config"
	"github.com/go-sql-driver/mysql"
	"github.com/pressly/goose/v3"
	"github.com/urfave/cli/v3"
)

func initCmd(cfg **config.Config) *cli.Command {
	return &cli.Command{
		Name:  "init",
		Usage: "Create the database if not exists and run all migrations",
		Action: func(ctx context.Context, cmd *cli.Command) error {
			if err := createDatabase((*cfg).DSN); err != nil {
				return fmt.Errorf("create database: %w", err)
			}
			log.Println("database ready")

			if err := runMigrations((*cfg).DSN); err != nil {
				return fmt.Errorf("migrations: %w", err)
			}
			log.Println("migrations applied")

			return nil
		},
	}
}

func createDatabase(dsn string) error {
	parsed, err := mysql.ParseDSN(dsn)
	if err != nil {
		return fmt.Errorf("parse dsn: %w", err)
	}

	dbName := parsed.DBName
	parsed.DBName = ""

	db, err := sql.Open("mysql", parsed.FormatDSN())
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec(fmt.Sprintf("CREATE DATABASE IF NOT EXISTS `%s`", dbName))
	return err
}

func runMigrations(dsn string) error {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return err
	}
	defer db.Close()

	if err := goose.SetDialect("mysql"); err != nil {
		return err
	}

	return goose.Up(db, "migrations")
}
