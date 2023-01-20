package main

import (
	"fmt"
	"mongo/db/config"
	"mongo/db/handlers"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"github.com/labstack/gommon/random"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/net/context"

)

const (
	correlationID = "X-Correlation-ID"
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

// middleware
func addCorelationID(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Request().Header.Get(correlationID)
		var newId string
		if id == "" {
			newId = random.String(12)
		} else {
			newId = id
		}

		c.Request().Header.Set(correlationID, newId)
		c.Response().Header().Set(correlationID, newId)
		return next(c)
	}
}

func main() {
	e := echo.New()
	e.Logger.SetLevel(log.ERROR)
	e.Pre(middleware.RemoveTrailingSlash())
	e.Pre(addCorelationID)

	h := handlers.ProductHandler{
		Col: col,
	}

	e.POST("/", h.CreateProducts, middleware.BodyLimit("1M"))
	e.GET("/", h.GetProduct)
	e.GET("/", h.GetProductID)
	//end
	e.Logger.Fatal(e.Start(fmt.Sprintf("%s:%s", cfg.Host, cfg.Port)))
}
