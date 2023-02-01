package main

import (
	"context"
	"entdemo/ent"
	"entdemo/handlers"
	_ "github.com/lib/pq"
	"log"
	"net/http"
)

func main() {
	client, err := ent.Open("postgres", "host=localhost port=5432 user=postgres dbname=Densityasmt password=Anirudh@123 sslmode=disable")
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}
	defer client.Close()
	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	http.HandleFunc("/addbook", handlers.AddBook)
	http.HandleFunc("/getbookbyname", handlers.GetBookByName)
	http.HandleFunc("/delbookbyname", handlers.DelBookByName)
	http.HandleFunc("/getbookbyauthor", handlers.GetBookByAuthor)
	http.HandleFunc("/getallbooks", handlers.GetAllBook)
	http.HandleFunc("/getbookbypublishdate", handlers.GetBookByPublishDate)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
