package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/cobra"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/mysqldialect"
	"github.com/uptrace/bun/migrate"
	"yahtzmen/controller"
	"yahtzmen/core/http/requests"
	"yahtzmen/core/server"
	"yahtzmen/model/migrations"
	"yahtzmen/repository"
)

var rootCmd = &cobra.Command{
	Use:   "app",
	Short: "A web framework build for a E-Commerce store",
	Run: func(cmd *cobra.Command, args []string) {
		instance := echo.New()

		instance.Use(middleware.CORSWithConfig(middleware.CORSConfig{AllowHeaders: []string{"http://localhost:5173"}}))
		instance.Validator = &requests.CustomValidator{Validator: validator.New()}

		db, err := repository.NewDB()
		if err != nil {
			panic(err)
		}
		productRepo := repository.NewProductRepository(db)
		controller.NewProductController(
			instance,
			productRepo,
		)

		httpServer := server.NewHttpServer(instance)

		httpServer.Start()
		defer httpServer.Stop()

		// Listen for OS signals to perform a graceful shutdown
		log.Println("listening signals...")
		c := make(
			chan os.Signal,
			1,
		)
		signal.Notify(
			c,
			os.Interrupt,
			syscall.SIGHUP,
			syscall.SIGINT,
			syscall.SIGQUIT,
			syscall.SIGTERM,
		)
		<-c
		log.Println("graceful shutdown...")

	},
}

func init() {
	sqldb, err := sql.Open(
		"mysql",
		"root:my-secret-pw@/yahtzmen",
	)
	if err != nil {
		panic(err)
	}
	db := bun.NewDB(
		sqldb,
		mysqldialect.New(),
	)

	migrator := migrate.NewMigrator(
		db,
		migrations.Migrations,
	)
	rootCmd.AddCommand(NewDbCommands(migrator))
}
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func main() {
	Execute()
}
