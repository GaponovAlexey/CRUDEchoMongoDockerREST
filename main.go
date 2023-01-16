package main

import (
	"fmt"
	"log"
	"mongo/db/config"
	"net/http"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/net/context"

)

var (
	c   *mongo.Client
	db  *mongo.Database
	col *mongo.Collection
	cfg config.Properties
	ctx = context.Background()
)

func cancel(e error) {
	if e != nil {
		log.Fatal("ERROR ----> ", e)
	}
}

func init() {
	err := cleanenv.ReadEnv(&cfg)
	cancel(err)

	connectURI := fmt.Sprintf("mongodb://%s:%s", cfg.DBHost, cfg.DBPort)
	c, err := mongo.Connect(ctx, options.Client().ApplyURI(connectURI))
	cancel(err)
	db = c.Database(cfg.DBName)
	col = db.Collection(cfg.CollectionName)
}

func main() {

	log.Println("start - >>", cfg.DBPort)

	e := echo.New()

	e.GET("/", createProduct)

	e.Logger.Fatal(e.Start(fmt.Sprintf("%s:%s", cfg.Host, cfg.Port)))
	log.Println("end")
}

func createProduct(e echo.Context) error {
	return e.JSON(http.StatusOK, "you")
}
