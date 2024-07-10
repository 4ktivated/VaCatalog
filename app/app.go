package app

import (
	"context"
	"log"
	"net/http"
	"os"
	"some_code/dbase"

	"github.com/gocraft/web"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Run(ctx context.Context) error {
	godotenv.Load()
	uri := os.Getenv("MONGO_URI")
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		log.Print(err)
		return err
	}

	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			log.Print(err)
		}
	}()

	VacsDAO, err := dbase.NewVacsDAO(ctx, client, "vacs")
	if err != nil {
		log.Print(err)
		return err
	}

	service := NewService(VacsDAO)
	httpHandler := NewHandler(service)

	return http.ListenAndServe("localhost:8000", endPointInit(httpHandler))

}

func endPointInit(h *Handler) *web.Router {
	router := web.New(*h)
	router.Get("/", WrapEndpoint(h.Main_Page))
	router.Get("/:lang", WrapEndpoint(h.GetbyLang))
	// router.Post("/update/:shortUrl", WrapEndpoint(h.Update))
	router.Get("/ping", WrapEndpoint(h.Ping))
	return router

}
