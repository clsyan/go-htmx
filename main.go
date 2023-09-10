package main

import (
	"go-htmx/configuration"
	"go-htmx/pkg/controllers"
	"go-htmx/pkg/handlers"
	"log"
	"path/filepath"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/josharian/tstruct"
)

func main() {
	path, _ := filepath.Abs("./views")

	engine := html.New(path, ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	db := configuration.NewDatabase()

	defer db.Close()

	driver, err := mysql.WithInstance(db, &mysql.Config{})

	if err != nil {
		panic(err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://./migrations",
		"go-htmx", driver)

	if err != nil {
		panic(err)
	}

	m.Up()

	postService := handlers.NewPostService(db)

	postController := controllers.NewPostController(&postService)

	postController.Route(app)

	log.Fatal(app.Listen(":3000"))
}
