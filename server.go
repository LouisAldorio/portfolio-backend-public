package main

import (
	"log"
	"myapp/dataloaders"
	"myapp/graph"
	"myapp/graph/generated"
	"myapp/utils"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/rs/cors"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

const defaultPort = "8080"

func main() {

	// err := godotenv.Load()
	// if err != nil {
	// 	log.Fatal("Error loading .env file")
	// }

	// grpcClient.SendMail(model.NewMail{
	// 	Subject: "",
	// 	Title: "",
	// 	Message: "",
	// 	DestinationEmail: "",
	// })

	router := chi.NewRouter()
	router.Use(utils.Auth())
	router.Use(dataloaders.Middleware)

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	handler := cors.New(cors.Options{
		AllowedHeaders: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "HEAD", "OPTIONS"},
	}).Handler(srv)

	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", handler)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
